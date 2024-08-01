package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() error {
	err := godotenv.Load("db.env")

	if err != nil {
		return err
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSERNAME"),
		Passwd: os.Getenv("DBPASSWORD"),
		Addr:   os.Getenv("DBHOST"),
		DBName: os.Getenv("DBNAME"),
	}

	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		return pingErr
	}
	return nil
}

func ReturnValueFromDB[T any](query string, db *sql.DB) (T, error) {
	var result T

	row := db.QueryRow(query)

	err := row.Scan(&result)

	if err != nil {
		return result, err
	}

	err = row.Err()

	if err != nil {
		return result, err
	}
	return result, err

}
