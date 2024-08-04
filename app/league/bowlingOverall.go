package league

import (
	"database/sql"
	"fmt"
)

func GetMostWickets(season int, db *sql.DB) BowlOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			ROUND(SUM(FLOOR(b / 6) + (b % 6) / 10.0), 1) AS overs,
			SUM(r) AS runs, 
			SUM(w) AS wickets,
			SUM(0s) AS dots,
			SUM(m) AS maidens,
			ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
			ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
			ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
			SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
			SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
		FROM 
			bowling
		GROUP BY 
			player_id, player
	),
	best_bowling AS (
		SELECT 
			player_id,
			CONCAT(w, '/', r) AS bestBowling
		FROM 
			bowling b1
		WHERE NOT EXISTS (
			SELECT 1
			FROM bowling b2
			WHERE b1.player_id = b2.player_id
			AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
		)
	)
	SELECT 
		a.player,
		a.innings,
		a.overs,
		a.runs,
		a.wickets,
		a.dots,
		a.maidens,
		b.bestBowling,
		a.bowlingaverage,
		a.economy,
		a.strikeRate,
		a.fourWicket,
		a.fiveWicket
	FROM 
		aggregated a
	JOIN 
		best_bowling b ON a.player_id = b.player_id
	ORDER BY 
		a.wickets DESC, a.runs ASC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			WITH aggregated AS (
			SELECT 
				player_id,
				player, 
				COUNT(*) AS innings,
				ROUND(SUM(FLOOR(b / 6) + (b %% 6) / 10.0), 1) AS overs,
				SUM(r) AS runs, 
				SUM(w) AS wickets,
				SUM(0s) AS dots,
				SUM(m) AS maidens,
				ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
				ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
				ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
				SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
				SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
			FROM 
				bowling
			WHERE 
				season = %d
			GROUP BY 
				player_id, player
		),
		best_bowling AS (
			SELECT 
				player_id,
				CONCAT(w, '/', r) AS bestBowling
			FROM 
				bowling b1
			WHERE 
				season = %d
				AND NOT EXISTS (
					SELECT 1
					FROM bowling b2
					WHERE b1.player_id = b2.player_id
					AND b2.season = %d
					AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
			)
		)
		SELECT 
			a.player,
			a.innings,
			a.overs,
			a.runs,
			a.wickets,
			a.dots,
			a.maidens,
			b.bestBowling,
			a.bowlingaverage,
			a.economy,
			a.strikeRate,
			a.fourWicket,
			a.fiveWicket
		FROM 
			aggregated a
		JOIN 
			best_bowling b ON a.player_id = b.player_id
		ORDER BY 
			a.wickets DESC, a.runs ASC
		LIMIT 1;
			`, season,
			season, season,
		)
	}

	bowl := GetBowlingOverall(query, db)
	bowl.Team = GetPlayerTeam(bowl.Player, season, db)

	return bowl
}

func GetMostMaidens(season int, db *sql.DB) BowlOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			ROUND(SUM(FLOOR(b / 6) + (b % 6) / 10.0), 1) AS overs,
			SUM(r) AS runs, 
			SUM(w) AS wickets,
			SUM(0s) AS dots,
			SUM(m) AS maidens,
			ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
			ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
			ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
			SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
			SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
		FROM 
			bowling
		GROUP BY 
			player_id, player
	),
	best_bowling AS (
		SELECT 
			player_id,
			CONCAT(w, '/', r) AS bestBowling
		FROM 
			bowling b1
		WHERE NOT EXISTS (
			SELECT 1
			FROM bowling b2
			WHERE b1.player_id = b2.player_id
			AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
		)
	)
	SELECT 
		a.player,
		a.innings,
		a.overs,
		a.runs,
		a.wickets,
		a.dots,
		a.maidens,
		b.bestBowling,
		a.bowlingaverage,
		a.economy,
		a.strikeRate,
		a.fourWicket,
		a.fiveWicket
	FROM 
		aggregated a
	JOIN 
		best_bowling b ON a.player_id = b.player_id
	ORDER BY 
		a.maidens DESC, a.runs ASC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			WITH aggregated AS (
			SELECT 
				player_id,
				player, 
				COUNT(*) AS innings,
				ROUND(SUM(FLOOR(b / 6) + (b %% 6) / 10.0), 1) AS overs,
				SUM(r) AS runs, 
				SUM(w) AS wickets,
				SUM(0s) AS dots,
				SUM(m) AS maidens,
				ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
				ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
				ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
				SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
				SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
			FROM 
				bowling
			WHERE 
				season = %d
			GROUP BY 
				player_id, player
		),
		best_bowling AS (
			SELECT 
				player_id,
				CONCAT(w, '/', r) AS bestBowling
			FROM 
				bowling b1
			WHERE 
				season = %d
				AND NOT EXISTS (
					SELECT 1
					FROM bowling b2
					WHERE b1.player_id = b2.player_id
					AND b2.season = %d
					AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
			)
		)
		SELECT 
			a.player,
			a.innings,
			a.overs,
			a.runs,
			a.wickets,
			a.dots,
			a.maidens,
			b.bestBowling,
			a.bowlingaverage,
			a.economy,
			a.strikeRate,
			a.fourWicket,
			a.fiveWicket
		FROM 
			aggregated a
		JOIN 
			best_bowling b ON a.player_id = b.player_id
		ORDER BY 
			a.maidens DESC, a.runs ASC
		LIMIT 1;
			`, season,
			season, season,
		)
	}

	bowl := GetBowlingOverall(query, db)
	bowl.Team = GetPlayerTeam(bowl.Player, season, db)

	return bowl
}

