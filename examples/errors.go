package main

import (
	"errors"
	"fmt"
)

// imagine importing this from a deep dependency
var myError = errors.New("something went wrong")

func try() error {
	// imagine calling a deep stack of functions to get to
	return fmt.Errorf("try: %w", &myError)
}

func doIt() {
	err := try()
	if err != nil {
		if errors.Is(err, myError) {
			// we now know the type of the error
		}
	}
}
