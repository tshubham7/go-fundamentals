package main

import (
	"fmt"
	"sort"
)

func main() {
}

// sort the input array/slice
func sortIntegerSlice() {
	input := []int{2, 4, 9, 6, 7, 2, 1, 5, 3, 6, 10, 15}

	// // simple way of sorting int slice
	// sort.Ints(input)

	sort.Slice(input, func(i, j int) bool {
		// return input[i] > input[j] // descreasing order
		return input[i] < input[j] // increasing order
	})
	fmt.Println(input)
}

type student struct {
	Name  string
	Marks int
}

func sortStudent() {
	// sorting a slice of structs
	input := []student{
		{"jimenez", 50},
		{"arda", 60},
		{"morata", 40},
		{"miranda", 33},
		{"koke", 74},
		{"philippe", 87},
	}
	sort.Slice(input, func(i, j int) bool {
		// return input[i].Marks < input[j].Marks // sort by marks
		return input[i].Name < input[j].Name // sort by name
	})
	fmt.Println(input)
}
