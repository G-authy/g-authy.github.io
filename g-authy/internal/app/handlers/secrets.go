package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Secret struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var secrets []Secret

/*
	func route() {
		router := mux.NewRouter()

		router.HandleFunc("/secrets", getSecrets).Methods("GET")
		router.HandleFunc("/secrets", createSecret).Methods("POST")
		router.HandleFunc("/secrets/{id}", updateSecret).Methods("PUT")
		router.HandleFunc("/secrets/{id}", deleteSecret).Methods("DELETE")

		fmt.Println("Starting server at port 8080")
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatal(err)
		}
	}
*/

func getSecrets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(secrets)
}

func createSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var secret Secret
	_ = json.NewDecoder(r.Body).Decode(&secret)
	secrets = append(secrets, secret)
	json.NewEncoder(w).Encode(secret)
}

func updateSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range secrets {
		if item.ID == params["id"] {
			secrets = append(secrets[:index], secrets[index+1:]...)
			var secret Secret
			_ = json.NewDecoder(r.Body).Decode(&secret)
			secret.ID = params["id"]
			secrets = append(secrets, secret)
			json.NewEncoder(w).Encode(secret)
			return
		}
	}
	json.NewEncoder(w).Encode(&Secret{})
}

func deleteSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range secrets {
		if item.ID == params["id"] {
			secrets = append(secrets[:index], secrets[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(secrets)
}
