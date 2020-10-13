package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

//Subscription struct
type Subscription struct {
	Product string `json:"product"`
	Type    string `json:"type"`
}

var subs []Subscription

// OptionsTask get all the task route
func OptionsTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	getAllTask(w)
	//payload := getAllTask()
	//json.NewEncoder(w).Encode(payload)
}

// GetAllTask get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var task Subscription
	_ = json.NewDecoder(r.Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneTask(w, r, task)
	json.NewEncoder(w).Encode(task)
}

// get all task from the DB and return it
func getAllTask(w http.ResponseWriter) {
	fmt.Println("Get all Records ")

	subs := Subscription{
		Product: "test",
		Type:    "type-1",
	}
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

// ValidateMiddleware method call
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					//json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					//json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			//json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}
