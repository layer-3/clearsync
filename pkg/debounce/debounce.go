// Package debounce provides a mechanism for rate-limiting and retrying RPC requests
// in a controlled manner. It includes support for handling rate limit errors by
// dynamically adjusting retry intervals and limiting burst rates. The package uses
// a rate limiter to ensure that requests stay within allowed limits, and provides
// detailed handling of HTTP and JSON-RPC error codes to decide whether retries
// should be attempted or the error is unrecoverable.
package debounce

import (
	"context"
	"strings"
	"sync/atomic"

	"github.com/ipfs/go-log/v2"
	"golang.org/x/time/rate"
)

const (
	// rateLimit is the maximum number of requests per second.
	// As of spring 2024, Infura enables 10 req/s rate limit for free plan.
	rateLimit = 10
	// rateLimitDebounceSeconds is the upper limit of seconds to wait before retrying the request.
	// Wait time is increased slowly until it matches this value.
	rateLimitDebounceSeconds = 60
	maxBurst                 = rateLimit * rateLimitDebounceSeconds
	// burstDecrementAfter "debounces" the burst decrement.
	//
	// It is a number of consecutive successful requests
	// after which the burst is decremented.
	// This is to find the right tempo (request rate)
	// if the RPC is under heavy load or is unstable.
	//
	// Consider the following case: if request Nth-2 was rate-limited and Nth-1 was successful,
	// it does not mean that Nth request will be successful.
	// Maybe the rate limit is right on the edge,
	// then it does not make any sense to decrease the burst after Nth-1.
	// Hence burstDecrementAfter
	burstDecrementAfter = 5
	// burstBias is to avoid rate limiting on the first few requests.
	// The client will eventually work its way up to the actual rate limit.
	burstBias = 3
)

