package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello Wolrd!"
	printMe(printValue)
	int1 := 10
	int2 := 0
	var result, remainder, err = intDivision(int1, int2)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	// }

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(int1 int, int2 int) (int, int, error) {
	var err error
	if int2 == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = int1 / int2
	var remainder int = int1 % int2
	return result, remainder, err
}
