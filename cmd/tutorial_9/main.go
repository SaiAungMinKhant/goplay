package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 6
var MAX_TOFU_PRICE float32 = 3

func main(){
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites  = []string{"walmart.com","costco.com","wholefoods.com"}
	for i:= range websites{
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string){
	for{
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice<=MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, c chan string){
	for{
		time.Sleep(time.Second*1)
		var TofuPrice = rand.Float32()*20
		if TofuPrice<=MAX_TOFU_PRICE{
			c <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string){
	select{
		case website := <-chickenChannel:
			fmt.Printf("\nText Sent: Found deal on chicken at %s", website)
		case website := <-tofuChannel:
			fmt.Printf("\nText Sent: Found deal on tofu at %s", website)
	}
	
	// fmt.Printf("\nFound a deal on chicken at %s", <-chickenChannel)
	
}

// func main(){
// 	var c = make(chan int)
// 	go process(c)
// 	for i:= range c{
// 		fmt.Println(i)
// 		time.Sleep(time.Second*1) // some work
// 	}
	
// 	// c <- 1
// 	// var i = <- c
// }

// func process(c chan int){
// 	defer close(c)
// 	for i:=0 ; i<5; i++{
// 		c <- i
// 	}
// 	fmt.Println("Exiting Process")
// 	// close(c)
// }