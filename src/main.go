package main

import (
	"encoding/json"
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

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			break
		}
		if flag == false {
			json.NewEncoder(w).Encode(map[string]string{"status": "error"})
		}
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This is the HomePage")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the HomePage")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the HomePage")
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	handleRoutes()
}
