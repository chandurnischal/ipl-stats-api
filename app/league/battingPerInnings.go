package league

import (
	"database/sql"
	"fmt"
)

func GetMost4sPerInnings(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
		SELECT 
			a.player, 
			a.team, 
			b.team AS against, 
			a.season,
			a.r, 
			a.b, 
			a.r * 100 / a.b AS strikerate, 
			a.4s, 
			a.6s 
		FROM 
		(
			SELECT 
				matchID, 
				team, 
				player, 
				season,
				r, 
				b, 
				4s, 
				6s 
			FROM 
				batting 
			ORDER BY 
				4s DESC, 
				r DESC 
			LIMIT 1
		) AS a 
		JOIN 
		(
			SELECT 
				matchID, 
				team 
			FROM 
				scores
		) AS b 
		ON 
			a.matchID = b.matchID 
			AND a.team != b.team;
		`
	} else {
		query = fmt.Sprintf(
			`
		SELECT 
			a.player, 
			a.team, 
			b.team AS against,
			a.season, 
			a.r, 
			a.b, 
			a.r * 100 / a.b AS strikerate, 
			a.4s, 
			a.6s 
		FROM 
		(
			SELECT 
				matchID, 
				team, 
				player, 
				season,
				r, 
				b, 
				4s, 
				6s 
			FROM 
				batting WHERE season = %d
			ORDER BY 
				4s DESC, 
				r DESC 
			LIMIT 1
		) AS a 
		JOIN 
		(
			SELECT 
				matchID, 
				team 
			FROM 
				scores
		) AS b 
		ON 
			a.matchID = b.matchID 
			AND a.team != b.team;
		`, season,
		)
	}

	return GetBattingInnings(query, db)
}

func GetMost6sPerInnings(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
		SELECT 
			a.player, 
			a.team, 
			b.team AS against, 
			a.season,
			a.r, 
			a.b, 
			a.r * 100 / a.b AS strikerate, 
			a.4s, 
			a.6s 
		FROM 
		(
			SELECT 
				matchID, 
				team, 
				player,
				season, 
				r, 
				b, 
				4s, 
				6s 
			FROM 
				batting 
			ORDER BY 
				6s DESC, 
				r DESC 
			LIMIT 1
		) AS a 
		JOIN 
		(
			SELECT 
				matchID, 
				team 
			FROM 
				scores
		) AS b 
		ON 
			a.matchID = b.matchID 
			AND a.team != b.team;
		`
	} else {
		query = fmt.Sprintf(
			`
		SELECT 
			a.player, 
			a.team, 
			b.team AS against, 
			a.season,
			a.r, 
			a.b, 
			a.r * 100 / a.b AS strikerate, 
			a.4s, 
			a.6s 
		FROM 
		(
			SELECT 
				matchID, 
				team, 
				player, 
				season,
				r, 
				b, 
				4s, 
				6s 
			FROM 
				batting WHERE season = %d
			ORDER BY 
				6s DESC, 
				r DESC 
			LIMIT 1
		) AS a 
		JOIN 
		(
			SELECT 
				matchID, 
				team 
			FROM 
				scores
		) AS b 
		ON 
			a.matchID = b.matchID 
			AND a.team != b.team;
		`, season,
		)
	}

	return GetBattingInnings(query, db)
}

func GetFastest50(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
			SELECT 
				a.player, 
				a.team, 
				b.team AS against, 
				a.season,
				a.r, 
				a.b, 
				a.r * 100 / a.b AS strikerate, 
				a.4s, 
				a.6s 
			FROM 
			(
				SELECT 
					matchID, 
					team, 
					player, 
					season,
					r, 
					b, 
					4s, 
					6s 
				FROM 
					batting WHERE r >= 50
				ORDER BY 
					b
				LIMIT 1
			) AS a 
			JOIN 
			(
				SELECT 
					matchID, 
					team 
				FROM 
					scores
			) AS b 
			ON 
				a.matchID = b.matchID 
				AND a.team != b.team;		
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT 
				a.player, 
				a.team, 
				b.team AS against, 
				a.season,
				a.r, 
				a.b, 
				a.r * 100 / a.b AS strikerate, 
				a.4s, 
				a.6s 
			FROM 
			(
				SELECT 
					matchID, 
					team, 
					player,
					season, 
					r, 
					b, 
					4s, 
					6s 
				FROM 
					batting WHERE r >= 50 AND season = %d 
				ORDER BY 
					b
				LIMIT 1
			) AS a 
			JOIN 
			(
				SELECT 
					matchID, 
					team 
				FROM 
					scores
			) AS b 
			ON 
				a.matchID = b.matchID 
				AND a.team != b.team;`,
			season,
		)
	}

	return GetBattingInnings(query, db)
}

func GetFastest100(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
			SELECT 
				a.player, 
				a.team, 
				b.team AS against, 
				a.season,
				a.r, 
				a.b, 
				a.r * 100 / a.b AS strikerate, 
				a.4s, 
				a.6s 
			FROM 
			(
				SELECT 
					matchID, 
					team, 
					player, 
					season,
					r, 
					b, 
					4s, 
					6s 
				FROM 
					batting WHERE r >= 100
				ORDER BY 
					b
				LIMIT 1
			) AS a 
			JOIN 
			(
				SELECT 
					matchID, 
					team 
				FROM 
					scores
			) AS b 
			ON 
				a.matchID = b.matchID 
				AND a.team != b.team;		
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT 
				a.player, 
				a.team, 
				b.team AS against, 
				a.season,
				a.r, 
				a.b, 
				a.r * 100 / a.b AS strikerate, 
				a.4s, 
				a.6s 
			FROM 
			(
				SELECT 
					matchID, 
					team, 
					player,
					season, 
					r, 
					b, 
					4s, 
					6s 
				FROM 
					batting WHERE r >= 100 AND season = %d 
				ORDER BY 
					b
				LIMIT 1
			) AS a 
			JOIN 
			(
				SELECT 
					matchID, 
					team 
				FROM 
					scores
			) AS b 
			ON 
				a.matchID = b.matchID 
				AND a.team != b.team;`,
			season,
		)
	}

	return GetBattingInnings(query, db)
}

func GetHighestScore(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
			SELECT 
				a.player, 
				a.team, 
				b.team AS against, 
				a.season,
				a.r, 
				a.b, 
				a.r * 100 / a.b AS strikerate, 
				a.4s, 
				a.6s 
			FROM 
			(
				SELECT 
					matchID, 
					team, 
					player, 
					season,
					r, 
					b, 
					4s, 
					6s 
				FROM 
					batting
				ORDER BY 
					r DESC, b ASC
				LIMIT 1
			) AS a 
			JOIN 
			(
				SELECT 
					matchID, 
					team 
				FROM 
					scores
			) AS b 
			ON 
				a.matchID = b.matchID 
				AND a.team != b.team;		
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT 
				a.player, 
				a.team, 
				b.team AS against, 
				a.season,
				a.r, 
				a.b, 
				a.r * 100 / a.b AS strikerate, 
				a.4s, 
				a.6s 
			FROM 
			(
				SELECT 
					matchID, 
					team, 
					player,
					season, 
					r, 
					b, 
					4s, 
					6s 
				FROM 
					batting WHERE season = %d 
				ORDER BY 
					r DESC, b ASC
				LIMIT 1
			) AS a 
			JOIN 
			(
				SELECT 
					matchID, 
					team 
				FROM 
					scores
			) AS b 
			ON 
				a.matchID = b.matchID 
				AND a.team != b.team;`,
			season,
		)
	}

	return GetBattingInnings(query, db)
}
