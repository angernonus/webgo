package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func reqLog(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, r.Method)
	log.Println(r.URL.Path, r.Method)
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	reqLog(w, r)
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!")
}

func login(w http.ResponseWriter, r *http.Request) {
	reqLog(w, r)
	if r.Method == "GET" {
		r.ParseForm()
		for k, v := range r.Form {
			if k == "failed" && v[0] == "true" {
				fmt.Fprintf(w, "wrong username/passoword")
			}
		}
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		if len(r.Form["username"]) == 0 {
			http.Redirect(w, r, "http://localhost:8080/login?failed=true", 400)
		} else if len(r.Form["passoword"]) == 0 {
			http.Redirect(w, r, "http://localhost:8080/login?failed=true", 400)
		} else {
			fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
			fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("passoword")))
			fmt.Fprintf(w, "welcome, ")
			template.HTMLEscape(w, []byte(r.Form.Get("username"))) // printed out after form has been completed
		}
	}
}

func main() {
	// set up logfile
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening logfile: %v", err)
	}
	defer f.Close()
	log.SetOutput(f) // makes the logfile the default log

	http.HandleFunc("/", sayHi)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
