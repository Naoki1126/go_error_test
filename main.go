package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	if err := run(); err != nil {
		fmt.Println(err)
		var c *CustomError
		if errors.As(err, &c) {
			fmt.Println(111111)
		}
	}

	err1 := msgError("Error")
	fmt.Println("Normal Error", err1)

	err2 := WrapError("Test Error Message")
	fmt.Println("[Wrapping Error]", err2)

	var strErr *StrError
	if errors.As(err2, &strErr) {
		fmt.Printf("[Failed] %s\n", strErr)
	}
}

type StrError struct {
	Msg string
}

func (e *StrError) Error() string {
	return fmt.Sprintf("Message : %s", e.Msg)
}

//Errorが実装されていないとcompileErrorになる
func msgError(msg string) error {
	return &StrError{Msg: msg}
}
func WrapError(msg string) error {
	err := msgError(msg)
	return fmt.Errorf("(Wrapping) %w", err)
}

type CustomError struct {
	When time.Time
	What string
	Err  error
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("custom error", c.When, c.What)
}

func run() error {
	return &CustomError{
		When: time.Now(),
		What: "it didint conten",
		Err:  fmt.Errorf("custom error"),
	}
}

func (c *CustomError) Unwrap() error {
	return c.Err
}
