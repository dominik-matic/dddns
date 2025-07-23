package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dominik-matic/dddns/apiserver/internal/apiserver"
	"github.com/dominik-matic/dddns/apiserver/internal/db"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	authToken := os.Getenv("AUTH_TOKEN")
	err := db.Connect(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	http.HandleFunc("/", apiserver.NewUpdateHandler(authToken))

	port := "53530"
	log.Printf("Listening on :%s", port)
	if http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
