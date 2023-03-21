package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:         "1",
		TaskName:   "New project",
		TaskDetail: "Finish the project before deadline",
		Date:       "2023-03-20",
	}

	tasks = append(tasks, task)

	task1 := Tasks{
		ID:         "2",
		TaskName:   "Power project",
		TaskDetail: "The delete function needs to be fixed",
		Date:       "2023-03-21",
	}

	tasks = append(tasks, task1)

	fmt.Println("your tasks are", tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the HomePage")
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	allTasks()
	handleRoutes()
}
