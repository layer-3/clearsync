package opendax_protocol

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ErrIterationDone is returned when no more elements in array to iterate
var ErrIterationDone = errors.New("iteration done")

// ErrInvalidType is returned when type conversion is not successful
var ErrInvalidType = errors.New("invalid type")

// ErrValueNil is returned when next value is nil
var ErrValueNil = errors.New("value is nil")

// ArgIterator is a helper for msg arguments processing.
type ArgIterator struct {
	ind  int
	args []interface{}

	err error
}

// NewArgIterator creates new ready-to-use iterator
func NewArgIterator(args []interface{}) *ArgIterator {
	return &ArgIterator{0, args, nil}
}

// Index returns current iterator position
func (a *ArgIterator) Index() int {
	return a.ind
}

// Err returns first met iterator error1
func (a *ArgIterator) Err() error {
	return a.err
}

// More returns true if there are values not scanned, intended to be used for iterating
func (a *ArgIterator) More() bool {
	return a.Err() == nil && a.ind < len(a.args)
}

// Next returns next value or nil, sets ErrIterationDone
func (a *ArgIterator) Next() interface{} {
	if a.Err() != nil {
		return nil
	}

	if a.ind >= len(a.args) {
		a.err = ErrIterationDone
		return nil
	}

	el := a.args[a.ind]
	a.ind++

	if el == nil {
		a.err = ErrValueNil
		return nil
	}

	return el
}

// NextBool returns next value converted to bool
func (a *ArgIterator) NextBool() bool {
	el := a.Next()
	if err := a.Err(); err != nil {
		return false
	}

	res, ok := el.(bool)
	if !ok {
		a.err = ErrInvalidType
		return false
	}

	return res
}

// NextString returns next value converted to string
func (a *ArgIterator) NextString() string {
	el := a.Next()
	if err := a.Err(); err != nil {
		return ""
	}

	var res string
	var err error

	switch v := el.(type) {
	case float64:
		res, err = strconv.FormatFloat(v, 'f', -1, 64), nil
	case string:
		res, err = v, nil
	default:
		res, err = "", ErrInvalidType
	}

	a.err = err

	return res
}

// NextDecimal returns next value converted to decimal
func (a *ArgIterator) NextDecimal() decimal.Decimal {
	el := a.Next()
	if err := a.Err(); err != nil {
		return decimal.Zero
	}

	var res decimal.Decimal
	var err error

	switch v := el.(type) {
	case float64:
		res, err = decimal.NewFromFloat(v), nil
	case string:
		res, err = decimal.NewFromString(v)
		if err != nil && v == "" {
			return decimal.Zero
		}
	default:
		res, err = decimal.Zero, ErrInvalidType
	}

	a.err = err

	return res
}

// NextUUID returns next value parsed as UUID
func (a *ArgIterator) NextUUID() uuid.UUID {
	el := a.Next()
	if err := a.Err(); err != nil {
		return uuid.UUID{}
	}

	v, ok := el.(string)
	if !ok {
		a.err = ErrInvalidType
		return uuid.UUID{}
	}

	parsed, err := uuid.Parse(v)
	if err != nil {
		a.err = ErrInvalidType
		return uuid.UUID{}
	}

	return parsed
}

// NextFloat64 returns next value converted to float64
func (a *ArgIterator) NextFloat64() float64 {
	el := a.Next()
	if err := a.Err(); err != nil {
		return 0
	}

	switch v := el.(type) {
	case float64:
		return v
	case string:
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			a.err = ErrInvalidType
		}

		return parsed
	default:
		a.err = ErrInvalidType
		return 0
	}
}

// NextUint64 returns next value converted to uint64
func (a *ArgIterator) NextUint64() uint64 {
	el := a.Next()
	if err := a.Err(); err != nil {
		return 0
	}

	switch v := el.(type) {
	case float64:
		intV := uint64(v)

		if float64(intV) != v {
			a.err = ErrInvalidType
		}

		return intV
	case string:
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			a.err = ErrInvalidType
		}

		return parsed
	default:
		a.err = ErrInvalidType
		return 0
	}
}

// NextInt64 returns next value converted to int64
func (a *ArgIterator) NextInt64() int64 {
	el := a.Next()
	if err := a.Err(); err != nil {
		return 0
	}

	switch v := el.(type) {
	case float64:
		intV := int64(v)

		if float64(intV) != v {
			a.err = ErrInvalidType
		}

		return intV
	case string:
		parsed, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			a.err = ErrInvalidType
		}

		return parsed
	default:
		a.err = ErrInvalidType
		return 0
	}
}

// NextInt returns next value converted to int
func (a *ArgIterator) NextInt() int {
	el := a.Next()
	if err := a.Err(); err != nil {
		return 0
	}

	switch v := el.(type) {
	case float64:
		intV := int(v)

		if float64(intV) != v {
			a.err = ErrInvalidType
		}

		return intV
	case string:
		parsed, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			a.err = ErrInvalidType
		}

		return int(parsed)
	default:
		a.err = ErrInvalidType
		return 0
	}
}

// NextSlice returns next value converted to []interface{}
func (a *ArgIterator) NextSlice() []interface{} {
	el := a.Next()
	if err := a.Err(); err != nil {
		return nil
	}

	converted, ok := el.([]interface{})
	if !ok {
		a.err = ErrInvalidType
		return nil
	}

	return converted
}

// NextTimestamp returns next value converted to Unix()
func (a *ArgIterator) NextTimestamp() int64 {
	el := a.Next()

	switch v := el.(type) {
	case string:
		t1, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			t2, err := strconv.ParseFloat(v, 64)
			if err != nil {
				a.err = ErrInvalidType
				return 0
			}
			return int64(t2)
		}
		return t1.Unix()
	case float64:
		t := int64(v)
		return t
	default:
		a.err = ErrInvalidType
		return 0
	}
}

// NextInterface returns next value converted to interface{}
func (a *ArgIterator) NextInterface() interface{} {
	el := a.Next()
	if err := a.Err(); err != nil {
		return nil
	}
	return el
}
