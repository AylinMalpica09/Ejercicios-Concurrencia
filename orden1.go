package main

import (
	"fmt"
	"sync"
)

func f1 ( wg *sync.WaitGroup){
	fmt.Println("a")
	wg.Done()
}
func f2 ( wg *sync.WaitGroup){
	fmt.Println("b")
	wg.Done()
}
func f3 ( wg *sync.WaitGroup){
	fmt.Println("c")
	wg.Done()
}

// func main(){
// 	//var wg sync.WaitGroup
// 	wg := sync.WaitGroup{}
// 	wg.Add(3)

// 	go f1(&wg)
// 	go f2(&wg)
// 	go f3(&wg)

// 	wg.Wait()
// }