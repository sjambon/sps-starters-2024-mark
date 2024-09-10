package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	_ "github.com/lib/pq"
)

type Campaign struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var db *sql.DB

func main() {
	// Connect to the database
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/campaigns", getCampaigns).Methods("GET")
	r.HandleFunc("/campaigns", createCampaign).Methods("POST")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	})

	// Start server
	handler := c.Handler(r)
	log.Println("Server starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func getCampaigns(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description FROM campaigns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	campaigns := []Campaign{}
	for rows.Next() {
		var c Campaign
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		campaigns = append(campaigns, c)
	}

	json.NewEncoder(w).Encode(campaigns)
}

func createCampaign(w http.ResponseWriter, r *http.Request) {
	var c Campaign
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.QueryRow("INSERT INTO campaigns (name, description) VALUES ($1, $2) RETURNING id", c.Name, c.Description).Scan(&c.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}