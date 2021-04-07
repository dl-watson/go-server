package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Models
type Villager struct {
	ID string `json:"_id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Species string `json:"species"`
	Personality string `json:"personality"`
	Birthday string `json:"birthday"`
	Quote string `json:"quote"`
}

// CRUD handler funcs
func getVillagers(w http.ResponseWriter, r *http.Request) {
	ACNHClient := http.Client{ Timeout: time.Second * 10 }
	url := "https://ac-vill.herokuapp.com/villagers?perPage=397"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := ACNHClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var villagers []Villager
	jsonErr := json.Unmarshal(body, &villagers)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villagers)
}

func getVillager(w http.ResponseWriter, r *http.Request) {
	ACNHClient := http.Client{ Timeout: time.Second * 10 }
	params := mux.Vars(r)
	url := "https://ac-vill.herokuapp.com/villagers?name=" + params["name"]

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := ACNHClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var villager []Villager
	jsonErr := json.Unmarshal(body, &villager)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villager)
}


func Controller() {
	port := "7890"
	fmt.Printf("Starting server on port %q...\n", port)

	// Initialize router
	r := mux.NewRouter()
	
	// Route handlers / endpoints
	r.HandleFunc("/api/villagers", getVillagers).Methods("GET")
	r.HandleFunc("/api/villagers/{name}", getVillager).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":"+port, r))
}
