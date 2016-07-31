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
        fmt.Println(r.URL.Path, r.Method, r.Header.Get("X-Forwarded-For"))
        log.Println(r.URL.Path, r.Method, r.Header.Get("X-Forwarded-For"))
}

func sayHi(w http.ResponseWriter, r *http.Request) {
    reqLog(w,r)
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
    r.ParseForm()
    if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("passoword")))
		fmt.Fprintf(w, "welcome, ")
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // printed out after form has been completed
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
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
