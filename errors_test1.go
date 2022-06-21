package main

import (
	"errors"
	"fmt"
)

type myError struct {
	s string
}

func (obj *myError) Error() string {
	return ""
}

type myErrorImplementedIs struct {
	s string
}

func (obj *myErrorImplementedIs) Is(err error) bool {
	return true
}

func (obj *myErrorImplementedIs) Error() string {
	return ""
}

func ExcecuteTest1() {

	// 比較可能で同じオブジェクトを返す場合は真
	{
		err := &myError{s: "hoge"}
		fmt.Println(errors.Is(err, err))
	}

	// 比較可能でもオブジェクトが違えば偽
	{
		err := &myError{s: "hoge"}
		target := errors.New("hoge")
		fmt.Println(errors.Is(err, target))
	}

	// オブジェクトが違っていてもerrにIsが実装されていればそれを返す
	{
		err := &myErrorImplementedIs{s: "hoge"}
		target := errors.New("hoge")
		fmt.Println(errors.Is(err, target))
	}

	// Unwrapが実装されていれば返り値が比較可能かIsで判定する
	// errのwrapのどれかに同一のオブジェクトか真を返すIsがあれば通る
	{
		target := &myError{s: "hoge"}
		err := fmt.Errorf("%w", target)
		fmt.Println(errors.Is(err, target))
	}

}
