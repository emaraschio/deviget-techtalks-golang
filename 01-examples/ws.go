package main

import(
  "net/http"
  "log"
  "fmt"
)

func main() {
  http.HandleFunc("/mypage", handler)
  http.HandleFunc("/", defaultPage)
  log.Println("listening on port 1337...")
  log.Fatal(http.ListenAndServe("localhost:1337", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "MY PAGEEEE JELOOOU")
}

func defaultPage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "MY ROOT")
}
