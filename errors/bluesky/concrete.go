package bluesky

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type InvalidUsernameError struct {
	username string
}

func (e *InvalidUsernameError) Error() string {
	return fmt.Sprintf("invalid username %q", e.username)
}

type UnknownAvailabilityError struct {
	cause error
}

func (*UnknownAvailabilityError) Error() string {
	return "unknown availability"
}

func (e *UnknownAvailabilityError) Unwrap() error {
	return e.cause
}

func IsAvailableWithConcreteError(username string) (bool, error) {
	switch rand.IntN(3) {
	case 0:
		return false, &InvalidUsernameError{
			username: username,
		}
	case 1:
		return false, &UnknownAvailabilityError{
			cause: errors.New("oh no"),
		}
	default:
		return true, nil
	}
}
