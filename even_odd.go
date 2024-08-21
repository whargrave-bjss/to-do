package main

import (
	"math/rand"
	"time"
	"fmt"
	"sync"
)
// Write a program to simulate a race condition occurring when one goroutine updates a data variable with odd numbers, 
// while another updates the same data variable with even numbers. After each update , 
// attempt to display the data contained in the data variable to screen. [Goroutines][Concurrency][Race Conditions]

// data variable

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
 
func randomNumber() int {
	return rng.Intn(10) + 1
}

var mutex = &sync.Mutex{}
func EvenOddSum() {
	data := make(map[string]int)
	dataChannel := make(chan int, 50)
	var wg = sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		number := randomNumber()
		wg.Add(1)
		if number%2 == 0 {
			go func(num int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			data["even"] += num
			dataChannel <- num
		    fmt.Printf("Even: %d Added to map: %v\n", num, data)
			}(number)
		} else {	
				go func(num int) {
				defer wg.Done()
				mutex.Lock()
				defer mutex.Unlock()
				data["Odd"] += num
				dataChannel <- num
				fmt.Printf("Odd: %d Added to map: %v\n", num, data)
				}(number)
		}
	}
	wg.Wait()
	close(dataChannel)
}