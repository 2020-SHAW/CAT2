package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./contact.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable()

	r := mux.NewRouter()
	r.HandleFunc("/contact", handleContact).Methods("POST")

	// Add CORS support
	handler := cors.Default().Handler(r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS contact (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"email" TEXT,
		"message" TEXT
	);`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	var contactForm ContactForm
	err := json.NewDecoder(r.Body).Decode(&contactForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = insertContact(contactForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact form submitted successfully"))
}

func insertContact(contactForm ContactForm) error {
	insertSQL := `INSERT INTO contact (name, email, message) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(contactForm.Name, contactForm.Email, contactForm.Message)
	if err != nil {
		return err
	}
	return nil
}
