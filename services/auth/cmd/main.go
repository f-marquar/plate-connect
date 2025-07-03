package main

// main package for the auth service
import (
	"fmt"
	"log"
	"net/http"

	"plate-connect/services/auth/db"

	"encoding/json"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal("Error while connecting to DB", err)
	}

	db.Migrate()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		sqlDB, err := db.DB.DB()
		if err != nil {
			http.Error(w, "Fehler bei DB-Instanz", http.StatusInternalServerError)
			return
		}
		err = sqlDB.Ping()
		if err != nil {
			http.Error(w, "DB nicht erreichbar", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Auth service and DB are healthy (GORM)")
	})


	http.HandleFunc("/register", registerHandler)

	fmt.Println("Starting Auth service on :7001")
	err = http.ListenAndServe(":7001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Nur POST erlaubt", http.StatusMethodNotAllowed)
		return
	}

	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Ungültige Eingabe", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Fehler beim Speichern: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User mit Kennzeichen %s registriert", user.Kennzeichen)
}