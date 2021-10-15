package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Zdravo ljubitelji Petsona i Findusa!</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Pettson je stari farmer koji živi u neuređenoj falu crveno obojenoj drvenoj farmi u švedskom selu.</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)

}
