package main

import (
	"fmt"
	"reflect"
)

func printsomething[K any](num K) {
	fmt.Println("num", num, reflect.TypeOf(num))
}

type number interface {
	int | int32 | int64 | float32 | float64
}

func addnumbers[T number](a, b T) T {
	// Generics with Constraints
	// You can restrict what types are allowed using constraints.
	fmt.Println("num", reflect.TypeOf(a), reflect.TypeOf(b))
	return a + b
}

type Box[T any] struct {
	// Generics with Structs
	value T
}

func Filter[T any](nums []T, f func(T) bool) []T {
	result := []T{}
	for _, val := range nums {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}
func main() {
	printsomething("prasuna")
	printsomething(34)
	printsomething(2.1)
	printsomething('q')
	printsomething(true)
	////////////////////////////////
	fmt.Println("", addnumbers(1, 2))
	fmt.Println("", addnumbers(1.0, 2.0))
	fmt.Println("", addnumbers(1.0, 2))
	///////////////////////////////////
	intbox := Box[int]{value: 10}
	floatbox := Box[float32]{value: 2.1}
	fmt.Println("intbox", intbox.value)
	fmt.Println("floatbox", floatbox.value)
	/////////////////////////////////////////////
	fmt.Println("", Filter([]int{2, 1, 3, 4, 5}, func(n int) bool {
		return n%2 == 0
	}))
	names := []string{"apple", "ass", "bat", "ball", "anar"}
	fmt.Println("", Filter(names, func(n string) bool {
		return n[0] == 'a'
	}))
}
