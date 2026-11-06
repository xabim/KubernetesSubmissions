package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func timestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
}

func main() {
	id := rand.Text()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	go func() {
		for {
			fmt.Printf("%s: %s\n", timestamp(), id)
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s: %s\n", timestamp(), id)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
