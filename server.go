package main

import "net/http"
import "os"
import "fmt"

func main() {
	http.HandleFunc("/",  Hello)
	http.ListenAndServe(":8000", nil)
}

 func Hello(w http.ResponseWriter, r *http.Request) {
    name := os.Getenv("NAME")
    age := os.Getenv("AGE")

	fmt.Fprintf(w, "I'm %s and I'm %s years old", name, age)
 }
