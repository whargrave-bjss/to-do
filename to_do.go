package main

// Create a program that prints a list of things To Do 
// and the current status of the To Do item using two goroutines which alternate between 
// To Do Items and To Do statuses [Concurrency][Waitgroups][Workerpools][Mutexes]


// to-do statuse: complete / incomplete (bool?)
// to-do items list map[string]
// item struct & status struct

type todoItem struct {
	ID int
	item string
}

type todoStatus struct {
	ID int
	status bool
}


