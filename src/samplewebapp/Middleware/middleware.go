package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Subscription struct
type Subscription struct {
	Product string `json:"product"`
	Type    string `json:"type"`
}

var subs []Subscription

// GetAllTask get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	getAllTask(w)
	//payload := getAllTask()
	//json.NewEncoder(w).Encode(payload)
}

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task Subscription
	_ = json.NewDecoder(r.Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneTask(w, r, task)
	json.NewEncoder(w).Encode(task)
}

// get all task from the DB and return it
func getAllTask(w http.ResponseWriter) {
	fmt.Println("Get all Records ")

	//Convert the "subs" variable to json
	subsListBytes, err := json.Marshal(subs)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of subs to the response
	w.Write(subsListBytes)
}

// Insert one task in the DB
func insertOneTask(w http.ResponseWriter, r *http.Request, task Subscription) {
	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sub := Subscription{}

	sub.Product = r.Form.Get("Task")
	sub.Type = "Test task"

	subs = append(subs, sub)
	fmt.Println("Inserted a Single Record ")
}
