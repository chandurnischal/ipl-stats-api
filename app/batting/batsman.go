package batting

import (
	"database/sql"
	"fmt"
)

func GetPlayerID(name string, db *sql.DB) (int, string, error) {
	var id int
	var playerName string

	query := fmt.Sprintf(`SELECT * FROM players WHERE LOWER(player_name) LIKE '%%%s%%'`, name)

	row := db.QueryRow(query)

	err := row.Scan(&id, &playerName)

	if err != nil {
		return id, playerName, err
	}

	return id, playerName, err

}

func GetTeams(id int, db *sql.DB) ([]string, error) {
	var teamsPlayed []string

	query := fmt.Sprintf(`
	SELECT DISTINCT team FROM (SELECT season, team FROM bowling WHERE player_id = %d UNION SELECT season, team FROM batting WHERE player_id = %d ORDER BY season DESC) AS a;
	`, id,
		id,
	)

	rows, err := db.Query(query)

	if err != nil {
		return teamsPlayed, err
	}

	for rows.Next() {
		var team string
		err = rows.Scan(&team)
		if err != nil {
			return teamsPlayed, err
		}
		teamsPlayed = append(teamsPlayed, team)
	}

	err = rows.Err()

	if err != nil {
		return teamsPlayed, err
	}

	return teamsPlayed, err
}

func GetCareerStats(query string, db *sql.DB) (CareerStats, error) {
	var career CareerStats
	row := db.QueryRow(query)

	err := row.Scan(
		&career.Matches,
		&career.Innings,
		&career.TotalRuns,
		&career.TotalBalls,
		&career.Centuries,
		&career.HalfCenturies,
		&career.Fours,
		&career.Sixes,
		&career.StrikeRate,
		&career.BattingAverage,
		&career.BoundariesPerInnings,
		&career.HighestScore,
		&career.Ducks,
	)

	if err != nil {
		return career, err
	}

	err = row.Err()

	if err != nil {
		return career, err
	}

	return career, err
}

func GetAppearances(query string, db *sql.DB) (Appearances, error) {
	var appearances Appearances

	row := db.QueryRow(query)

	err := row.Scan(
		&appearances.SeasonsPlayed,
		&appearances.PlayoffAppearances,
		&appearances.FinalAppearances,
		&appearances.ChampionshipsWon,
	)

	if err != nil {
		return appearances, err
	}

	err = row.Err()

	if err != nil {
		return appearances, err
	}

	return appearances, err
}
