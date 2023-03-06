package main

import (
  "fmt"
  "net/http"
  "html"
  "log"
)

func main() {
  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":8081", nil))
}
