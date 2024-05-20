package quotes

import (
	"context"
	"strings"

	"github.com/ipfs/go-log/v2"
	"golang.org/x/time/rate"
)

var (
	// rpcRateLimiter limits the number of requests to the RPC provider for DEX drivers.
	// As of spring 2024, Infura enables 10 req/s rate limit for free plan.
	// A lower limit of 5 req/s is used here just to be safe.
	rpcRateLimiter = rate.NewLimiter(5, 1)
	httpRpcErrors  = map[int]rpcError{
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

type rpcError struct {
	Message     string
	Patterns    []string
	Recoverable bool
}

// debounce is a wrapper around the rate limiter
// that retries the request if it fails with rate limit error.
func debounce(logger *log.ZapEventLogger, f func() error) error {
	for {
		if err := rpcRateLimiter.Wait(context.TODO()); err != nil {
			logger.Warnf("failed to acquire rate limiter: %s", err)
		}

		err := f()
		if err == nil {
			return nil
		}

		for _, httpRpcError := range httpRpcErrors {
			for _, pattern := range httpRpcError.Patterns {
				if strings.Contains(err.Error(), pattern) {
					logger.Warn("recoverable error",
						"message", httpRpcError.Message,
						"error", err)
					continue // retry the request after a while
				}
			}
		}

		return err
	}
}
