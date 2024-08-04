package league

import (
	"database/sql"
	"fmt"
)

func GetMost4sPerInnings(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
		WITH top_fours AS (
		SELECT 
			matchID, 
			player, 
			team_id, 
			team, 
			season, 
			r, 
			b, 
			r * 100 / NULLIF(b, 0) AS strikerate, 
			4s, 
			6s
		FROM 
			batting
		ORDER BY 
			4s DESC
		LIMIT 1
	)
	SELECT 
		t.player,
		t.team,
		s.team AS against,
		t.season,
		t.r,
		t.b,
		t.strikerate,
		t.4s,
		t.6s
	FROM 
		top_fours t
	JOIN 
		scores s
	ON 
		t.matchID = s.matchID
		AND t.team_id != s.team_id;
		`
	} else {
		query = fmt.Sprintf(`
		WITH top_fours AS (
		SELECT 
			matchID, 
			player, 
			team_id, 
			team, 
			season, 
			r, 
			b, 
			r * 100 / NULLIF(b, 0) AS strikerate, 
			4s, 
			6s
		FROM 
			batting WHERE season = %d
		ORDER BY 
			4s DESC
		LIMIT 1
	)
	SELECT 
		t.player,
		t.team,
		s.team AS against,
		t.season,
		t.r,
		t.b,
		t.strikerate,
		t.4s,
		t.6s
	FROM 
		top_fours t
	JOIN 
		scores s
	ON 
		t.matchID = s.matchID
		AND t.team_id != s.team_id;
		`, season)
	}

	return GetBattingInnings(query, db)
}

func GetMost6sPerInnings(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
		WITH top_fours AS (
		SELECT 
			matchID, 
			player, 
			team_id, 
			team, 
			season, 
			r, 
			b, 
			r * 100 / NULLIF(b, 0) AS strikerate, 
			4s, 
			6s
		FROM 
			batting
		ORDER BY 
			6s DESC
		LIMIT 1
	)
	SELECT 
		t.player,
		t.team,
		s.team AS against,
		t.season,
		t.r,
		t.b,
		t.strikerate,
		t.4s,
		t.6s
	FROM 
		top_fours t
	JOIN 
		scores s
	ON 
		t.matchID = s.matchID
		AND t.team_id != s.team_id;
		`
	} else {
		query = fmt.Sprintf(`
		WITH top_fours AS (
		SELECT 
			matchID, 
			player, 
			team_id, 
			team, 
			season, 
			r, 
			b, 
			r * 100 / NULLIF(b, 0) AS strikerate, 
			4s, 
			6s
		FROM 
			batting WHERE season = %d
		ORDER BY 
			6s DESC
		LIMIT 1
	)
	SELECT 
		t.player,
		t.team,
		s.team AS against,
		t.season,
		t.r,
		t.b,
		t.strikerate,
		t.4s,
		t.6s
	FROM 
		top_fours t
	JOIN 
		scores s
	ON 
		t.matchID = s.matchID
		AND t.team_id != s.team_id;
		`, season)
	}

	return GetBattingInnings(query, db)
}

