package dnsserver

import (
	"log"

	"github.com/miekg/dns"
)

func HandleDNSRequest(writer dns.ResponseWriter, requestMessage *dns.Msg) {
	var responseMessage = new(dns.Msg)
	responseMessage.SetReply(requestMessage)
	responseMessage.Authoritative = true

	for _, question := range requestMessage.Question {
		log.Printf("Query: %s %s", question.Name, dns.TypeToString[question.Qtype])
		// perhaps add some type of logic if Resolve returns nil
		var resourceRecords []dns.RR = Resolve(question.Name, question.Qtype)
		responseMessage.Answer = append(responseMessage.Answer, resourceRecords...)
	}

	err := writer.WriteMsg(responseMessage)
	if err != nil {
		log.Printf("Failed to write DNS response: %v", err)
	}
}
