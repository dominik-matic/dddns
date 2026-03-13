package db

import (
	"database/sql"

	"github.com/dominik-matic/dddns/apiserver/pkg/models"
	_ "github.com/go-sql-driver/mysql"
)

// Currently no support for multiple records of the same
// (name, type), but I don't really need that.
// Maybe something to implement in the future,
// should be pretty easy

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
		WHERE name = ? AND type = ?`,
		data.Name, data.Type,
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
