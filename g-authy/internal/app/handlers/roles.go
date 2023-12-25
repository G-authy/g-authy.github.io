package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var roles []Role

/*
func route() {
	router := mux.NewRouter()

	router.HandleFunc("/roles", getRoles).Methods("GET")
	router.HandleFunc("/roles", createRole).Methods("POST")
	router.HandleFunc("/roles/{id}", updateRole).Methods("PUT")
	router.HandleFunc("/roles/{id}", deleteRole).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}*/

func getRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

func createRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var role Role
	_ = json.NewDecoder(r.Body).Decode(&role)
	roles = append(roles, role)
	json.NewEncoder(w).Encode(role)
}

func updateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range roles {
		if item.ID == params["id"] {
			roles = append(roles[:index], roles[index+1:]...)
			var role Role
			_ = json.NewDecoder(r.Body).Decode(&role)
			role.ID = params["id"]
			roles = append(roles, role)
			json.NewEncoder(w).Encode(role)
			return
		}
	}
	json.NewEncoder(w).Encode(&Role{})
}

func deleteRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range roles {
		if item.ID == params["id"] {
			roles = append(roles[:index], roles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(roles)
}
