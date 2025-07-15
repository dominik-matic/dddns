package db

import (
	"database/sql"
	"strings"

	"github.com/dominik-matic/dddns/authdns/pkg/models"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dataSourceName string) error {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func QueryRecords(name, recordType string) ([]models.DNSRecord, error) {
	var records []models.DNSRecord

	rows, err := DB.Query("SELECT name, type, value, ttl FROM dns_records WHERE name = ? AND type = ?", name, recordType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record models.DNSRecord
		err := rows.Scan(&record.Name, &record.Type, &record.Value, &record.TTL)
		if err != nil {
			continue
		}
		records = append(records, record)
	}

	// wildcard fallback
	if len(records) == 0 {
		trimmed := strings.TrimPrefix(name, "*.")
		parts := strings.SplitN(trimmed, ".", 2)
		if len(parts) == 2 {
			wildcard := "*." + parts[1]
			return QueryRecords(wildcard, recordType)
		}
	}

	return records, nil
}
