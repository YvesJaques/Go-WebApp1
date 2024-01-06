package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {
		nb, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes Written: %d", nb)
	})
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err)
}
