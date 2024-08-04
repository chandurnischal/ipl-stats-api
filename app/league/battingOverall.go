package league

import (
	"database/sql"
	"fmt"
)

func GetMostRuns(season int, db *sql.DB) BatOverall {
	var query string
	var mostRuns BatOverall

	if season == 0 {
		query = `
		SELECT a.player, a.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
		a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries
		FROM
		(
			SELECT
				player_id,
				player,
				team,
				COUNT(*) AS innings, 
				SUM(r) AS totalruns, 
				SUM(b) AS totalballs, 
				SUM(4s) AS fours,
				SUM(6s) AS sixes,
				ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
				ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
				SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
				SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
			FROM batting 
			GROUP BY player_id, player, team
			ORDER BY SUM(r) DESC, SUM(b)
			LIMIT 1
		) AS a
		JOIN
		(
			SELECT player_id, CONCAT(MAX(r), '/', b) AS highestScore 
			FROM batting 
			GROUP BY player_id, b
			ORDER BY MAX(r) DESC, b 
		) AS b ON a.player_id = b.player_id LIMIT 1;
			`
	} else {
		query = fmt.Sprintf(`
		SELECT a.player, a.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
       	a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries
		FROM
		(
			SELECT
				player_id,
				player,
				team,
				COUNT(*) AS innings, 
				SUM(r) AS totalruns, 
				SUM(b) AS totalballs, 
				SUM(4s) AS fours,
				SUM(6s) AS sixes,
				ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
				ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
				SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
				SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
			FROM batting WHERE season = %d
			GROUP BY player_id, player, team
			ORDER BY SUM(r) DESC, SUM(b)
			LIMIT 1
		) AS a
		JOIN
		(
			SELECT player_id, CONCAT(MAX(r), '/', b) AS highestScore 
			FROM batting WHERE season = %d
			GROUP BY player_id, b
			ORDER BY MAX(r) DESC, b 
		) AS b ON a.player_id = b.player_id LIMIT 1;
		`, season,
			season,
		)
	}
	mostRuns = GetBattingOverall(query, db)

	return mostRuns
}

func GetMostFours(season int, db *sql.DB) BatOverall {
	var query string
	if season == 0 {
		query = `
		SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
       a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries 
		FROM
		(
			SELECT
				player_id,
				player,
				COUNT(*) AS innings, 
				SUM(r) AS totalruns, 
				SUM(b) AS totalballs, 
				SUM(4s) AS fours,
				SUM(6s) AS sixes,
				ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
				ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
				SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
				SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
			FROM batting 
			GROUP BY player_id, player
			ORDER BY SUM(4s) DESC, SUM(r) DESC
			LIMIT 1
		) AS a
		JOIN
		(
			SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
			FROM batting 
			GROUP BY player_id, team, b
			ORDER BY MAX(r) DESC, b 
		) AS b ON a.player_id = b.player_id LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
				a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries
			FROM
			(
				SELECT
					player_id,
					player,
					COUNT(*) AS innings, 
					SUM(r) AS totalruns, 
					SUM(b) AS totalballs, 
					SUM(4s) AS fours,
					SUM(6s) AS sixes,
					ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
					ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
					SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
					SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
				FROM batting WHERE season = %d
				GROUP BY player_id, player
				ORDER BY SUM(4s) DESC, SUM(r) DESC
				LIMIT 1
			) AS a
			JOIN
			(
				SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
				FROM batting WHERE season = %d
				GROUP BY player_id, team, b
				ORDER BY MAX(r) DESC, b 
			) AS b ON a.player_id = b.player_id LIMIT 1;
			`, season,
			season,
		)
	}

	return GetBattingOverall(query, db)
}

