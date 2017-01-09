package main

import (
    "log"
    "encoding/json"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "JEEELOOOU")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/planet/yavin", planets)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

type Planet struct {
    Name string
    Climate string
    Terrain string
}

func planets (w http.ResponseWriter, r *http.Request) {
    //Planet{Name: "Yavin IV", Climate: "temperate, tropical", Terrain: "jungle, rainforests"}
    planet := Planet{Name: "Yavin IV", Climate: "temperate, tropical", Terrain: "jungle, rainforests"}
    body, error := json.Marshal(planet)

    if error != nil {
        panic(error)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(body)
}