package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hey", handleSegment)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"
	log.Printf("Listening on port %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleSegment(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	writeResponse(writer, "Segmenting!")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	writeResponse(writer, "Server Status Ok")
}

func writeResponse(writer http.ResponseWriter, responseStr string) {
	response := []byte(responseStr)
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
