package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	// errors.Is is capable of detecting sentinel errors wrapped in the error chain
	if errors.Is(err, errSentinel) {
		fmt.Println("errSentinel detected in the error chain with errors.Is")
	}
	if err == errSentinel {
		fmt.Println("errSentinel detected in the error chain by ==")
	}
}

var errSentinel = errors.New("test")

func foo() error {
	return fmt.Errorf("error : %w", bar())
}

func bar() error {
	return errSentinel
}
