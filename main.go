package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello! welcome")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s %s\n ", firstname, lastname)
	fmt.Fprintf(w, "Phone = %s\n", phone)
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Address = %s\n", address)
}
