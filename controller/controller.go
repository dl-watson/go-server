package controller

import (
	"fmt"
	"log"
	"net/http"

	"local.ex/main/pages/about"
	"local.ex/main/pages/home"
)

func Server() {
	port := "7890"
	fmt.Printf("Starting server on port %q...\n", port)

	http.HandleFunc("/", home.Page)
	http.HandleFunc("/about", about.Page)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting the HTTP server : ", err)
		return 
	}
}


