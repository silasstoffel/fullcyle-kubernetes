package main

import "net/http"

func main() {
	http.HandleFunc("/",  Hello)
	http.ListenAndServe(":8000", nil)
}

 func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Kubernetes</h1><p>v1.1</p>"))
 }
