package batting

import (
	"database/sql"
	"fmt"
	"strings"
)

func GetPlayerID(name string, db *sql.DB) (int, string, error) {
	var id int
	var playerName string
	var err error

	query := fmt.Sprintf(`SELECT * FROM players WHERE LOWER(player_name) LIKE '%%%s%%'`, strings.ToLower(name))
	row := db.QueryRow(query)

	err = row.Scan(&id, &playerName)

	if err != nil {
		return id, playerName, err
	}

	err = row.Err()

	if err != nil {
		return id, playerName, err
	}

	return id, playerName, err
}

func GetPlayerHistory(query string, db *sql.DB) (History, error) {
	var history History
	var err error

	row := db.QueryRow(query)

	err = row.Scan(
		&history.Matches,
		&history.Seasons,
		&history.Playoffs,
		&history.Finals,
		&history.Championships,
	)

	if err != nil {
		return history, err
	}

	err = row.Err()

	if err != nil {
		return history, err
	}

	return history, err
}
