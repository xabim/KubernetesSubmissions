package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var counter atomic.Int64

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong %d\n", counter.Add(1)-1)
	})

	fmt.Printf("Server started in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
