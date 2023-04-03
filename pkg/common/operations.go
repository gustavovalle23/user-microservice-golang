package main

import (
	"fmt"
	"strings"
)

func Map[T, U any](mappingFunc func(T) U, inputSlice []T) []U {
	outputSlice := make([]U, len(inputSlice))

	for i, elem := range inputSlice {
		outputSlice[i] = mappingFunc(elem)
	}

	return outputSlice
}

func Filter[T any](data []T, filterFunc func(T) bool) []T {
	filteredData := make([]T, 0, len(data))

	for _, elem := range data {
		if filterFunc(elem) {
			filteredData = append(filteredData, elem)
		}
	}

	return filteredData
}

func main() {
	input := []int{1, 2, 3, 4}
	result := Map(func(x int) int { return x * 2 }, input)
	fmt.Println(result)

	words := []string{"war", "cup", "water", "tree", "storm"}

	res := Filter(words, func(s string) bool {
		return strings.HasPrefix(s, "w")
	})

	fmt.Println(res)
}
