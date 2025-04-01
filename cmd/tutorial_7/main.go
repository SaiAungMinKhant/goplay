package main

import "fmt"

func main() {
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of thing1 array is %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is %v", result)
	fmt.Printf("\nThe value of thing1 is %v", thing1)
}

func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("\nThe memory location of thing2 array is %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

// func main() {
// var slice = []int32{1,2,3}
// var sliceCopy = slice
// sliceCopy[2] = 4
// fmt.Println(slice)
// fmt.Println(sliceCopy)
// 	var p *int32 = new(int32)
// 	var i int32
// 	fmt.Printf("The value of p points to is: %v\n", *p)
// 	fmt.Printf("The value of i is: %v\n", i)
// 	p = &i
// 	*p = 12
// 	fmt.Printf("The value of p points to is: %v\n", *p)
// 	fmt.Printf("The value of i is: %v", i)
// }