func GetFastest50(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
		WITH top_fours AS (
		SELECT 
			matchID, 
			player, 
			team_id, 
			team, 
			season, 
			r, 
			b, 
			r * 100 / NULLIF(b, 0) AS strikerate, 
			4s, 
			6s
		FROM 
			batting WHERE r >= 50 AND r < 100
		ORDER BY 
			r ASC, b ASC
		LIMIT 1
	)
	SELECT 
		t.player,
		t.team,
		s.team AS against,
		t.season,
		t.r,
		t.b,
		t.strikerate,
		t.4s,
		t.6s
	FROM 
		top_fours t
	JOIN 
		scores s
	ON 
		t.matchID = s.matchID
		AND t.team_id != s.team_id;
		`
	} else {
		query = fmt.Sprintf(`
		WITH top_fours AS (
		SELECT 
			matchID, 
			player, 
			team_id, 
			team, 
			season, 
			r, 
			b, 
			r * 100 / NULLIF(b, 0) AS strikerate, 
			4s, 
			6s
		FROM 
			batting WHERE season = %d AND r >= 50 AND r < 100
		ORDER BY 
			r ASC, b ASC
		LIMIT 1
	)
	SELECT 
		t.player,
		t.team,
		s.team AS against,
		t.season,
		t.r,
		t.b,
		t.strikerate,
		t.4s,
		t.6s
	FROM 
		top_fours t
	JOIN 
		scores s
	ON 
		t.matchID = s.matchID
		AND t.team_id != s.team_id;
		`, season)
	}

	return GetBattingInnings(query, db)
}

func GetFastest100(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
		WITH top_fours AS (
		SELECT 
			matchID, 
			player, 
			team_id, 
			team, 
			season, 
			r, 
			b, 
			r * 100 / NULLIF(b, 0) AS strikerate, 
			4s, 
			6s
		FROM 
			batting WHERE r >= 100
		ORDER BY 
			r ASC, b ASC
		LIMIT 1
	)
	SELECT 
		t.player,
		t.team,
		s.team AS against,
		t.season,
		t.r,
		t.b,
		t.strikerate,
		t.4s,
		t.6s
	FROM 
		top_fours t
	JOIN 
		scores s
	ON 
		t.matchID = s.matchID
		AND t.team_id != s.team_id;
		`
	} else {
		query = fmt.Sprintf(`
		WITH top_fours AS (
		SELECT 
			matchID, 
			player, 
			team_id, 
			team, 
			season, 
			r, 
			b, 
			r * 100 / NULLIF(b, 0) AS strikerate, 
			4s, 
			6s
		FROM 
			batting WHERE season = %d AND r >= 100
		ORDER BY 
			r ASC, b ASC
		LIMIT 1
	)
	SELECT 
		t.player,
		t.team,
		s.team AS against,
		t.season,
		t.r,
		t.b,
		t.strikerate,
		t.4s,
		t.6s
	FROM 
		top_fours t
	JOIN 
		scores s
	ON 
		t.matchID = s.matchID
		AND t.team_id != s.team_id;
		`, season)
	}

	return GetBattingInnings(query, db)
}

func GetHighestScore(season int, db *sql.DB) BatInnings {
	var query string

	if season == 0 {
		query = `
		WITH top_fours AS (
			SELECT 
				matchID, 
				player, 
				team_id, 
				team, 
				season, 
				r, 
				b, 
				r * 100 / NULLIF(b, 0) AS strikerate, 
				4s, 
				6s
			FROM 
				batting
			ORDER BY 
				r DESC, b ASC
			LIMIT 1
		)
		SELECT 
			t.player,
			t.team,
			s.team AS against,
			t.season,
			t.r,
			t.b,
			t.strikerate,
			t.4s,
			t.6s
		FROM 
			top_fours t
		JOIN 
			scores s
		ON 
			t.matchID = s.matchID
			AND t.team_id != s.team_id;
		`
	} else {
		query = fmt.Sprintf(`
			WITH top_fours AS (
				SELECT 
					matchID, 
					player, 
					team_id, 
					team, 
					season, 
					r, 
					b, 
					r * 100 / NULLIF(b, 0) AS strikerate, 
					4s, 
					6s
				FROM 
					batting WHERE season = %d
				ORDER BY 
					r DESC, b ASC
				LIMIT 1
			)
			SELECT 
				t.player,
				t.team,
				s.team AS against,
				t.season,
				t.r,
				t.b,
				t.strikerate,
				t.4s,
				t.6s
			FROM 
				top_fours t
			JOIN 
				scores s
			ON 
				t.matchID = s.matchID
				AND t.team_id != s.team_id;
		`, season)
	}

	return GetBattingInnings(query, db)
}
