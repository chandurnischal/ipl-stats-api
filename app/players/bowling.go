package players

import "database/sql"

func GetBowlingStats(query string, db *sql.DB) BowlStats {
	var stats BowlStats

	row := db.QueryRow(query)

	_ = row.Scan(
		&stats.Innings,
		&stats.TotalRuns,
		&stats.TotalWickets,
		&stats.Dots,
		&stats.Maidens,
		&stats.BowlingAverage,
		&stats.StrikeRate,
		&stats.Economy,
		&stats.FourWicketHauls,
		&stats.FiveWicketHauls,
	)

	return stats
}

func GetBowlingPerformanceStats(query string, db *sql.DB) BowlPerformanceStats {
	var performance BowlPerformanceStats

	row := db.QueryRow(query)

	_ = row.Scan(
		&performance.BestBowling.Wickets,
		&performance.BestBowling.Runs,
		&performance.BestBowling.Team,
		&performance.BestBowling.Against,
		&performance.BestBowling.Year,
		&performance.Last5Innings.Wickets,
		&performance.Last5Innings.Runs,
		&performance.Last5Innings.Economy,
		&performance.Last5Innings.BowlingAverage,
		&performance.Last10Innings.Wickets,
		&performance.Last10Innings.Runs,
		&performance.Last10Innings.Economy,
		&performance.Last10Innings.BowlingAverage,
	)

	return performance
}
