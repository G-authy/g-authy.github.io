package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Vault struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var vaults []Vault

/*
func route() {
	router := mux.NewRouter()

	router.HandleFunc("/vaults", getVaults).Methods("GET")
	router.HandleFunc("/vaults", createVault).Methods("POST")
	router.HandleFunc("/vaults/{id}", updateVault).Methods("PUT")
	router.HandleFunc("/vaults/{id}", deleteVault).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
*/

func getVaults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vaults)
}

func createVault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vault Vault
	_ = json.NewDecoder(r.Body).Decode(&vault)
	vaults = append(vaults, vault)
	json.NewEncoder(w).Encode(vault)
}

func updateVault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range vaults {
		if item.ID == params["id"] {
			vaults = append(vaults[:index], vaults[index+1:]...)
			var vault Vault
			_ = json.NewDecoder(r.Body).Decode(&vault)
			vault.ID = params["id"]
			vaults = append(vaults, vault)
			json.NewEncoder(w).Encode(vault)
			return
		}
	}
	json.NewEncoder(w).Encode(&Vault{})
}

func deleteVault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range vaults {
		if item.ID == params["id"] {
			vaults = append(vaults[:index], vaults[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(vaults)
}