func GetMostDots(season int, db *sql.DB) BowlOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			ROUND(SUM(FLOOR(b / 6) + (b % 6) / 10.0), 1) AS overs,
			SUM(r) AS runs, 
			SUM(w) AS wickets,
			SUM(0s) AS dots,
			SUM(m) AS maidens,
			ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
			ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
			ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
			SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
			SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
		FROM 
			bowling
		GROUP BY 
			player_id, player
	),
	best_bowling AS (
		SELECT 
			player_id,
			CONCAT(w, '/', r) AS bestBowling
		FROM 
			bowling b1
		WHERE NOT EXISTS (
			SELECT 1
			FROM bowling b2
			WHERE b1.player_id = b2.player_id
			AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
		)
	)
	SELECT 
		a.player,
		a.innings,
		a.overs,
		a.runs,
		a.wickets,
		a.dots,
		a.maidens,
		b.bestBowling,
		a.bowlingaverage,
		a.economy,
		a.strikeRate,
		a.fourWicket,
		a.fiveWicket
	FROM 
		aggregated a
	JOIN 
		best_bowling b ON a.player_id = b.player_id
	ORDER BY 
		a.dots DESC, a.runs ASC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			WITH aggregated AS (
			SELECT 
				player_id,
				player, 
				COUNT(*) AS innings,
				ROUND(SUM(FLOOR(b / 6) + (b %% 6) / 10.0), 1) AS overs,
				SUM(r) AS runs, 
				SUM(w) AS wickets,
				SUM(0s) AS dots,
				SUM(m) AS maidens,
				ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
				ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
				ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
				SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
				SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
			FROM 
				bowling
			WHERE 
				season = %d
			GROUP BY 
				player_id, player
		),
		best_bowling AS (
			SELECT 
				player_id,
				CONCAT(w, '/', r) AS bestBowling
			FROM 
				bowling b1
			WHERE 
				season = %d
				AND NOT EXISTS (
					SELECT 1
					FROM bowling b2
					WHERE b1.player_id = b2.player_id
					AND b2.season = %d
					AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
			)
		)
		SELECT 
			a.player,
			a.innings,
			a.overs,
			a.runs,
			a.wickets,
			a.dots,
			a.maidens,
			b.bestBowling,
			a.bowlingaverage,
			a.economy,
			a.strikeRate,
			a.fourWicket,
			a.fiveWicket
		FROM 
			aggregated a
		JOIN 
			best_bowling b ON a.player_id = b.player_id
		ORDER BY 
			a.dots DESC, a.runs ASC
		LIMIT 1;
			`, season,
			season, season,
		)
	}

	bowl := GetBowlingOverall(query, db)
	bowl.Team = GetPlayerTeam(bowl.Player, season, db)

	return bowl
}

func GetBestBowlingAverage(season int, db *sql.DB) BowlOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			ROUND(SUM(FLOOR(b / 6) + (b % 6) / 10.0), 1) AS overs,
			SUM(r) AS runs, 
			SUM(w) AS wickets,
			SUM(0s) AS dots,
			SUM(m) AS maidens,
			ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
			ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
			ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
			SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
			SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
		FROM 
			bowling
		GROUP BY 
			player_id, player
	),
	best_bowling AS (
		SELECT 
			player_id,
			CONCAT(w, '/', r) AS bestBowling
		FROM 
			bowling b1
		WHERE NOT EXISTS (
			SELECT 1
			FROM bowling b2
			WHERE b1.player_id = b2.player_id
			AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
		)
	)
	SELECT 
		a.player,
		a.innings,
		a.overs,
		a.runs,
		a.wickets,
		a.dots,
		a.maidens,
		b.bestBowling,
		a.bowlingaverage,
		a.economy,
		a.strikeRate,
		a.fourWicket,
		a.fiveWicket
	FROM 
		aggregated a
	JOIN 
		best_bowling b ON a.player_id = b.player_id
	ORDER BY 
		a.bowlingaverage ASC, a.runs ASC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			WITH aggregated AS (
			SELECT 
				player_id,
				player, 
				COUNT(*) AS innings,
				ROUND(SUM(FLOOR(b / 6) + (b %% 6) / 10.0), 1) AS overs,
				SUM(r) AS runs, 
				SUM(w) AS wickets,
				SUM(0s) AS dots,
				SUM(m) AS maidens,
				ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
				ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
				ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
				SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
				SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
			FROM 
				bowling
			WHERE 
				season = %d
			GROUP BY 
				player_id, player
		),
		best_bowling AS (
			SELECT 
				player_id,
				CONCAT(w, '/', r) AS bestBowling
			FROM 
				bowling b1
			WHERE 
				season = %d
				AND NOT EXISTS (
					SELECT 1
					FROM bowling b2
					WHERE b1.player_id = b2.player_id
					AND b2.season = %d
					AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
			)
		)
		SELECT 
			a.player,
			a.innings,
			a.overs,
			a.runs,
			a.wickets,
			a.dots,
			a.maidens,
			b.bestBowling,
			a.bowlingaverage,
			a.economy,
			a.strikeRate,
			a.fourWicket,
			a.fiveWicket
		FROM 
			aggregated a
		JOIN 
			best_bowling b ON a.player_id = b.player_id
		ORDER BY 
			a.bowlingaverage ASC, a.runs ASC
		LIMIT 1;
			`, season,
			season, season,
		)
	}

	bowl := GetBowlingOverall(query, db)
	bowl.Team = GetPlayerTeam(bowl.Player, season, db)

	return bowl
}

