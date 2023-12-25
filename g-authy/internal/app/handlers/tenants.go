package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Tenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var tenants []Tenant

/*
func route() {
	router := mux.NewRouter()

	router.HandleFunc("/tenants", getTenants).Methods("GET")
	router.HandleFunc("/tenants", createTenant).Methods("POST")
	router.HandleFunc("/tenants/{id}", updateTenant).Methods("PUT")
	router.HandleFunc("/tenants/{id}", deleteTenant).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
*/

func getTenants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tenants)
}

func createTenant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tenant Tenant
	_ = json.NewDecoder(r.Body).Decode(&tenant)
	tenants = append(tenants, tenant)
	json.NewEncoder(w).Encode(tenant)
}

func updateTenant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tenants {
		if item.ID == params["id"] {
			tenants = append(tenants[:index], tenants[index+1:]...)
			var tenant Tenant
			_ = json.NewDecoder(r.Body).Decode(&tenant)
			tenant.ID = params["id"]
			tenants = append(tenants, tenant)
			json.NewEncoder(w).Encode(tenant)
			return
		}
	}
	json.NewEncoder(w).Encode(&Tenant{})
}

func deleteTenant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tenants {
		if item.ID == params["id"] {
			tenants = append(tenants[:index], tenants[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tenants)
}
