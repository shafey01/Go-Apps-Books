package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go Web %s", r.URL.Path[1:])
}
func handlerHeader(w http.ResponseWriter, r *http.Request) {
	s := r.Header.Clone()
	for i, v := range s {
		fmt.Fprintf(w, "headers are: %s : %s \n", i, v)
	}
}
func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/string", handlerHeader)
	http.ListenAndServe(":8080", nil)

}
