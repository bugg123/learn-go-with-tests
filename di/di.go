package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(myGreeterHandler))
}

// Greet will print hello name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func myGreeterHandler(w http.ResponseWriter, r *http.Request){
	Greet(w, "world");
}