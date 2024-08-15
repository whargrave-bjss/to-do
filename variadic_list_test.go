package main

import (
	"testing"
	"reflect"
)

func VariadicListTest(t *testing.T) {
	
	got := VariadicList("task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10")
	want := []string{"task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10"}

	if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func VariadicListJSONTest(t *testing.T) { 

	got := VariadicListJSON("task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10")
	want := "{\n    \"tasks\": [\n        \"task1\",\n        \"task2\",\n        \"task3\",\n        \"task4\",\n        \"task5\",\n        \"task6\",\n        \"task7\",\n        \"task8\",\n        \"task9\",\n        \"task10\"\n    ]\n}"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}