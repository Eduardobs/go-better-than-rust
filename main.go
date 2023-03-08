package main

import (
	"fmt"
	"net/http"
	"sync"
)

var counter int
var mutex sync.Mutex

func incrementCounter() int {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	return counter
}

func getCurrentCounter() int {
	mutex.Lock()
	defer mutex.Unlock()
	return counter
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Fprintf(w, "%d", incrementCounter())
		} else {
			fmt.Fprintf(w, "%d", getCurrentCounter())
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
