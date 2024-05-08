package main

import (
	"errors"
	"fmt"

	m "github.com/Hunnnn77/monad/src"
)

func main() {
	r := Test(2).Unwrap()
	fmt.Println(r)

	r2 := Test(1).UnwrapOr(func(e error) int {
		return -1
	})
	fmt.Println(r2)

	r3 := Test(1).Match(func(value int) *int {
		return &value
	}, func(e error) *int {
		fb := -1
		return &fb
	})
	fmt.Println(*r3)
}

func Test(i int) m.Result[int] {
	if i == 1 {
		return m.Result[int]{
			Err: &m.Err{E: errors.New("err")},
		}
	}
	return m.Result[int]{
		Ok: &m.Ok[int]{
			Value: 1,
		},
	}
}
