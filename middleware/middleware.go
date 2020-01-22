package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", timer(hello))
	http.ListenAndServe(":1234", nil)
}
func timer(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		fmt.Println("The request took", end.Sub(start))
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi everyone")
}
