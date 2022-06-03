package main

import "net/http"
import "os"
import "fmt"
import "io/ioutil"
import "log"

func main() {
    http.HandleFunc("/configmap",  ConfigMap)
    http.HandleFunc("/secret",  Secret)
	http.HandleFunc("/",  Hello)    
	http.ListenAndServe(":8000", nil)
}

 func Hello(w http.ResponseWriter, r *http.Request) {
    name := os.Getenv("NAME")
    age := os.Getenv("AGE")

	fmt.Fprintf(w, "I'm %s and I'm %s years old", name, age)
 }

 func Secret(w http.ResponseWriter, r *http.Request) {
    user := os.Getenv("USER")
    password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "User: %s Password: %s", user, password)
 }

 func ConfigMap(w http.ResponseWriter, r *http.Request) {
    file, err := ioutil.ReadFile("/go/heroes/heroes.txt")

    if err != nil {
        log.Fatalf("Error reading file: ", err)
    }

	fmt.Fprintf(w, "Heroes: %s", string(file))
 }
