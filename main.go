package main

import (
	"fmt"
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

func main() {
	http.HandleFunc("/", sayHelloName)       //set router
	err := http.ListenAndServe(":9090", nil) //set listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
