package main

import (
	"fmt"
	"sync"
	"time"
)

// var m = sync.Mutex{}
// var m = sync.RWMutex{}

var wg = sync.WaitGroup{}
// var dbData = []string{"id1","id2","id3","id4","id5","id6"}
// var results = []string{}
func main(){
	t0 := time.Now()
	for i:=0; i<1000; i++{
		wg.Add(1)
		// go dbCall(i)
		go count()
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	// fmt.Printf("\nThe results are: %v", results)
}

func count(){
	var res int
	for i:=0; i<1000000; i++{
		res+=1
	}
	wg.Done()
}

// func dbCall(i int){
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay)*time.Millisecond)

// 	wg.Done()
// }

// func save(result string){
// 	m.Lock()
// 	results = append(results, result)
// 	m.Unlock()
// }

// func log(){
// 	m.RLock()
// 	fmt.Printf("\nThe current results are: %v", results)
// 	m.RUnlock()
// }