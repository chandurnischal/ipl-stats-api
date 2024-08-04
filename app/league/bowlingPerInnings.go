package league

import (
	"database/sql"
	"fmt"
)

func GetMostDotsPerInnings(season int, db *sql.DB) BowlInnings {
	var query string

	if season == 0 {
		query = `
		SELECT 
			b1.player,
			b1.team,
			s.team AS against,
			b1.o,
			b1.r,
			b1.w,
			b1.0s,
			b1.econ,
			ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
		FROM 
			bowling b1
		JOIN 
			scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
		ORDER BY 
			b1.0s DESC 
		LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT 
				b1.player,
				b1.team,
				s.team AS against,
				b1.o,
				b1.r,
				b1.w,
				b1.0s,
				b1.econ,
				ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
			FROM 
				bowling b1
			JOIN 
				scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
			WHERE 
				b1.season = %d
			ORDER BY 
				b1.0s DESC 
			LIMIT 1;
			`,
			season,
		)
	}

	return GetBowlingInnings(query, db)

}

func GetBestEconomyPerInnings(season int, db *sql.DB) BowlInnings {
	var query string

	if season == 0 {
		query = `
		SELECT 
			b1.player,
			b1.team,
			s.team AS against,
			b1.o,
			b1.r,
			b1.w,
			b1.0s,
			b1.econ,
			ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
		FROM 
			bowling b1
		JOIN 
			scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
		ORDER BY 
			b1.econ ASC
		LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT 
				b1.player,
				b1.team,
				s.team AS against,
				b1.o,
				b1.r,
				b1.w,
				b1.0s,
				b1.econ,
				ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
			FROM 
				bowling b1
			JOIN 
				scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
			WHERE 
				b1.season = %d
			ORDER BY 
				b1.econ ASC 
			LIMIT 1;
			`,
			season,
		)
	}

	return GetBowlingInnings(query, db)

}

func GetBestBowlStrikeRatePerInnings(season int, db *sql.DB) BowlInnings {
	var query string

	if season == 0 {
		query = `
		SELECT 
			b1.player,
			b1.team,
			s.team AS against,
			b1.o,
			b1.r,
			b1.w,
			b1.0s,
			b1.econ,
			ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
		FROM 
			bowling b1
		JOIN 
			scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
		ORDER BY 
			strikeRate ASC
		LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT 
				b1.player,
				b1.team,
				s.team AS against,
				b1.o,
				b1.r,
				b1.w,
				b1.0s,
				b1.econ,
				ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
			FROM 
				bowling b1
			JOIN 
				scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
			WHERE 
				b1.season = %d
			ORDER BY 
				strikeRate ASC
			LIMIT 1;
			`,
			season,
		)
	}

	return GetBowlingInnings(query, db)

}

func GetMostConcededRuns(season int, db *sql.DB) BowlInnings {
	var query string

	if season == 0 {
		query = `
		SELECT 
			b1.player,
			b1.team,
			s.team AS against,
			b1.o,
			b1.r,
			b1.w,
			b1.0s,
			b1.econ,
			ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
		FROM 
			bowling b1
		JOIN 
			scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
		ORDER BY 
			b1.r DESC
		LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT 
				b1.player,
				b1.team,
				s.team AS against,
				b1.o,
				b1.r,
				b1.w,
				b1.0s,
				b1.econ,
				ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
			FROM 
				bowling b1
			JOIN 
				scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
			WHERE 
				b1.season = %d
			ORDER BY 
				b1.r DESC
			LIMIT 1;
			`,
			season,
		)
	}

	return GetBowlingInnings(query, db)
}

func GetBestBowlingFigures(season int, db *sql.DB) BowlInnings {
	var query string

	if season == 0 {
		query = `
		SELECT 
			b1.player,
			b1.team,
			s.team AS against,
			b1.o,
			b1.r,
			b1.w,
			b1.0s,
			b1.econ,
			ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
		FROM 
			bowling b1
		JOIN 
			scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
		ORDER BY 
			b1.w DESC, b1.r ASC
		LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT 
				b1.player,
				b1.team,
				s.team AS against,
				b1.o,
				b1.r,
				b1.w,
				b1.0s,
				b1.econ,
				ROUND(b1.b / NULLIF(b1.w, 0), 2) AS strikeRate
			FROM 
				bowling b1
			JOIN 
				scores s ON b1.matchID = s.matchID AND b1.team_id != s.team_id
			WHERE 
				b1.season = %d
			ORDER BY 
				b1.w DESC, b1.r ASC
			LIMIT 1;
			`,
			season,
		)
	}

	return GetBowlingInnings(query, db)

}
