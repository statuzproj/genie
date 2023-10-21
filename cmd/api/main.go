package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("genie is out of the jar")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Genie says Hello World!")
		if err != nil {
			return
		}
		log.Println("responded with hello world")
	})
	http.ListenAndServe(":8081", nil)
}
