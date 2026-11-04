package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	id := rand.Text()

	for {
		fmt.Printf("%s: %s\n", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"), id)
		time.Sleep(5 * time.Second)
	}
}
