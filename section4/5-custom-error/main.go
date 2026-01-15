package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDenominatorZeroValue = errors.New("denominator cannot zero")
var ErrNumeratorToLarge = errors.New("number too large")

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (op OpError) Error() string {
	return op.Message
}

func NewOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

func DoSomething() error {
	return NewOpError("doSomething", 100, "do something failed", time.Now())
}

func divide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, ErrDenominatorZeroValue
	}

	if numerator > 1000 {
		return 0, ErrNumeratorToLarge
	}

	return numerator / denominator, nil
}

func main() {
	value, err := divide(1001, 2)
	if err != nil {
		if errors.Is(err, ErrDenominatorZeroValue) {
			fmt.Println("do something else")
		} else if errors.Is(err, ErrNumeratorToLarge) {
			fmt.Println("number too large")
		}
		return
	}

	fmt.Println(value)
}
