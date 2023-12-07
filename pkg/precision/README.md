# Precision Package

The `precision` package provides a method for converting decimal numbers to significant figures while maintaining precision within a defined scale.

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
	"github.com/layer-3/clearsync/precision"
)
```

### Function: `ToSignificant`

The `ToSignificant` function takes a decimal input, the desired number of significant digits (`sigDigits`), and a maximum scale (`maxScale`).
It performs rounding operations to ensure the output has the specified significant digits within the defined scale.

#### Parameters

- `input decimal.Decimal`: The input decimal number.
- `sigDigits int32`: Number of significant digits desired in the output.
- `maxScale int32`: Maximum scale allowed for the output decimal.

#### Returns

- `decimal.Decimal`: The decimal value with the specified significant digits and within the defined scale.

#### Example

```go
input := decimal.NewFromFloat(123.456789)
significantDigits := int32(5)
maxScale := int32(2)

result := precision.ToSignificant(input, significantDigits, maxScale)
fmt.Println(result) // Output: 123.46
```

## Contributing

Feel free to contribute by forking the repository, making changes, and creating pull requests. Ensure you maintain the existing code conventions and add tests for any new functionality.
