package players

import (
	"database/sql"
)

func GetBattingStats(query string, db *sql.DB) BatStats {
	var stats BatStats

	row := db.QueryRow(query)

	row.Scan(
		&stats.Innings,
		&stats.TotalRuns,
		&stats.TotalBalls,
		&stats.Centuries,
		&stats.HalfCenturies,
		&stats.Fours,
		&stats.Sixes,
		&stats.StrikeRate,
		&stats.BattingAverage,
		&stats.Ducks,
	)

	return stats

}

func GetBattingPerformanceStats(query string, db *sql.DB) BatPerformanceStats {
	var stats BatPerformanceStats

	row := db.QueryRow(query)

	row.Scan(
		&stats.HighestScore.Runs,
		&stats.HighestScore.Balls,
		&stats.HighestScore.Team,
		&stats.HighestScore.Against,
		&stats.HighestScore.Year,
		&stats.Last5Innings.Runs,
		&stats.Last5Innings.Boundaries,
		&stats.Last5Innings.StrikeRate,
		&stats.Last5Innings.BattingAverage,
		&stats.Last10Innings.Runs,
		&stats.Last10Innings.Boundaries,
		&stats.Last10Innings.StrikeRate,
		&stats.Last10Innings.BattingAverage,
	)

	return stats
}
