package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vsrrevanth/gocrud/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/student", api.GetStudents).Methods("GET")
	r.HandleFunc("/student/{id}", api.GetStudentById).Methods("GET")
	r.HandleFunc("/student", api.CreateStudent).Methods("POST")
	r.HandleFunc("/student/{id}", api.UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/{id}", api.DeleteStudent).Methods("DELETE")

	fmt.Println("Starting Server on Port 8000")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
