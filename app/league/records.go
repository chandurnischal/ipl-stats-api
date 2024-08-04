package league

import (
	"database/sql"
)

func GetBattingOverall(query string, db *sql.DB) BatOverall {
	var overall BatOverall

	row := db.QueryRow(query)

	row.Scan(
		&overall.Player,
		&overall.Team,
		&overall.Innings,
		&overall.Runs,
		&overall.Balls,
		&overall.Fours,
		&overall.Sixes,
		&overall.Average,
		&overall.StrikeRate,
		&overall.HighestScore,
		&overall.HalfCenturies,
		&overall.Centuries,
	)

	return overall
}

func GetBattingInnings(query string, db *sql.DB) BatInnings {
	var innings BatInnings

	row := db.QueryRow(query)

	row.Scan(
		&innings.Player,
		&innings.Team,
		&innings.Against,
		&innings.Year,
		&innings.Runs,
		&innings.Balls,
		&innings.StrikeRate,
		&innings.Fours,
		&innings.Sixes,
	)
	return innings
}

func GetBowlingOverall(query string, db *sql.DB) BowlOverall {
	var overall BowlOverall

	row := db.QueryRow(query)

	row.Scan(
		&overall.Player,
		&overall.Team,
		&overall.Innings,
		&overall.Overs,
		&overall.Runs,
		&overall.Wickets,
		&overall.Dots,
		&overall.Maidens,
		&overall.BestBowling,
		&overall.Average,
		&overall.Economy,
		&overall.StrikeRate,
		&overall.FourWicketHaul,
		&overall.FiveWicketHaul,
	)

	return overall
}

func GetBowlingInnings(query string, db *sql.DB) BowlInnings {
	var innings BowlInnings

	row := db.QueryRow(query)

	row.Scan(
		&innings.Player,
		&innings.Team,
		&innings.Against,
		&innings.Overs,
		&innings.Runs,
		&innings.Wickets,
		&innings.Dots,
		&innings.Economy,
		&innings.StrikeRate,
	)

	return innings
}
