package main

import (
	"context"
	"fmt"
	"sync"
)

type (
	reportable interface {
		Report(ctx context.Context)
	}

	retryable interface {
		CanRetry() bool
	}

	mySpecialError struct {
		once               sync.Once
		retries            int
		importantDebugInfo string
	}
)

func (e *mySpecialError) Error() string {
	return e.importantDebugInfo
}

func (e *mySpecialError) Report(ctx context.Context) {
	e.once.Do(func() {
		fmt.Println(e.Error())
	})
}

func (e *mySpecialError) CanRetry() bool {
	return e.retries > 0
}

func DoSomethingThatMightFail(b bool) error {
	if b {
		return nil
	}

	return &mySpecialError{
		retries:            1,
		importantDebugInfo: "test stuff",
	}
}

func main() {
	err := DoSomethingThatMightFail(false)
	switch e := err.(type) {
	case reportable:
		e.Report(context.TODO())
		fallthrough
	case retryable:
		if e.CanRetry() {
			fmt.Println("try again!")
		}
	case error:
		fmt.Println("any other type of error")
	default:
		fmt.Println("nil?", e)
	}
}
