package main

import (
    "strings"
    "fmt"
    "net/http"
    "html/template"
    "log"
)

func sayHi(w http.ResponseWriter, r *http.Request) {
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
    fmt.Println("method:", r.Method) // request method ex. GET PUT POST
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
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
    http.HandleFunc("/", sayHi)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
