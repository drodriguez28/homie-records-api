package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Homie struct {
	ID    string `json:"id"`
	Age   string `json:"age"`
	City  string `json:"city"`
	Alias string `json:"alias"`
	Name  *Name  `json:"name"`
}

type Name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var homies []Homie

func getHomies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(homies)
}

func deleteHomie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range homies {

		if item.ID == params["id"] {
			homies = append(homies[:index], homies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(homies)
}

func getHomie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range homies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createHomie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var homie Homie
	_ = json.NewDecoder(r.Body).Decode(&homie)
	homie.ID = strconv.Itoa(rand.Intn(10000000))
	homies = append(homies, homie)
	json.NewEncoder(w).Encode(homie)
}

func updateHomie(w http.ResponseWriter, r *http.Request) {
	//set json
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range homies {

		if item.ID == params["id"] {
			homies = append(homies[:index], homies[index+1:]...)
			var homie Homie
			_ = json.NewDecoder(r.Body).Decode(&homie)
			homie.ID = params["id"]
			homies = append(homies, homie)
			json.NewEncoder(w).Encode(homies)

		}
	}

}

func main() {
	r := mux.NewRouter()

	homies = append(homies, Homie{ID: "1", Age: "23", City: "Roswell", Alias: "cepi", Name: &Name{Firstname: "Deivid", Lastname: "Rodriguez"}})
	homies = append(homies, Homie{ID: "2", Age: "18", City: "Roswell", Alias: "tex", Name: &Name{Firstname: "Alexis", Lastname: "Rodriguez"}})

	r.HandleFunc("/homies", getHomies).Methods("GET")
	r.HandleFunc("/homies/{id}", getHomie).Methods("GET")
	r.HandleFunc("/homies", createHomie).Methods("POST")
	r.HandleFunc("/homies/{id}", updateHomie).Methods("PUT")
	r.HandleFunc("/homies/{id}", deleteHomie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
