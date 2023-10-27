package precision

import (
	"strings"
	"unsafe"

	"github.com/shopspring/decimal"
)

type params struct {
	digitsBefore      int
	dot               int
	leadingZeros      int
	roundTo           int
	significantDigits int
	maxDecimals       int
}

func ToSignificant(x decimal.Decimal, significantDigits, maxDecimals uint) decimal.Decimal {
	if x.IsZero() {
		return x
	}

	bytes := []byte(x.String())
	dot := strings.IndexRune(unsafeToString(bytes), '.')
	p := params{
		digitsBefore:      getDigitsBefore(bytes, dot),
		dot:               dot,
		significantDigits: int(significantDigits),
		maxDecimals:       int(maxDecimals),
	}

	if p.digitsBefore > p.significantDigits {
		roundInt(bytes, p.digitsBefore-p.significantDigits)
		return buildDecimal(bytes)
	} else if p.digitsBefore > 0 && p.dot == -1 {
		// For some minuscule numbers like `math.SmallestNonzeroFloat64`
		// string representation may be different from its decimal representation,
		// So it's better to rebuild it from scratch.
		return buildDecimal(bytes)
	}

	if truncated := truncateFloat(bytes, p); truncated != nil {
		bytes = truncated
	}
	p.leadingZeros = getLeadingZeros(bytes, p.dot)

	// Calculate the position of the last significant digit
	if p.digitsBefore > 0 {
		p.roundTo = int(significantDigits) - p.digitsBefore - p.leadingZeros
		if p.roundTo < 0 {
			p.roundTo = p.dot + 1
		}
	} else {
		p.roundTo = p.dot + p.leadingZeros + int(significantDigits)
		if p.roundTo > int(maxDecimals) {
			p.roundTo = int(maxDecimals)
		}
	}

	// Make sure rounding occurs inside buffer
	if p.roundTo > len(bytes) { // e.g. 0.0001
		p.roundTo = len(bytes) - 1
	}

	roundFloat(bytes, p)
	return buildDecimal(bytes)
}

// unsafeToString effectively tricks the Go runtime into treating
// the underlying memory of the byte slice as if it were a string
// to avoid additional allocation.
//
// !!! DO NOT MODIFY STRING YOU GET BACK FROM THIS FUNCTION !!!
//
// Since strings in Go are immutable,
// **YOU MUST ENSURE THAT THE ORIGINAL BYTE SLICE ISN'T MODIFIED** after this conversion,
// or else you risk causing undefined behavior.
func unsafeToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func getDigitsBefore(bytes []byte, dot int) int {
	if dot == -1 {
		return len(bytes)
	}

	var digitsBefore int
	for _, b := range bytes[:dot] {
		if b == '0' {
			continue
		}
		digitsBefore++
	}
	return digitsBefore
}

func getLeadingZeros(bytes []byte, dot int) int {
	var leadingZeros int
	for _, b := range bytes[dot+1:] {
		if b != '0' {
			continue
		}
		leadingZeros++
	}
	return leadingZeros
}

// truncateFloat truncates decimals that are greater than max allowed decimals
func truncateFloat(bytes []byte, p params) []byte {
	bytesBeforeDot := p.dot + 1
	bytesAfterDot := len(bytes) - bytesBeforeDot
	if bytesAfterDot > p.maxDecimals {
		return bytes[:bytesBeforeDot+p.maxDecimals]
	}
	return nil
}

func roundInt(bytes []byte, to int) {
	bytesLen := len(bytes)
	for i := bytesLen - 1; i > bytesLen-to-2; i-- {
		if bytes[i] == '9' {
			bytes[i] = '0'
		} else {
			bytes[i]++

			if i != bytesLen-1 {
				bytes[i+1] = '0'
			}
		}
	}
}

func roundFloat(bytes []byte, p params) {
	existingSignificantDigits := len(bytes[p.dot+1+p.leadingZeros:])
	enoughDecimals := p.leadingZeros+existingSignificantDigits <= p.maxDecimals
	if p.digitsBefore == 0 && enoughDecimals && existingSignificantDigits <= p.significantDigits {
		// No rounding is required for 0.0001,
		// as it already has enough significant digits
		// and there's nothing to round
		return
	}

	// Carry mechanism to round up
	var carry bool
	for i := len(bytes) - 1; i > p.roundTo-1; i-- {
		if bytes[i] == '9' {
			bytes[i] = '0'
			carry = true
		} else if bytes[i] >= '5' {
			bytes[i] = '0'
			carry = true
		} else if carry {
			bytes[i]++
			carry = false
		} else {
			bytes[i] = '0'
		}
	}

	// If there's still a carry after the dot, increment the part before the dot
	if carry {
		i := p.dot - 1
		for i >= 0 {
			if bytes[i] == '9' {
				bytes[i] = '0'
				i--
			} else {
				bytes[i]++
				break
			}
		}
		// If we've gone through all the bytes, and they were all '9'
		if i == -1 {
			bytes = append([]byte{'1'}, bytes...)
		}
	}
}

func buildDecimal(bytes []byte) decimal.Decimal {
	result, err := decimal.NewFromString(unsafeToString(bytes))
	if err != nil {
		panic("failed to parse converted bytes to decimal")
	}
	return result
}
