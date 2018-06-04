package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "<h1>Hello world: MAU 2018/06/08")
}

func main() {
  http.Handle("/", handler)
  http.ListenAndServe(":80", nil)
}
