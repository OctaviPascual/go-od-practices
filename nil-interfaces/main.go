package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

type wrappedError struct {
	err     error
	details string
}

func newWrappedError(err error, details string) *wrappedError {
	return &wrappedError{
		err:     err,
		details: details,
	}
}

func (e *wrappedError) Error() string {
	// This checks if both type T and concrete value V of e.err are nil.
	if e.err == nil {
		return fmt.Sprintf("error=nil_interface, details=%s", e.details)
	}
	// This checks if concrete value V of e.err is nil.
	if reflect.ValueOf(e.err).IsNil() {
		return fmt.Sprintf("error=nil_concrete_value, details=%s", e.details)
	}
	return fmt.Sprintf("error=%s, details=%s", e.err.Error(), e.details)
}

func main() {
	wrappedError := newWrappedError(nil, "details_1")
	fmt.Println(wrappedError.Error())

	wrappedError = newWrappedError((*os.PathError)(nil), "details_2")
	fmt.Println(wrappedError.Error())

	wrappedError = newWrappedError(errors.New("some error"), "details_3")
	fmt.Println(wrappedError.Error())
}
