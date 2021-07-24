package main

import (
  "fmt"
  "log"
  "net/http"
)

// Go built-in types as required by function handlers
func home (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Home!")
}

func main () {
	// First argument is path, second argument is callback
	http.HandleFunc("/", home)
	fmt.Println("Server is running on port :8080")
	// Second argument is kind of optional
	// The ListenAndServe is always executed.
	// If something goes wrong, it will return an error which then gets
	// handled by log.Fatal()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
