package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users []User

/*
	func route() {
		router := mux.NewRouter()

		router.HandleFunc("/users", getUsers).Methods("GET")
		router.HandleFunc("/users", createUser).Methods("POST")
		router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
		router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

		fmt.Println("Starting server at port 8080")
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatal(err)
		}
	}
*/
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
