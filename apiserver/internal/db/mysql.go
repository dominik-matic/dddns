package db

import (
	"database/sql"

	"github.com/dominik-matic/dddns/apiserver/pkg/models"
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

func InsertOrUpdate(data models.RequestData) error {
	var id int
	err := DB.QueryRow(`
		SELECT id FROM dns_records
		WHERE name = ? AND type = ? AND value = ?`,
		data.Name, data.Type, data.Value,
	).Scan(&id)

	switch err {
	case sql.ErrNoRows:
		_, err = DB.Exec(`
			INSERT INTO dns_records (name, type, value)
			VALUES (?, ?, ?)`,
			data.Name, data.Type, data.Value)
		return err
	case nil:
		_, err = DB.Exec(`
			UPDATE dns_records
			SET value = ?
			WHERE id = ?`,
			data.Value, id)
		return err
	}
	return err
}

func Delete(data models.RequestData) error {
	_, err := DB.Exec(`
		DELETE FROM dns_records
		WHERE name = ?
		AND type = ?`,
		data.Name, data.Type)
	return err
}
