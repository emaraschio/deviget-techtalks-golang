package main
 
import (
    "encoding/json"
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
type Person struct {
    ID          string      `json:"id,omitempty"`
    Name        string      `json:"firstname,omitempty"`
    Gender      string      `json:"lastname,omitempty"`
    Homeworld   *Homeworld  `json:"homeworld,omitempty"`
}
 
type Homeworld struct {
    Name      string `json:"city,omitempty"`
    Terrain   string `json:"state,omitempty"`
}
 
var people []Person
 
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }

    json.NewEncoder(w).Encode(&Person{})
}
 
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    json.NewEncoder(w).Encode(people)
}
 
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    json.NewEncoder(w).Encode(people)
}
 
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    json.NewEncoder(w).Encode(people)
}

func MyFunc(s string, func(h http.ResponseWriter, w *http.Request)) {
  log.Fatal("test")
}

func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Name: "Wilhuff Tarkin", Gender: "male", Homeworld: &Homeworld{Name: "Eriadu", Terrain: "cityscape"}})
    people = append(people, Person{ID: "2", Name: "Anakin Skywalker", Gender: "male", Homeworld: &Homeworld{Name: "Tatooine", Terrain: "desert"}})
    router.HandleFunc = MyFunc
    /*router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":1337", router))*/
}