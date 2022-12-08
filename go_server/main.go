package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST requesr successful\n")
	name := r.FormValue("name")
	quote := r.FormValue("quote")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "QUote of the day is -\n %s\n", quote)
}

func helloHandler (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080.\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}