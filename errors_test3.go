package main

import (
	"errors"
	"fmt"
	"os"
)

type ErrorTest3First struct {
	Err error
}

func (e ErrorTest3First) Error() string {
	return fmt.Sprintf("ErrorTest3First : %v", e.Err)
}

type ErrorTest3Second struct {
	Err error
}

func (e ErrorTest3Second) Error() string {
	return fmt.Sprintf("ErrorTest3Second : %v", e.Err)
}

func TestError3Generate() error {
	_, err := os.Open("./testtest")

	if err != nil {
		return ErrorTest3Second{Err: err}
	}

	return nil
}

func CallTestError3() {
	err := TestError3Generate()

	// var e ErrorTest3First
	if errors.As(err, &ErrorTest3First{}) {
		fmt.Println("Error Test3 First")
	}

	if errors.As(err, &ErrorTest3Second{}) {
		fmt.Println("Error Test3 Second")
	}
}
