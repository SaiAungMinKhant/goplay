package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int
	fmt.Println(intNum)

	var floatNum float64 = 123213.321
	fmt.Println(floatNum)

	var floatNum32 float32 = 12.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 2
	var intNum2 int = 3
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello \nWorld"
	var myString1 string = `Hello 
	World`
	var myString2 string = "Hello" + " " + "World!"
	fmt.Println(myString)
	fmt.Println(myString1)
	fmt.Println(myString2)
	fmt.Println(utf8.RuneCountInString("Y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune = 3
	fmt.Println(intNum3)

	var myVar = "text"
	myVar1 := "text1"
	fmt.Println(myVar)
	fmt.Println(myVar1)

	var1, var2 := "var1", "var2"
	fmt.Println(var1 + " + " + var2)

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.14
	fmt.Println(pi)

}
