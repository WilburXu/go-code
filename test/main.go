package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello")
	})

	go func() {
		if err := http.ListenAndServe(":8811", nil); err != nil {
			log.Println(err)
		}
	}()

	select {}
}
