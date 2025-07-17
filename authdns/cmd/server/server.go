package main

import (
	"log"
	"os"

	"github.com/dominik-matic/dddns/authdns/internal/db"
	"github.com/dominik-matic/dddns/authdns/internal/dnsserver"
	"github.com/miekg/dns"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	err := db.Connect(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	dnsServer := &dns.Server{Addr: ":53", Net: "udp"}
	dns.HandleFunc("dominikmatic.com", dnsserver.HandleDNSRequest)

	log.Println("Starting authdns on :53")
	err = dnsServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start authdns: %s", err.Error())
	}
}
