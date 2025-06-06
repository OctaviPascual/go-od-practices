package bluesky

import (
	"errors"
	"math/rand/v2"
)

var ErrInvalidUsername = errors.New("invalid username")

var ErrUnknownAvailability = errors.New("unknown availability")

func IsAvailableWithSentinel(_ string) (bool, error) {
	switch rand.IntN(3) {
	case 0:
		return false, ErrInvalidUsername
	case 1:
		return false, ErrUnknownAvailability
	default:
		return true, nil
	}
}
