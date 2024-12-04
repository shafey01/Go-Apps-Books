package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime/debug"
)

// 1. Logs the error, as well as the stack trace.  \/
// 2. Sets the status code to http.StatusInternalServerError (500)
// whenever a panic occurs. \/
// 3. Write a "Something went wrong" message when a panic occurs. \/
// 4. Ensure that partial writes and 200 headers aren't set even if
// the handler started writing to the http.ResponseWriter BEFORE the panic
// occurred (this one may be trickier) \/
// 5. If the environment is set to be development, print the stack trace and \/
// the error to the webpage as well as to the logs.
// Otherwise default to the "Something went wrong" message described in (3).

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	panicHandler := panicHandler{mux}
	log.Fatal(http.ListenAndServe(":3000", panicHandler))
}

type panicHandler struct {
	fallback http.Handler
}

func (ph panicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {

		if err := recover(); err != nil {
			log.Printf("%v ", err)
			debug.PrintStack()
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Something went wrong!")

			if env, ok := os.LookupEnv("ENV"); ok && env == "dev" {
				fmt.Fprintf(w, string(debug.Stack()))
				fmt.Fprintf(w, "Panic: %v \n", err)

			}
		}

	}()
	rec := httptest.NewRecorder()

	ph.fallback.ServeHTTP(rec, r)

	w.WriteHeader(rec.Code)
	for key, values := range rec.Result().Header {
		for _, v := range values {
			w.Header().Set(key, v)
		}

	}

	rec.Body.WriteTo(w)

}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}
