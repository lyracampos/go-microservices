package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("welcome to golang api")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "welcome %s", d)
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
