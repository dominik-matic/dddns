package main

import (
	"log"

	"github.com/dominik-matic/dddns/authdns/internal/db"
	"github.com/dominik-matic/dddns/authdns/internal/dnsserver"
	"github.com/miekg/dns"
)

func main() {
	err := db.Connect("root:password@tcp(mysql:3306)/dnsdb")
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
