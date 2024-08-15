package main

// 10. Create a program using a variadic function to print a list of 10 things To Do. 

import "fmt"
func VariadicList(list ...string) []string {
	for _, n := range list {
		fmt.Printf("%v\n", n)
	}
	return list
}
