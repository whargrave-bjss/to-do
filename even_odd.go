package main

import (
	"math/rand"
	"time"
	"fmt"
)
// Write a program to simulate a race condition occurring when one goroutine updates a data variable with odd numbers, 
// while another updates the same data variable with even numbers. After each update , 
// attempt to display the data contained in the data variable to screen. [Goroutines][Concurrency][Race Conditions]

// data variable

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomNumber() int {
	return rng.Intn(10) + 1
}
var data int = 0

func EvenOddSum() {
	for i := 0; i < 10; i++ {
		number := randomNumber()
		fmt.Printf("Generated number: %d\n", number)
		time.Sleep(1 * time.Second)
		if number%2 == 0 {
			go func() {
			data += number
		    fmt.Printf("Even: %d Added to make: %d\n", number, data)
			}()
		} else {
			go func() {
			data += number
			fmt.Printf("Odd: %d Added to make: %d\n", number, data)
			}()
		}
	}
}