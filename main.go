package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Car struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Year   string `json:"Year"`
	S_type string `json:"S_type"`
}

var Models []Car

func main() {

	Models = []Car{
		Car{Id: "1", Name: "Tesla", Year: "2020", S_type: "Robot"},
		Car{Id: "2", Name: "Volvo", Year: "2021", S_type: "Automation"},
		Car{Id: "3", Name: "BMW", Year: "2022", S_type: "Mechanic"},
		Car{Id: "4", Name: "Cherry", Year: "2023", S_type: "Robot"},
		Car{Id: "5", Name: "Audi", Year: "2024", S_type: "Robot"},
	}
	apiRequest()
}

func apiRequest() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", homePage)
	route.HandleFunc("/car/{id}", getCar)
	log.Fatal(http.ListenAndServe(":3000", route))
}

func getCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, cars := range Models {
		if cars.Id == key {
			json.NewEncoder(w).Encode(cars)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API Home page!")
}
