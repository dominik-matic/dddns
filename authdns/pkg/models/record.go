package models

type DNSRecord struct {
	Name  string
	Type  string
	Value string
	TTL   uint32
}
