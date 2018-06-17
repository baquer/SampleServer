package main

import (
  "fmt"
  "html/template"
  "strings"
  "net/http"
  "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request)
{
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r,Forn["url_lomg"])
  for k , v := range r.Form {
    fmt.Println("key", k)
    fmt.Println("val", strings.Join(v, ""))
  }
  fmt.Println(w, "Hello baquer!")
}

func login(w http.ResponseWriter, r *http.Request) {
  fmt.Println("method:", r.Method)
   if r.Method == "GET" {
     t, _ := template.ParseFiles("login.gtpl")
     t.Execute(w, nil)
   }
   else {
     r.ParseForm()
     fmt.Println("username:", r.Form["username"] )
     fmt.Println("password:", r.Form["password"])
   }

}

func main() {
  http.HandleFunc("/", sayhelloname)
  http.handleFunc("login", login)
  err := http.ListenAndServe("9090", nil)
  if err != nil {
    log.Fatal("ListenAndServe", err)
  }
}
