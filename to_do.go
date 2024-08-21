package main

// Create a program that prints a list of things To Do 
// and the current status of the To Do item using two goroutines which alternate between 
// To Do Items and To Do statuses [Concurrency][Waitgroups][Workerpools][Mutexes]

import "fmt"
// to-do statuse: complete / incomplete (bool?)
// to-do items list map[string]
// item struct & status struct

type TodoItem struct {
	ID int
	Item string
}

type TodoStatus struct {
	ID int
	Status bool
}

var mockTodoItems = []TodoItem{
    {ID: 1, Item: "Buy fish"},
    {ID: 2, Item: "Make todo list"},
    {ID: 3, Item: "Call dentist"},
    {ID: 4, Item: "Go for a run"},
    {ID: 5, Item: "Read a chapter of book"},
}

var mockTodoStatuses = []TodoStatus{
    {ID: 1, Status: false},
    {ID: 2, Status: true},
    {ID: 3, Status: false},
    {ID: 4, Status: true},
    {ID: 5, Status: false},
}

// items goroutine
// print all items

// status goroutine
// check status of each item
// print status 




func TodoListPrintAllWithStatus(todoItems []TodoItem, todoStatuses []TodoStatus) {

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, item := range todoItems {
			mutex.Lock()
			fmt.Println("Item:", item.Item)
			mutex.Unlock()
		}
	}()
    go func() {
		defer wg.Done()
		for _, status := range todoStatuses {
			mutex.Lock()
			fmt.Println("Status:", status.Status)
			mutex.Unlock()
		}
	}()
	wg.Wait()
}