package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getPizzasHandler(w http.ResponseWriter, r *http.Request) {
	pizzas := []string{"Margherita", "Pepperoni", "Four Cheese", "Vegetarian", "Neapolitan"}
	fmt.Fprintf(w, "%s", pizzas)
}

func main() {
	// Register the helloHandler function to handle all requests to the root URL "/"
	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/pizzas", getPizzasHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
