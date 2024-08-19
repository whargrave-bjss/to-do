package main

import (
	"testing"
	"reflect"
	"os"
	"bytes"
	"io"
)

func TestVariadicList(t *testing.T) {
	
	got := VariadicList("task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10")
	want := []string{"task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10"}

	if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func TestVariadicListToJSON(t *testing.T) { 

	got := VariadicListToJSON("task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10")
	want := "{\n    \"tasks\": [\n        \"task1\",\n        \"task2\",\n        \"task3\",\n        \"task4\",\n        \"task5\",\n        \"task6\",\n        \"task7\",\n        \"task8\",\n        \"task9\",\n        \"task10\"\n    ]\n}"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestVariadicJSONFile(t *testing.T) {

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	VariadicJSONFile()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	got := buf.String()
	want := "Your Todo list:\n" +
        "1, Buy groceries\n" +
        ":2, Do washing\n" +
        ":3, return book to libray\n" +
        ":4, Finish report\n" +
        ":5, Go to gym\n" +
        ":6, Read a chapter\n" +
        ":7, Water plants\n" +
        ":8, Pay bills\n" +
        ":9, Clean kitchen\n" +
        ":10, Prepare presentation\n" +
        ":"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}	
	
}