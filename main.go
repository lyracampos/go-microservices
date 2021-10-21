package main

import (
	"log"
	"net/http"
)

func welcome(http.ResponseWriter, *http.Request) {
	log.Println("welcome to golang api")
}

func goodBye(http.ResponseWriter, *http.Request) {
	log.Println("golang api saing good bye...")
}

func handlerRequests() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/goodbye", goodBye)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
func main() {
	handlerRequests()
}
