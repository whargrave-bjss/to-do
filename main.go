package main

import (
		"time" 
		"math/rand"
	)

func main() {
	// VariadicList("task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10")
	// VariadicListToJSON("task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10")
	// VariadicJSONFile()
	rand.NewSource(time.Now().UnixNano())
	EvenOddSum()
}