package main

// import (
// 	"fmt"
// 	"sync"
// )

// var total int = 0
// var mutex = &sync.Mutex{}

// func sumar(wg *sync.WaitGroup) {
// 	mutex.Lock()
// 	for i := 1; i <= 215000; i++ {
// 		total++
// 	}
// 	mutex.Unlock()
// 	wg.Done()
// }
// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)

// 	go sumar(&wg)
// 	go sumar(&wg)
// 	wg.Wait()
// 	fmt.Println(total)
// }
