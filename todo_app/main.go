package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, "<h1>todo app</h1><p>Server running on port "+port+"</p>")
	})

	fmt.Printf("Server started in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
