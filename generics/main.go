package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sliceString := []string{"a", "b", "c", "d", "e"}
	newInts := reverse(slice)
	newStrings := reverse(sliceString)
	fmt.Println(newInts)
	fmt.Println(newStrings)
}

type constraintsCustom interface {
	int | string
}

func reverse[T constraintsCustom](slice []T) []T {
	newSlice := make([]T, len(slice))

	newSliceLen := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newSlice[newSliceLen] = slice[i]
		newSliceLen--
	}
	return newSlice

}