var (
	// rpcRateLimiter limits the number of requests to the RPC provider for DEX drivers.
	rpcRateLimiter = rate.NewLimiter(rate.Limit(rateLimit), maxBurst)

	// burst is a shared counter that is used to adjust the rate limiter's burst.
	burst = atomic.Int64{}
	// burstSuccessCount is a shared counter
	// that is used to track the number of successful requests
	// after which the burst is decremented.
	burstSuccessCount = atomic.Int64{}

	httpRpcErrors = map[int]rpcError{
		400: {
			Recoverable: false,
			Patterns:    []string{"400"},
			Message:     "Bad request: Incorrect HTTP Request type or invalid characters, ensure that your request body and format is correct.",
		},
		401: {
			Recoverable: false,
			Patterns: []string{
				"401",
				"project id required in the url",
				"invalid project id",
				"invalid project id or project secret",
				"invalid JWT",
			},
			Message: "Unauthorized: This can happen when one or multiple security requirements are not met.",
		},
		403: {
			Recoverable: false,
			Patterns:    []string{"403"},
			Message:     "Forbidden: The request was intentionally refused due to specific settings mismatch, check your key settings.",
		},
		404: {
			Recoverable: false,
			Patterns:    []string{"404"},
			Message:     "RPC endpoint doesn't exist",
		},
		405: {
			Recoverable: false,
			Patterns:    []string{"405"},
			Message:     "HTTP Method Not Allowed",
		},
		429: {
			Recoverable: true,
			Patterns: []string{
				"429",
				"project ID request rate exceeded",
				"daily request count exceeded",
				"request rate limited",
				"Too Many Requests",
			},
			Message: "Too Many Requests: The daily request total or request per second are higher than your plan allows. Refer to the Avoid rate limiting topic for more information.",
		},
		500: {
			Recoverable: true,
			Patterns:    []string{"500"},
			Message:     "Internal Server Error: Error while processing the request on the server side.",
		},
		502: {
			Recoverable: true,
			Patterns:    []string{"502"},
			Message:     "Bad Gateway: Indicates a communication error which can have various causes, from networking issues to invalid response received from the server.",
		},
		503: {
			Recoverable: true,
			Patterns: []string{
				"503",
				"HTTP status 503 Service Unavailable",
				"Service Unavailable",
				"service unavailable",
			},
			Message: "Service Unavailable: Indicates that the server is not ready to handle the request.",
		},
		504: {
			Recoverable: true,
			Patterns:    []string{"504"},
			Message:     "Gateway Timeout: The request ended with a timeout, it can indicate a networking issue or a delayed or missing response from the server.",
		},
	}
	jsonRpcErrors = map[int]rpcError{
		-32700: {
			Recoverable: false,
			Message:     "Parse error: The JSON request is invalid, this can be due to syntax errors. (Standard error)",
		},
		-32600: {
			Recoverable: false,
			Message:     "Invalid request: The JSON request is possibly malformed. (Standard error)",
		},
		-32601: {
			Recoverable: false,
			Message:     "Method not found: The method does not exist, often due to a typo in the method name or the method not being supported. (Standard error)",
		},
		-32602: {
			Recoverable: false,
			Message:     "Invalid argument: Invalid method parameters. (Standard error)",
		},
		-32603: {
			Recoverable: false,
			Message:     "Internal error: An internal JSON-RPC error, often caused by a bad or invalid payload. (Standard error)",
		},
		-32000: {
			Recoverable: false,
			Message:     "Invalid input: Missing or invalid parameters, possibly due to server issues or a block not being processed yet. (Non-standard error)",
		},
		-32001: {
			Recoverable: false,
			Message:     "Resource not found: The requested resource cannot be found, possibly when calling an unsupported method. (Non-standard error)",
		},
		-32002: {
			Recoverable: false,
			Message:     "Resource unavailable: The requested resource is not available. (Non-standard error)",
		},
		-32003: {
			Recoverable: false,
			Message:     "Transaction rejected: The transaction could not be created. (Non-standard error)",
		},
		-32004: {
			Recoverable: false,
			Message:     "Method not supported: The requested method is not implemented. (Non-standard error)",
		},
		-32005: {
			Recoverable: false,
			Message:     "Limit exceeded: The request exceeds your request limit. For more information, refer to Avoid rate limiting. (Non-standard error)",
		},
		-32006: {
			Recoverable: false,
			Message:     "JSON-RPC version not supported: The version of the JSON-RPC protocol is not supported. (Non-standard error)",
		},
	}
)

func init() {
	burst.Store(rateLimit * burstBias) // to avoid rate limiting on the first few requests
}

type rpcError struct {
	Message     string
	Patterns    []string
	Recoverable bool
}

// Debounce is a wrapper around the rate limiter
// that retries the request if it fails with rate limit error.
func Debounce(
	ctx context.Context,
	logger *log.ZapEventLogger,
	f func(context.Context) error,
) error {
outer:
	for {
		currBurst := int(burst.Load())
		if err := rpcRateLimiter.WaitN(ctx, int(currBurst)); err != nil {
			if logger != nil {
				logger.Warnw("failed to acquire rate limiter", "error", err)
			}
			return err
		}

		err := f(ctx)
		if err == nil {
			successCount := burstSuccessCount.Add(1)
			if currBurst > 1 && successCount >= burstDecrementAfter {
				burstSuccessCount.Store(0) // reset counter
				burst.Add(-1)              // decrease burst on success
				logger.Debugw("decremented burst", "new_burst", currBurst-1)
			}

			return nil
		}

		// Search for the error in the list of known HTTP RPC errors
		for _, httpRpcError := range httpRpcErrors {
			for _, pattern := range httpRpcError.Patterns {
				if strings.Contains(err.Error(), pattern) && httpRpcError.Recoverable {
					if logger != nil {
						logger.Warnw("recoverable error",
							"burst", currBurst,
							"message", httpRpcError.Message,
							"error", err)
					}

					// Adjust burst to wait a bit longer
					if currBurst < maxBurst {
						burst.Add(1)
					}

					continue outer // retry the request after a while
				}
			}
		}

		return err
	}
}
