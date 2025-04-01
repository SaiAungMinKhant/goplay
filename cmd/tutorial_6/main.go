package main

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
	// owner
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

// type owner struct{
// 	name string
// }

func (e gasEngine) milesLeft() uint8{
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8{
	return e.mpkwh * e.kwh
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need to feul up first!")
	}
} 

func main(){
	var myEngine gasEngine = gasEngine{12, 4}
	var myElectricEngine electricEngine = electricEngine{10, 15}
	// var myEngine gasEngine = gasEngine{12, 4, owner{"Alex"}}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in the tank: %v\n", myEngine.milesLeft())
	// canMakeIt(myEngine, 50)
	canMakeIt(myEngine, 40)
	fmt.Printf("Total miles left in the battery: %v\n", myElectricEngine.milesLeft())
	canMakeIt(myElectricEngine, 200)
}