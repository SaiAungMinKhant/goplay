package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice(float32Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}

// func sumIntSlice(slice []int) int{
// 	var sum int
// 	for _, v := range slice{
// 		sum += v
// 	}
// 	return sum
// }

// func sumFloat32Slice(slice []float32) float32{
// 	var sum float32
// 	for _, v := range slice{
// 		sum += v
// 	}
// 	return sum
// }

// func sumFloat64Slice(slice []float64) float64{
// 	var sum float64
// 	for _, v := range slice{
// 		sum += v
// 	}
// 	return sum
// }