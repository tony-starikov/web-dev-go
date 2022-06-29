package main

import (
	"fmt"
	"net/http"
)

func main() {
	var router Router
	fmt.Println("Starting the server on port 3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}

type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		errorHandler(w, r)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:tony.starikov.1992@gmail.com\">tony.starikov.1992@gmail.com</a></p>")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome!</h1>")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 page not found</h1>")
}