func GetMostSixes(season int, db *sql.DB) BatOverall {
	var query string
	if season == 0 {
		query = `
		SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
       a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries 
		FROM
		(
			SELECT
				player_id,
				player,
				COUNT(*) AS innings, 
				SUM(r) AS totalruns, 
				SUM(b) AS totalballs, 
				SUM(4s) AS fours,
				SUM(6s) AS sixes,
				ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
				ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
				SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
				SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
			FROM batting 
			GROUP BY player_id, player
			ORDER BY SUM(6s) DESC, SUM(r) DESC
			LIMIT 1
		) AS a
		JOIN
		(
			SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
			FROM batting 
			GROUP BY player_id, team, b
			ORDER BY MAX(r) DESC, b 
		) AS b ON a.player_id = b.player_id LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
				a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries
			FROM
			(
				SELECT
					player_id,
					player,
					COUNT(*) AS innings, 
					SUM(r) AS totalruns, 
					SUM(b) AS totalballs, 
					SUM(4s) AS fours,
					SUM(6s) AS sixes,
					ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
					ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
					SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
					SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
				FROM batting WHERE season = %d
				GROUP BY player_id, player
				ORDER BY SUM(6s) DESC, SUM(r) DESC
				LIMIT 1
			) AS a
			JOIN
			(
				SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
				FROM batting WHERE season = %d
				GROUP BY player_id, team, b
				ORDER BY MAX(r) DESC, b 
			) AS b ON a.player_id = b.player_id LIMIT 1;
			`, season,
			season,
		)
	}

	return GetBattingOverall(query, db)
}

func GetMost50s(season int, db *sql.DB) BatOverall {
	var query string
	if season == 0 {
		query = `
		SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
       a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries 
		FROM
		(
			SELECT
				player_id,
				player,
				COUNT(*) AS innings, 
				SUM(r) AS totalruns, 
				SUM(b) AS totalballs, 
				SUM(4s) AS fours,
				SUM(6s) AS sixes,
				ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
				ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
				SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
				SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
			FROM batting 
			GROUP BY player_id, player
			ORDER BY halfcenturies DESC, SUM(r) DESC
			LIMIT 1
		) AS a
		JOIN
		(
			SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
			FROM batting 
			GROUP BY player_id, team, b
			ORDER BY MAX(r) DESC, b 
		) AS b ON a.player_id = b.player_id LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
				a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries
			FROM
			(
				SELECT
					player_id,
					player,
					COUNT(*) AS innings, 
					SUM(r) AS totalruns, 
					SUM(b) AS totalballs, 
					SUM(4s) AS fours,
					SUM(6s) AS sixes,
					ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
					ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
					SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
					SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
				FROM batting WHERE season = %d
				GROUP BY player_id, player
				ORDER BY halfcenturies DESC, SUM(r) DESC
				LIMIT 1
			) AS a
			JOIN
			(
				SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
				FROM batting WHERE season = %d
				GROUP BY player_id, team, b
				ORDER BY MAX(r) DESC, b 
			) AS b ON a.player_id = b.player_id LIMIT 1;
			`, season,
			season,
		)
	}

	return GetBattingOverall(query, db)
}

func GetMost100s(season int, db *sql.DB) BatOverall {
	var query string
	if season == 0 {
		query = `
		SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
       a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries 
		FROM
		(
			SELECT
				player_id,
				player,
				COUNT(*) AS innings, 
				SUM(r) AS totalruns, 
				SUM(b) AS totalballs, 
				SUM(4s) AS fours,
				SUM(6s) AS sixes,
				ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
				ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
				SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
				SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
			FROM batting 
			GROUP BY player_id, player
			ORDER BY centuries DESC, SUM(r) DESC
			LIMIT 1
		) AS a
		JOIN
		(
			SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
			FROM batting 
			GROUP BY player_id, team, b
			ORDER BY MAX(r) DESC, b 
		) AS b ON a.player_id = b.player_id LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
				a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries
			FROM
			(
				SELECT
					player_id,
					player,
					COUNT(*) AS innings, 
					SUM(r) AS totalruns, 
					SUM(b) AS totalballs, 
					SUM(4s) AS fours,
					SUM(6s) AS sixes,
					ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
					ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
					SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
					SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
				FROM batting WHERE season = %d
				GROUP BY player_id, player
				ORDER BY centuries DESC, SUM(r) DESC
				LIMIT 1
			) AS a
			JOIN
			(
				SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
				FROM batting WHERE season = %d
				GROUP BY player_id, team, b
				ORDER BY MAX(r) DESC, b 
			) AS b ON a.player_id = b.player_id LIMIT 1;
			`, season,
			season,
		)
	}

	return GetBattingOverall(query, db)
}

