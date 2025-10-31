package main

import (
	"fmt"
	"net/http"
	"time"
)


func main()  {
		fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":4000", nil)
}

var tasks = []string{
	"Task 1: Buy groceries",
	"Task 2: Walk the dog",
	"Task 3: Read a book",
}

func showTasks(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintln(writer, "Showing all tasks...")
	for _, task := range tasks {
		fmt.Fprintln(writer, task)
	}
}