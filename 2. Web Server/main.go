package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	// return users as json
	_ = []string{"user1", "user2", "user3"}

	// Set the Content-Type header so that the client knows how to interpret the response
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer (w) using "toJSON" function from utils.go
	// ex: toJSON(users) -> ["user1","user2","user3"]

	// fmt.Fprintf(...)

}

func newUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %s", err)
		return
	}

	// Get the "username" form value
	username := r.FormValue("username")

	// Return a success message
	fmt.Fprintf(w, "New user %s added successfully", username)
}

func main() {
	// Register the helloHandler function to handle all requests to the root URL "/"
	http.HandleFunc("/", helloHandler)

	// Register the usersHandler function to handle all requests to the "/users" URL
	http.HandleFunc("/users", usersHandler)

	// Register the newUsersHandler function to handle all requests to the "/new-users" URL (POST request)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
