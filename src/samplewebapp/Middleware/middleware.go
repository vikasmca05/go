package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/jwk"

	"github.com/dgrijalva/jwt-go"
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

		w.Header().Set("Access-Control-Allow-Origin", "*")
		log.Println("ValidateMiddleware method")

		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			log.Println("auth heder present")

			bearerToken := strings.Split(authorizationHeader, "Bearer ")
			//log.Println(bearerToken)

			// AWS Cognito public keys are available at address:
			// https://cognito-idp.{region}.amazonaws.com/{userPoolId}/.well-known/jwks.json
			publicKeysURL := "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_njGiiHrUB/.well-known/jwks.json"

			// Start with downloading public keys information
			// The .Fetch method is used from https://github.com/lestrrat-go/jwx package
			publicKeySet, err := jwk.Fetch(publicKeysURL)
			if err != nil {
				log.Println("failed to parse key: %s", err)
				log.Printf("failed to parse key: %s", err)
			}

			// Get access token as string from *principal
			// Access token is Base64-encoded JSON that contains user details - called "claims".
			// ---
			// Token is separated into 3 sections - header, payload and signature
			// You can test and validate your token with jwt.io
			tokenString := bearerToken[1]
			log.Println(tokenString)
			// We want to get details from the access token: client_id and unique user identifier.
			// Let's add client_id. We can verify, if it match our App cliet ID in AWS Cognito User Pool
			// We can also add user identifier (f.e. "username") to use it with our App
			type AWSCognitoClaims struct {
				Client_ID string   `json:"email"`
				Username  string   `json:"cognito:username"`
				Groups    []string `json:"cognito:groups"`
				jwt.StandardClaims
			}

			// JWT Parse - it's actually doing parsing, validation and returns back a token.
			// Use .Parse or .ParseWithClaims methods from https://github.com/dgrijalva/jwt-go
			token, err := jwt.ParseWithClaims(tokenString, &AWSCognitoClaims{}, func(token *jwt.Token) (interface{}, error) {

				// Verify if the token was signed with correct signing method
				// AWS Cognito is using RSA256 in my case
				_, ok := token.Method.(*jwt.SigningMethodRSA)

				if !ok {
					log.Println("Unexpected signing method: %v", token.Header["alg"])

					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				log.Println("signing method passed")

				// Get "kid" value from token header
				// "kid" is shorthand for Key ID
				kid, ok := token.Header["kid"].(string)
				if !ok {
					log.Println("Kid header not found")

					return nil, errors.New("kid header not found")
				}
				log.Println("kid header value = ")
				log.Println(kid)

				// Check client_id attribute from the access token
				claims, ok := token.Claims.(*AWSCognitoClaims)
				if !ok {
					log.Println("Problem to get claims")

					return nil, errors.New("There is problem to get claims")
				}
				log.Printf("client_id: %v", claims.Client_ID)
				log.Println("Claims Value - Expirest ")
				log.Println((claims.StandardClaims.ExpiresAt))
				log.Println("Claims Value - Audience ")
				log.Println((claims.StandardClaims.Audience))
				log.Println("Claims Value - Issuer ")
				log.Println((claims.StandardClaims.Issuer))
				log.Println("Claims Value - Subject ")
				log.Println((claims.StandardClaims.Subject))
				log.Println("Claims Value - Username ")
				log.Println((claims.Username))
				log.Println("Claims Value - Eail ")
				log.Println(claims.Client_ID)
				log.Println("Claims Value - Groups ")
				log.Println(claims.Groups[0])

				//w.Write([]byte(fmt.Sprintf("Welcome # %s !", claims.Username)))

				// "kid" must be present in the public keys set
				keys := publicKeySet.LookupKeyID(kid)
				if len(keys) == 0 {
					log.Println("Key not found")

					return nil, fmt.Errorf("key %v not found", kid)
				}
				log.Println("Key found")

				// In our case, we are returning only one key = keys[0]
				// Return token key as []byte{string} type
				var tokenKey interface{}
				if err := keys[0].Raw(&tokenKey); err != nil {
					return nil, errors.New("failed to create token key")
				}

				return tokenKey, nil
			})

			if err != nil {
				// This place can throw expiration error
				log.Printf("token problem: %s", err)
			}

			// Check if token is valid
			if !token.Valid {
				log.Println("token is invalid")
			}

		}
		next(w, req)
	})
}
