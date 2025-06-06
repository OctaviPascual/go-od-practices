package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/OctaviPascual/go-od-practices/errors/bluesky"
)

func main() {
	sentinelError()
	concreteError()
}

func sentinelError() {
	username := "😱"

	available, err := bluesky.IsAvailableWithSentinel(username)
	if errors.Is(err, bluesky.ErrInvalidUsername) {
		fmt.Printf("%q is not valid on Bluesky.\n", username)
		return
	}
	if errors.Is(err, bluesky.ErrUnknownAvailability) {
		fmt.Printf("The availability of %q on Bluesky could not be checked.\n", username)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(strconv.FormatBool(available))
}

func concreteError() {
	username := "😱"

	available, err := bluesky.IsAvailableWithConcreteError(username)
	var iuerr *bluesky.InvalidUsernameError
	if errors.As(err, &iuerr) {
		fmt.Printf("%q is not valid on Bluesky.\n", username)
		return
	}
	var uaerr *bluesky.UnknownAvailabilityError
	if errors.As(err, &uaerr) {
		fmt.Printf("The availability of %q on Bluesky could not be checked.\n", username)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(strconv.FormatBool(available))
}
