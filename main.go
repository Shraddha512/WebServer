package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//Check token validity
		} else {

		}
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("Empty Username. Please enter your Username")
		} else {
			fmt.Println("username:", r.Form["username"])
		}
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
