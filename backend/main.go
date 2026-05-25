package main

import (
	"fmt"
	"net/http"
        "starttech/db"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	db.ConnectMongo()

	http.HandleFunc("/health", health)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
