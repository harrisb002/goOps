package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hey", handleHey)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"
	log.Printf("Listening on port %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHey(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		// Pass 405 status code to Error handler to map to string
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(writer, "Hey There!")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(writer, "Ok")

}

func writeResponse(writer http.ResponseWriter, responseStr string) {
	response := []byte(responseStr)
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
