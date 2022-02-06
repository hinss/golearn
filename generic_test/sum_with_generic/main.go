package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var result V
	for _,v := range m {
		result += v
	}

	return result
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var result V
	for _,v := range m {
		result += v
		}

	return result
}


func main() {

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first": 34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first": 35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
