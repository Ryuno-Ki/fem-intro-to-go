package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// A StarWars Planet
type Planet struct {
  // JSON not XML
  Name string `json:"name"`
  Population string `json:"population"`
  Terrain string `json:"terrain"`
}

// A StarWars character
type Person struct {
  Name string `json:"name"`
  Homeworld Planet
  HomeworldURL string `json:"homeworld"`
}

// All StarWars characters
type AllPeople struct {
  // Alias JSON results field to People
  People []Person `json:"results"`
}

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

func (p *Person) getHomeworld () {
  res, err := http.Get(p.HomeworldURL)

  if err != nil {
    log.Print("Error fetching homeworld,", err)
  }

  // Make bytes available outside of if block
  var bytes []byte
  if bytes, err = ioutil.ReadAll(res.Body); err != nil {
    log.Print("Error parsing homeworld,", err)
  }

  json.Unmarshal(bytes, &p.Homeworld)
}

func getPeople (w http.ResponseWriter, r *http.Request) {
  response, err := http.Get(fmt.Sprintf("%speople", BaseURL))

  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    log.Print("Failed to request Star Wars people")
  }

  bytes, err := ioutil.ReadAll(response.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    log.Print("Failed to parse request body")
  }

  fmt.Println(string(bytes))

  var people AllPeople

  // Parse bytes as JSON and store the result in people
  if err = json.Unmarshal(bytes, &people); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    fmt.Println("Failed to parse JSON", err)
  }

  fmt.Println(people)
  // fmt.Fprintf(w, people)

  for _, pers := range people.People{
    pers.getHomeworld()
    fmt.Println(pers)
  }
}

func main() {
	fmt.Println(BaseURL)
	http.HandleFunc("/people/", getPeople)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
