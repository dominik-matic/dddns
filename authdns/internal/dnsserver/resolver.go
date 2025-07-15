package dnsserver

import (
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/dominik-matic/dddns/authdns/internal/db"
	"github.com/dominik-matic/dddns/authdns/pkg/models"
	"github.com/miekg/dns"
)

func Resolve(name string, qType uint16) []dns.RR {
	var results []dns.RR
	records, err := db.QueryRecords(strings.ToLower(strings.TrimSuffix(name, ".")), dns.TypeToString[qType])
	if err != nil {
		log.Printf("Resolve error: %v", err)
		return nil
	}

	for _, record := range records {
		resourceRecord := BuildRR(name, record)
		if resourceRecord != nil {
			results = append(results, resourceRecord)
		}
	}

	return results
}

func BuildRR(name string, record models.DNSRecord) dns.RR {
	header := dns.RR_Header{
		Name:   dns.Fqdn(name),
		Rrtype: dns.StringToType[record.Type],
		Class:  dns.ClassINET,
		Ttl:    record.TTL,
	}

	switch record.Type {
	case "A":
		return &dns.A{Hdr: header, A: net.ParseIP(record.Value).To4()}
	case "AAAA":
		return &dns.AAAA{Hdr: header, AAAA: net.ParseIP(record.Value)}
	case "CNAME":
		return &dns.CNAME{Hdr: header, Target: dns.Fqdn(record.Value)}
	case "TXT":
		return &dns.TXT{Hdr: header, Txt: []string{record.Value}}
	case "NS":
		return &dns.NS{Hdr: header, Ns: dns.Fqdn(record.Value)}
	case "SOA":
		var parts []string = strings.Fields(record.Value)
		if len(parts) != 7 {
			return nil
		}
		return &dns.SOA{
			Hdr:     header,
			Ns:      dns.Fqdn(parts[0]),
			Mbox:    dns.Fqdn(parts[1]),
			Serial:  parseSerial(parts[2]),
			Refresh: parseUint32(parts[3]),
			Retry:   parseUint32(parts[4]),
			Expire:  parseUint32(parts[5]),
			Minttl:  parseUint32(parts[6]),
		}
	default:
		return nil
	}
}

func parseSerial(serial string) uint32 {
	// go uses the following date for time formatting:
	// 2006-01-02 15:04:05
	t, err := time.Parse("2006010215", serial)
	if err != nil {
		return uint32(time.Now().Unix())
	}
	return uint32(t.Unix())
}

func parseUint32(s string) uint32 {
	u32, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Printf("parseUint32 error: %v", err)
	}
	return uint32(u32)
}
