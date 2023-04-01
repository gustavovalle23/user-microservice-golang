package main

import "fmt"

func Map[T, U any](fn func(T) U, slice []T) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 4}
	result := Map(func(x int) int { return x * 2 }, input)
	fmt.Println(result)
}
