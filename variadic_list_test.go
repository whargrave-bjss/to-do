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