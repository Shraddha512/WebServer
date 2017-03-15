package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) //prints server side information
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //sends data to client side

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request Method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}

}

func main() {
	http.HandleFunc("/", sayHelloName) //set router
	http.HandleFunc("/login", login)
	// nil --> DefaultServeMux, router variable which can cal handler functions fo specified URLs
	err := http.ListenAndServe(":9090", nil) //set listening port and initialize a server object
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
