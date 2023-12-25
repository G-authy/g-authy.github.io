package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Permission struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var permissions []Permission

/*
	func route() {
		router := mux.NewRouter()

		router.HandleFunc("/permissions", getPermissions).Methods("GET")
		router.HandleFunc("/permissions", createPermission).Methods("POST")
		router.HandleFunc("/permissions/{id}", updatePermission).Methods("PUT")
		router.HandleFunc("/permissions/{id}", deletePermission).Methods("DELETE")

		fmt.Println("Starting server at port 8080")
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatal(err)
		}
	}
*/
func getPermissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}

func createPermission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var permission Permission
	_ = json.NewDecoder(r.Body).Decode(&permission)
	permissions = append(permissions, permission)
	json.NewEncoder(w).Encode(permission)
}

func updatePermission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range permissions {
		if item.ID == params["id"] {
			permissions = append(permissions[:index], permissions[index+1:]...)
			var permission Permission
			_ = json.NewDecoder(r.Body).Decode(&permission)
			permission.ID = params["id"]
			permissions = append(permissions, permission)
			json.NewEncoder(w).Encode(permission)
			return
		}
	}
	json.NewEncoder(w).Encode(&Permission{})
}

func deletePermission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range permissions {
		if item.ID == params["id"] {
			permissions = append(permissions[:index], permissions[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(permissions)
}
