package league

import (
	"database/sql"
	"fmt"
)

func GetMostRuns(season int, db *sql.DB) BatOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.totalRuns DESC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(`
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting WHERE season = %d
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting WHERE season = %d
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.totalRuns DESC
	LIMIT 1;
		`, season,
			season,
		)
	}

	res := GetBattingOverall(query, db)

	if season == 0 {
		res.Team = GetPlayerTeam(res.Player, 0, db)
	} else {
		res.Team = GetPlayerTeam(res.Player, season, db)
	}

	return res
}

func GetMostFours(season int, db *sql.DB) BatOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.total4s DESC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(`
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting WHERE season = %d
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting WHERE season = %d
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.total4s DESC
	LIMIT 1;
		`, season,
			season,
		)
	}

	res := GetBattingOverall(query, db)

	if season == 0 {
		res.Team = GetPlayerTeam(res.Player, 0, db)
	} else {
		res.Team = GetPlayerTeam(res.Player, season, db)
	}

	return res
}

func GetMostSixes(season int, db *sql.DB) BatOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.total6s DESC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(`
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting WHERE season = %d
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting WHERE season = %d
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.total6s DESC
	LIMIT 1;
		`, season,
			season,
		)
	}

	res := GetBattingOverall(query, db)

	if season == 0 {
		res.Team = GetPlayerTeam(res.Player, 0, db)
	} else {
		res.Team = GetPlayerTeam(res.Player, season, db)
	}

	return res
}

func GetMostFifties(season int, db *sql.DB) BatOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.total50s DESC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(`
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting WHERE season = %d
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting WHERE season = %d
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.total50s DESC
	LIMIT 1;
		`, season,
			season,
		)
	}

	res := GetBattingOverall(query, db)

	if season == 0 {
		res.Team = GetPlayerTeam(res.Player, 0, db)
	} else {
		res.Team = GetPlayerTeam(res.Player, season, db)
	}

	return res
}

func GetMostHundrends(season int, db *sql.DB) BatOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.total100s DESC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(`
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting WHERE season = %d
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting WHERE season = %d
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.total100s DESC
	LIMIT 1;
		`, season,
			season,
		)
	}

	res := GetBattingOverall(query, db)

	if season == 0 {
		res.Team = GetPlayerTeam(res.Player, 0, db)
	} else {
		res.Team = GetPlayerTeam(res.Player, season, db)
	}

	return res
}

func GetBestStrikeRate(season int, db *sql.DB) BatOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.strikeRate DESC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(`
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting WHERE season = %d
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting WHERE season = %d
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.strikeRate DESC
	LIMIT 1;
		`, season,
			season,
		)
	}

	res := GetBattingOverall(query, db)

	if season == 0 {
		res.Team = GetPlayerTeam(res.Player, 0, db)
	} else {
		res.Team = GetPlayerTeam(res.Player, season, db)
	}

	return res
}

func GetBestBattingAverage(season int, db *sql.DB) BatOverall {
	var query string

	if season == 0 {
		query = `
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.battingaverage DESC
	LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(`
		WITH aggregated AS (
		SELECT 
			player_id,
			player, 
			COUNT(*) AS innings,
			SUM(r) AS totalRuns, 
			SUM(b) AS totalBalls,
			SUM(4s) AS total4s,
			SUM(6s) AS total6s,
			ROUND(SUM(r) / NULLIF(SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 0), 2) AS battingaverage,
			ROUND(SUM(r) * 100 / NULLIF(SUM(b), 0), 2) AS strikeRate,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS total50s,
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS total100s
		FROM 
			batting WHERE season = %d
		GROUP BY 
			player_id, player
	),
	highest_score AS (
		SELECT 
			player_id,
			CONCAT(r, '/', b) AS highestScore,
			ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY r DESC, b ASC) AS rn
		FROM 
			batting WHERE season = %d
	)
	SELECT 
		a.player,
		a.innings,
		a.totalRuns,
		a.totalBalls,
		a.total4s,
		a.total6s,
		a.battingaverage,
		a.strikeRate,
		h.highestScore,
		a.total50s,
		a.total100s
	FROM 
		aggregated a
	JOIN 
		highest_score h 
	ON 
		a.player_id = h.player_id
	WHERE 
		h.rn = 1
	ORDER BY 
		a.battingaverage DESC
	LIMIT 1;
		`, season,
			season,
		)
	}

	res := GetBattingOverall(query, db)

	if season == 0 {
		res.Team = GetPlayerTeam(res.Player, 0, db)
	} else {
		res.Team = GetPlayerTeam(res.Player, season, db)
	}

	return res
}
