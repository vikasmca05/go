package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "samplewebapp/Models"
)

// GetAllTask get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	getAllTask()
	//payload := getAllTask()
	//json.NewEncoder(w).Encode(payload)
}

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.Subscription
	_ = json.NewDecoder(r.Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

// get all task from the DB and return it
func getAllTask() {
	fmt.Println("Get all Records ")
}

// Insert one task in the DB
func insertOneTask(task models.Subscription) {

	fmt.Println("Inserted a Single Record ")
}

