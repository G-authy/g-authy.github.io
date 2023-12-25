package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var groups []Group

/*
func route() {
	router := mux.NewRouter()

	router.HandleFunc("/groups", getGroups).Methods("GET")
	router.HandleFunc("/groups", createGroup).Methods("POST")
	router.HandleFunc("/groups/{id}", updateGroup).Methods("PUT")
	router.HandleFunc("/groups/{id}", deleteGroup).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}*/

func getGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var group Group
	_ = json.NewDecoder(r.Body).Decode(&group)
	groups = append(groups, group)
	json.NewEncoder(w).Encode(group)
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range groups {
		if item.ID == params["id"] {
			groups = append(groups[:index], groups[index+1:]...)
			var group Group
			_ = json.NewDecoder(r.Body).Decode(&group)
			group.ID = params["id"]
			groups = append(groups, group)
			json.NewEncoder(w).Encode(group)
			return
		}
	}
	json.NewEncoder(w).Encode(&Group{})
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range groups {
		if item.ID == params["id"] {
			groups = append(groups[:index], groups[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(groups)
}
