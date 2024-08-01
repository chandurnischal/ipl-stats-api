package teams

import (
	"database/sql"
	"fmt"
)

func GetTeamID(name string, db *sql.DB) (int, string, error) {
	var id int
	var teamName string
	query := fmt.Sprintf(`
	SELECT * FROM teams WHERE LOWER(team_name) LIKE '%%%s%%'
	`, name)

	row := db.QueryRow(query)

	err := row.Scan(&id, &teamName)

	if err != nil {
		return id, teamName, err
	}

	err = row.Err()

	if err != nil {
		return id, teamName, err
	}

	return id, teamName, nil
}

func GetMatches(query string, db *sql.DB) (Matches, error) {
	var matches Matches

	row := db.QueryRow(query)

	err := row.Scan(
		&matches.Played,
		&matches.Won,
		&matches.Lost,
		&matches.Tied,
		&matches.NoResult,
		&matches.BattingFirstPerc,
		&matches.FieldingFirstPerc,
		&matches.WinPercentage,
	)

	if err != nil {
		return matches, err
	}

	err = row.Err()

	if err != nil {
		return matches, err
	}

	return matches, nil

}

func GetAppearances(query string, db *sql.DB) (Appearances, error) {
	var playoffs Appearances

	row := db.QueryRow(query)

	err := row.Scan(
		&playoffs.Played,
		&playoffs.Appearances,
		&playoffs.Finals,
		&playoffs.Championships,
	)

	if err != nil {
		return playoffs, err
	}

	return playoffs, nil

}

func GetTeamStats(query string, db *sql.DB) (Stats, error) {
	var stats Stats

	row := db.QueryRow(query)

	err := row.Scan(
		&stats.HighestScore,
		&stats.LowestScore,
		&stats.AverageRuns,
		&stats.AverageWickets,
		&stats.TotalRuns,
		&stats.TotalWickets,
	)

	if err != nil {
		return stats, err
	}

	err = row.Err()

	if err != nil {
		return stats, err
	}

	return stats, err

}

func GetIndividualPerformance(query string, db *sql.DB) (Indiviudal, error) {
	var individual Indiviudal

	row := db.QueryRow(query)

	err := row.Scan(
		&individual.TopRunScorer.Name,
		&individual.TopRunScorer.Runs,
		&individual.TopWicketTaker.Name,
		&individual.TopWicketTaker.Wickets,
		&individual.BestBatting,
		&individual.BestBowling,
	)

	if err != nil {
		return individual, err
	}

	err = row.Err()

	if err != nil {
		return individual, err
	}

	return individual, nil

}
