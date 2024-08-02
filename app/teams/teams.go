package teams

import (
	"database/sql"
	"fmt"
)

func GetTeamID(name string, db *sql.DB) (int, string, error) {
	var id int
	var teamName string
	var err error

	query := fmt.Sprintf(`SELECT * FROM teams WHERE LOWER(team_name) LIKE '%%%s%%'`, name)

	row := db.QueryRow(query)

	err = row.Scan(&id, &teamName)

	if err != nil {
		return id, teamName, err
	}

	err = row.Err()

	if err != nil {
		return id, teamName, err
	}

	return id, teamName, err

}

func GetTeamHistory(query string, db *sql.DB) (History, error) {
	var history History
	var err error

	row := db.QueryRow(query)

	err = row.Scan(
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

func GetTeamMatchRecord(query string, db *sql.DB) (MatchRecord, error) {
	var matchRecord MatchRecord
	var err error

	row := db.QueryRow(query)

	err = row.Scan(
		&matchRecord.Played,
		&matchRecord.Won,
		&matchRecord.Lost,
		&matchRecord.NoResult,
		&matchRecord.Tied,
	)

	if err != nil {
		return matchRecord, err
	}

	err = row.Err()

	if err != nil {
		return matchRecord, err
	}

	return matchRecord, err
}

func GetPerformancestats(query string, db *sql.DB) (PerformanceStats, error) {
	var performanceStats PerformanceStats
	var winPerc WinPercentage
	var err error

	row := db.QueryRow(query)

	err = row.Scan(
		&performanceStats.TotalRuns,
		&performanceStats.TotalWickets,
		&performanceStats.AverageRuns,
		&performanceStats.AverageWickets,
		&performanceStats.HighestScore,
		&performanceStats.BestBowling,
		&winPerc.Overall,
		&winPerc.BatFirst,
		&winPerc.FieldFirst,
		&winPerc.Last5Matches,
		&winPerc.Last10Matches,
	)

	if err != nil {
		return performanceStats, err
	}

	err = row.Err()

	if err != nil {
		return performanceStats, err
	}

	performanceStats.WinPerc = winPerc

	return performanceStats, err
}

func GetPlayerAchievements(query string, db *sql.DB) (PlayerAchievements, error) {
	var playerAchievements PlayerAchievements
	var err error

	row := db.QueryRow(query)

	err = row.Scan(
		&playerAchievements.TopScorer.Player,
		&playerAchievements.TopScorer.Score,
		&playerAchievements.TopScorer.Against,
		&playerAchievements.TopScorer.Year,
		&playerAchievements.TopBowler.Player,
		&playerAchievements.TopBowler.Figure,
		&playerAchievements.TopBowler.Against,
		&playerAchievements.TopBowler.Year,
		&playerAchievements.MostCenturies.Player,
		&playerAchievements.MostCenturies.Centuries,
		&playerAchievements.MostWickets.Player,
		&playerAchievements.MostWickets.Wickets,
	)

	if err != nil {
		return playerAchievements, err
	}

	err = row.Err()

	if err != nil {
		return playerAchievements, err
	}

	return playerAchievements, err

}
