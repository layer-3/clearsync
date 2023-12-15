# Precision Package

The `precision` package implements functionality for handling the precision of asset prices and quantities following the YIP-0001 - Asset Price Precision specification.

This package includes functions to ensure that asset prices and quantities are represented with the appropriate level of significant digits and precision, aligning with global standards for price-point (PIP) representation in financial markets.
## Installation

To use this package, you can install it using `go get`:

```bash
go get github.com/layer-3/clearsync/precision
```

## Usage

Import the package into your Go code:

```go
import (
	"github.com/shopspring/decimal"
	"github.com/layer-3/clearsync/pkg/precision"
)
```

## Key Features

### ToSignificant

The `ToSignificant` function truncates a decimal number to a specified number of significant digits as per the YIP-0001 specification for price representation.
It is crucial for maintaining standard price precision in trading scenarios.

#### Parameters

- `input decimal.Decimal`: The number to be truncated (non-negative).
- `sigDigits int32`: The number of significant digits to retain.

#### Returns

- `decimal.Decimal`: The truncated decimal number.

#### Example

```go
result := ToSignificant(decimal.NewFromFloat(123.456), 4) // Returns 123.4
```

### Validate

The `Validate` function ensures that a given decimal number adheres to precision rules specified in the YIP-0001. It checks that the input is non-negative and its precision does not exceed a specified limit.

#### Parameters

- `input decimal.Decimal`: The number to be validated.
- `maxPrecision int32`: The maximum allowed precision (number of digits after the decimal point).

#### Returns

- `error`: An error if the input is negative or exceeds the maximum precision.

#### Example

```go
err := Validate(decimal.NewFromFloat(1.234), 3) // Returns nil, as the precision is within the limit.
```

## Usage

Import the package into your Go code:

```go
import (
	"github.com/shopspring/decimal"
	"github.com/layer-3/clearsync/pkg/precision"
)
```

Feel free to adjust the information or add more details as needed!

## Benchmark Results

### ToSignificant Benchmark

```sh
goos: darwin
goarch: amd64
pkg: github.com/layer-3/clearsync/pkg/precision
cpu: VirtualApple @ 2.50GHz
BenchmarkToSignificant_DecimalWithLeadingZeros-10                         	 4991638	       238.6 ns/op	     152 B/op	       6 allocs/op
BenchmarkToSignificant_TruncateDecimals-10                                	 2944626	       407.3 ns/op	     312 B/op	      12 allocs/op
BenchmarkToSignificant_ExactSignificantDigits-10                          	 7927390	       154.6 ns/op	      56 B/op	       4 allocs/op
BenchmarkToSignificant_IntegralPartSizeGreaterThanSignificantDigits-10    	 4253004	       291.8 ns/op	     168 B/op	       9 allocs/op
```

### Validate Benchmark
```sh
goos: darwin
goarch: amd64
pkg: github.com/layer-3/clearsync/pkg/precision
cpu: VirtualApple @ 2.50GHz
BenchmarkValidate_SuccessfulCase-10                                       	490919691	         2.388 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidate_UnsuccessfulCase-10                                     	 8283324	       143.4 ns/op	      80 B/op	       2 allocs/op
```
