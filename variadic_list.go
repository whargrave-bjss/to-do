package main

// 10. Create a program using a variadic function to output a list of 10 things To Do to a JSON format file. [Variadic Functions][Structures][JSON]

import ("fmt"
		"encoding/json"
	)
func VariadicList(list ...string) []string {
	for _, n := range list {
		fmt.Printf("%v\n", n)
	}
	return list
}

func VariadicListJSON(list ...string) string {
	tasks := make(map[string][]string)
	tasks["tasks"] = list
	jsonData, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(jsonData)
}