func GetBestStrikeRate(season int, db *sql.DB) BatOverall {
	var query string
	if season == 0 {
		query = `
		SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
       a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries 
		FROM
		(
			SELECT
				player_id,
				player,
				COUNT(*) AS innings, 
				SUM(r) AS totalruns, 
				SUM(b) AS totalballs, 
				SUM(4s) AS fours,
				SUM(6s) AS sixes,
				ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
				ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
				SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
				SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
			FROM batting 
			GROUP BY player_id, player
			HAVING innings >= 12
			ORDER BY strikerate DESC
			LIMIT 1
		) AS a
		JOIN
		(
			SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
			FROM batting 
			GROUP BY player_id, team, b
			ORDER BY MAX(r) DESC, b 
		) AS b ON a.player_id = b.player_id LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
				a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries
			FROM
			(
				SELECT
					player_id,
					player,
					COUNT(*) AS innings, 
					SUM(r) AS totalruns, 
					SUM(b) AS totalballs, 
					SUM(4s) AS fours,
					SUM(6s) AS sixes,
					ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
					ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
					SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
					SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
				FROM batting WHERE season = %d
				GROUP BY player_id, player
				HAVING innings >= 12
				ORDER BY strikerate DESC
				LIMIT 1
			) AS a
			JOIN
			(
				SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
				FROM batting WHERE season = %d
				GROUP BY player_id, team, b
				ORDER BY MAX(r) DESC, b 
			) AS b ON a.player_id = b.player_id LIMIT 1;
			`, season,
			season,
		)
	}

	return GetBattingOverall(query, db)
}

func GetBestBattingAverage(season int, db *sql.DB) BatOverall {
	var query string
	if season == 0 {
		query = `
		SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
       a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries 
		FROM
		(
			SELECT
				player_id,
				player,
				COUNT(*) AS innings, 
				SUM(r) AS totalruns, 
				SUM(b) AS totalballs, 
				SUM(4s) AS fours,
				SUM(6s) AS sixes,
				ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%not out%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
				ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
				SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
				SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
			FROM batting 
			GROUP BY player_id, player
			HAVING innings >= 12
			ORDER BY battingaverage DESC
			LIMIT 1
		) AS a
		JOIN
		(
			SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
			FROM batting 
			GROUP BY player_id, team, b
			ORDER BY MAX(r) DESC, b 
		) AS b ON a.player_id = b.player_id LIMIT 1;
		`
	} else {
		query = fmt.Sprintf(
			`
			SELECT a.player, b.team, a.innings, a.totalruns, a.totalballs, a.fours, a.sixes, 
				a.battingaverage, a.strikerate, b.highestScore, a.halfcenturies, a.centuries
			FROM
			(
				SELECT
					player_id,
					player,
					COUNT(*) AS innings, 
					SUM(r) AS totalruns, 
					SUM(b) AS totalballs, 
					SUM(4s) AS fours,
					SUM(6s) AS sixes,
					ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
					ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
					SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
					SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries
				FROM batting WHERE season = %d
				GROUP BY player_id, player
				HAVING innings >= 12
				ORDER BY battingaverage DESC
				LIMIT 1
			) AS a
			JOIN
			(
				SELECT player_id, team, CONCAT(MAX(r), '/', b) AS highestScore 
				FROM batting WHERE season = %d
				GROUP BY player_id, team, b
				ORDER BY MAX(r) DESC, b 
			) AS b ON a.player_id = b.player_id LIMIT 1;
			`, season,
			season,
		)
	}

	return GetBattingOverall(query, db)
}
