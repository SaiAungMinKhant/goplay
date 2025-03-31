package main

import (
	"fmt"
	"time"
)

func timeLoop(slice []int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice)<n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func main() {

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice2, n))
	// var intArr [3]int32
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3}
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v \n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("\nThe length is %v with capacity %v \n", len(intSlice3), cap(intSlice3))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 21}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Unknown"])
	delete(myMap2, "Adam")
	var age, ok = myMap["Jason"]

	if ok {
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	// for name, age := range myMap2{
	// 	fmt.Printf("Name: %v, Age: %v\n", name, age)
	// }

	// for i, v := range intArr{
	// 	fmt.Printf("Index: %v, Value: %v\n", i, v)
	// }

	// var i int = 0
	// for {
	// 	if i >= 10{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i:=0; i<10; i++ {
	// 	fmt.Println(i)
	// }
}
