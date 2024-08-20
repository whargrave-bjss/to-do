package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TodoList struct {
	Todos []string `json:"todos"`
}
func VariadicList(list ...string) []string {
	for _, n := range list {
		fmt.Printf("%v\n", n)
	}
	return list
}

func VariadicListToJSON(list ...string) string {
	tasks := make(map[string][]string)
	tasks["tasks"] = list
	jsonData, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(jsonData)
}

func VariadicJSONFile() {
	jsonData, err := os.ReadFile("todo.json")
	if err != nil {
		fmt.Println("Error reading file:",err)
		return
	}

	var todoList TodoList

	err = json.Unmarshal(jsonData, &todoList) 
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Your Todo list:")
	for i, todo := range todoList.Todos {
		fmt.Printf("%d, %s\n:", i+1, todo)
	}


}