func GetBestEconomy(season int, db *sql.DB) BowlOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			ROUND(SUM(FLOOR(b / 6) + (b % 6) / 10.0), 1) AS overs,
			SUM(r) AS runs, 
			SUM(w) AS wickets,
			SUM(0s) AS dots,
			SUM(m) AS maidens,
			ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
			ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
			ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
			SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
			SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
		FROM 
			bowling
		GROUP BY 
			player_id, player
	),
	best_bowling AS (
		SELECT 
			player_id,
			CONCAT(w, '/', r) AS bestBowling
		FROM 
			bowling b1
		WHERE NOT EXISTS (
			SELECT 1
			FROM bowling b2
			WHERE b1.player_id = b2.player_id
			AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
		)
	)
	SELECT 
		a.player,
		a.innings,
		a.overs,
		a.runs,
		a.wickets,
		a.dots,
		a.maidens,
		b.bestBowling,
		a.bowlingaverage,
		a.economy,
		a.strikeRate,
		a.fourWicket,
		a.fiveWicket
	FROM 
		aggregated a
	JOIN 
		best_bowling b ON a.player_id = b.player_id
	ORDER BY 
		a.economy ASC, a.runs ASC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			WITH aggregated AS (
			SELECT 
				player_id,
				player, 
				COUNT(*) AS innings,
				ROUND(SUM(FLOOR(b / 6) + (b %% 6) / 10.0), 1) AS overs,
				SUM(r) AS runs, 
				SUM(w) AS wickets,
				SUM(0s) AS dots,
				SUM(m) AS maidens,
				ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
				ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
				ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
				SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
				SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
			FROM 
				bowling
			WHERE 
				season = %d
			GROUP BY 
				player_id, player
		),
		best_bowling AS (
			SELECT 
				player_id,
				CONCAT(w, '/', r) AS bestBowling
			FROM 
				bowling b1
			WHERE 
				season = %d
				AND NOT EXISTS (
					SELECT 1
					FROM bowling b2
					WHERE b1.player_id = b2.player_id
					AND b2.season = %d
					AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
			)
		)
		SELECT 
			a.player,
			a.innings,
			a.overs,
			a.runs,
			a.wickets,
			a.dots,
			a.maidens,
			b.bestBowling,
			a.bowlingaverage,
			a.economy,
			a.strikeRate,
			a.fourWicket,
			a.fiveWicket
		FROM 
			aggregated a
		JOIN 
			best_bowling b ON a.player_id = b.player_id
		ORDER BY 
			a.economy ASC, a.runs ASC
		LIMIT 1;
			`, season,
			season, season,
		)
	}

	bowl := GetBowlingOverall(query, db)
	bowl.Team = GetPlayerTeam(bowl.Player, season, db)

	return bowl
}

func GetBestBowlStrikeRate(season int, db *sql.DB) BowlOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			ROUND(SUM(FLOOR(b / 6) + (b % 6) / 10.0), 1) AS overs,
			SUM(r) AS runs, 
			SUM(w) AS wickets,
			SUM(0s) AS dots,
			SUM(m) AS maidens,
			ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
			ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
			ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
			SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
			SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
		FROM 
			bowling
		GROUP BY 
			player_id, player
	),
	best_bowling AS (
		SELECT 
			player_id,
			CONCAT(w, '/', r) AS bestBowling
		FROM 
			bowling b1
		WHERE NOT EXISTS (
			SELECT 1
			FROM bowling b2
			WHERE b1.player_id = b2.player_id
			AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
		)
	)
	SELECT 
		a.player,
		a.innings,
		a.overs,
		a.runs,
		a.wickets,
		a.dots,
		a.maidens,
		b.bestBowling,
		a.bowlingaverage,
		a.economy,
		a.strikeRate,
		a.fourWicket,
		a.fiveWicket
	FROM 
		aggregated a
	JOIN 
		best_bowling b ON a.player_id = b.player_id
	ORDER BY 
		a.strikeRate DESC, a.runs ASC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			WITH aggregated AS (
			SELECT 
				player_id,
				player, 
				COUNT(*) AS innings,
				ROUND(SUM(FLOOR(b / 6) + (b %% 6) / 10.0), 1) AS overs,
				SUM(r) AS runs, 
				SUM(w) AS wickets,
				SUM(0s) AS dots,
				SUM(m) AS maidens,
				ROUND(SUM(r) / NULLIF(SUM(w), 0), 2) AS bowlingaverage,
				ROUND(SUM(r) / NULLIF(SUM(b) / 6, 0), 2) AS economy,
				ROUND(SUM(b) / NULLIF(SUM(w), 0), 2) AS strikeRate,
				SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicket,
				SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicket 
			FROM 
				bowling
			WHERE 
				season = %d
			GROUP BY 
				player_id, player
		),
		best_bowling AS (
			SELECT 
				player_id,
				CONCAT(w, '/', r) AS bestBowling
			FROM 
				bowling b1
			WHERE 
				season = %d
				AND NOT EXISTS (
					SELECT 1
					FROM bowling b2
					WHERE b1.player_id = b2.player_id
					AND b2.season = %d
					AND (b2.w > b1.w OR (b2.w = b1.w AND b2.r < b1.r))
			)
		)
		SELECT 
			a.player,
			a.innings,
			a.overs,
			a.runs,
			a.wickets,
			a.dots,
			a.maidens,
			b.bestBowling,
			a.bowlingaverage,
			a.economy,
			a.strikeRate,
			a.fourWicket,
			a.fiveWicket
		FROM 
			aggregated a
		JOIN 
			best_bowling b ON a.player_id = b.player_id
		ORDER BY 
			a.strikeRate DESC, a.runs ASC
		LIMIT 1;
			`, season,
			season, season,
		)
	}

	bowl := GetBowlingOverall(query, db)
	bowl.Team = GetPlayerTeam(bowl.Player, season, db)

	return bowl
}
