package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "<h1><font color='blue'>Hello world: MAU 2018/06/08</font></h1>")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":80", nil)
}
