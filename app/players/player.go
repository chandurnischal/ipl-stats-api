package players

import (
	"database/sql"
	"fmt"
)

func GetPlayerByName(name string, db *sql.DB) (int, string, error) {
	var id int
	var playerName string

	query := fmt.Sprintf(`SELECT * FROM players WHERE LOWER(player_name) LIKE '%%%s%%'`, name)
	row := db.QueryRow(query)

	err := row.Scan(&id, &playerName)

	if err != nil {
		return id, playerName, err
	}

	err = row.Err()

	if err != nil {
		return id, playerName, err

	}

	return id, playerName, err

}

func GetTeams(query string, db *sql.DB) ([]string, error) {
	var teams []string
	var err error

	rows, err := db.Query(query)

	if err != nil {
		return teams, err
	}

	for rows.Next() {
		var team string
		err = rows.Scan(&team)

		if err != nil {
			return teams, err
		}

		teams = append(teams, team)
	}

	err = rows.Err()

	if err != nil {
		return teams, err
	}

	return teams, err
}

func GetCareerHistory(query string, db *sql.DB) (History, error) {
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
