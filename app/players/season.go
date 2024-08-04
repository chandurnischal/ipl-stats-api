package players

import (
	"database/sql"
	"fmt"
	"ipl-api/database"
)

func GetSeasonTeams(id, season int, db *sql.DB) ([]string, error) {
	query := fmt.Sprintf(
		`
		SELECT DISTINCT team
		FROM (
			SELECT team FROM batting WHERE player_id = %d AND season = %d
			UNION
			SELECT team FROM bowling WHERE player_id = %d AND season = %d
			UNION
			SELECT team FROM batting WHERE player_id = %d AND season != %d
			UNION
			SELECT team FROM bowling WHERE player_id = %d AND season != %d
		) AS combined_teams;
		`, id, season,
		id, season,
		id, season,
		id, season,
	)

	return GetTeams(query, db)
}

func GetSeasonHistory(id, season int, db *sql.DB) (History, error) {
	query := fmt.Sprintf(
		`
		SELECT
			COUNT(DISTINCT m.matchID) AS total_matches,
			COUNT(DISTINCT m.season) AS total_seasons,
			COUNT(DISTINCT CASE WHEN m.match_type NOT LIKE '%%match%%' THEN m.season END) AS playoffs,
			COUNT(DISTINCT CASE WHEN m.match_type LIKE 'Final%%' THEN m.matchID END) AS finals,
			COUNT(DISTINCT CASE WHEN m.match_type LIKE 'Final%%' AND m.team_id = m.winner_id THEN m.matchID END) AS championships
		FROM
			(SELECT b.matchID, b.season, b.team_id, m.match_type, m.winner_id
			FROM batting b
			JOIN matches m ON b.matchID = m.matchID
			WHERE b.player_id = %d AND m.season = %d
			UNION ALL
			SELECT b.matchID, b.season, b.team_id, m.match_type, m.winner_id
			FROM bowling b
			JOIN matches m ON b.matchID = m.matchID
			WHERE b.player_id = %d AND m.season = %d) AS m;
		`, id, season,
		id, season,
	)
	return GetCareerHistory(query, db)
}

func GetSeasonBattingPerformance(id, season int, db *sql.DB) BatPerformanceStats {
	query := fmt.Sprintf(
		`
		WITH highest_score AS (
		SELECT 
			b.matchID,
			b.r AS highestScoreRuns,
			b.b AS highestScoreBalls,
			b.team_id AS highestScoreTeamID,
			b.team AS highestScoreTeam,
			m.season AS highestScoreSeason
		FROM 
			batting b
		JOIN 
			matches m ON b.matchID = m.matchID
		WHERE 
			b.player_id = %d AND b.season = %d
		ORDER BY 
			b.r DESC, b.b
		LIMIT 1
	),
	opponent_team AS (
		SELECT 
			s.matchID,
			s.team_id AS highestScoreAgainstID,
			s.team AS highestScoreAgainst
		FROM 
			scores s
		JOIN 
			highest_score hs ON s.matchID = hs.matchID
		WHERE 
			s.team_id != hs.highestScoreTeamID
	),
	latest_matches_5 AS (
		SELECT 
			m.matchID
		FROM 
			matches m
		JOIN 
			batting b ON m.matchID = b.matchID
		WHERE 
			b.player_id = %d AND b.season = %d
		ORDER BY 
			m.date DESC
		LIMIT 5
	),
	latest_matches_10 AS (
		SELECT 
			m.matchID
		FROM 
			matches m
		JOIN 
			batting b ON m.matchID = b.matchID
		WHERE 
			b.player_id = %d AND b.season = %d
		ORDER BY 
			m.date DESC
		LIMIT 10
	),
	stats_5_matches AS (
		SELECT
			SUM(b.r) AS total_runs,
			SUM(b.4s + b.6s) AS total_boundaries,
			ROUND(SUM(b.r) / SUM(b.b) * 100, 2) AS strike_rate,
			ROUND(SUM(b.r) / COUNT(CASE WHEN b.dismissal_info NOT LIKE '%%not out%%' THEN 1 END), 2) AS batting_average
		FROM 
			batting b
		JOIN 
			latest_matches_5 lm ON b.matchID = lm.matchID
		WHERE 
			b.player_id = %d AND b.season = %d
	),
	stats_10_matches AS (
		SELECT
			SUM(b.r) AS total_runs,
			SUM(b.4s + b.6s) AS total_boundaries,
			ROUND(SUM(b.r) / SUM(b.b) * 100, 2) AS strike_rate,
			ROUND(SUM(b.r) / COUNT(CASE WHEN b.dismissal_info NOT LIKE '%%not out%%' THEN 1 END), 2) AS batting_average
		FROM 
			batting b
		JOIN 
			latest_matches_10 lm ON b.matchID = lm.matchID
		WHERE 
			b.player_id = %d AND b.season = %d
	)
	SELECT 
		hs.highestScoreRuns,
		hs.highestScoreBalls,
		hs.highestScoreTeam,
		ot.highestScoreAgainst,
		hs.highestScoreSeason,
		stats5.total_runs AS total_runs_5,
		stats5.total_boundaries AS total_boundaries_5,
		stats5.strike_rate AS strike_rate_5,
		stats5.batting_average AS batting_average_5,
		stats10.total_runs AS total_runs_10,
		stats10.total_boundaries AS total_boundaries_10,
		stats10.strike_rate AS strike_rate_10,
		stats10.batting_average AS batting_average_10
	FROM 
		highest_score hs
	JOIN 
		opponent_team ot ON hs.matchID = ot.matchID
	JOIN 
		stats_5_matches stats5
	JOIN 
		stats_10_matches stats10;
	`, id, season,
		id, season,
		id, season,
		id, season,
		id, season,
	)

	return GetBattingPerformanceStats(query, db)
}

func GetSeasonBattingStats(id, season int, db *sql.DB) BatStats {
	query := fmt.Sprintf(
		`SELECT 
			COUNT(*) AS innings, 
			SUM(r) AS totalruns, 
			SUM(b) AS totalballs, 
			SUM(CASE WHEN r >= 100 THEN 1 ELSE 0 END) AS centuries,
			SUM(CASE WHEN r >= 50 AND r < 100 THEN 1 ELSE 0 END) AS halfcenturies,
			SUM(4s) AS fours, 
			SUM(6s) AS sixes,
			ROUND((SUM(r) * 100) / SUM(b), 2) AS strikerate, 
			ROUND(SUM(r) / SUM(CASE WHEN dismissal_info NOT LIKE '%%not out%%' THEN 1 ELSE 0 END), 2) AS battingaverage, 
			SUM(CASE WHEN r = 0 THEN 1 ELSE 0 END) AS ducks
		FROM batting 
		WHERE player_id = %d AND season = %d;
	`,
		id, season,
	)

	return GetBattingStats(query, db)
}

func GetSeasonBowlingStats(id, season int, db *sql.DB) BowlStats {
	query := fmt.Sprintf(
		`
		SELECT 
		COUNT(*) AS innings,
		SUM(r) AS totalRuns,
		SUM(w) AS totalWickets,
		SUM(0s) AS totalDots,
		SUM(m) AS totalMaidens,
		CASE 
			WHEN SUM(w) = 0 THEN NULL
			ELSE SUM(r) / SUM(w)
		END AS bowlingAvg,
		CASE 
			WHEN SUM(w) = 0 THEN NULL
			ELSE SUM(b) / SUM(w)
		END AS strikeRate,
		CASE 
			WHEN SUM(b) = 0 THEN NULL
			ELSE SUM(r) / SUM(b)
		END AS economy,
		SUM(CASE WHEN w = 4 THEN 1 ELSE 0 END) AS fourWicketHaul,
		SUM(CASE WHEN w = 5 THEN 1 ELSE 0 END) AS fiveWicketHaul
	FROM bowling
	WHERE player_id = %d AND season = %d;
		`,
		id, season,
	)

	return GetBowlingStats(query, db)
}

func GetSeasonBowlingPerformance(id, season int, db *sql.DB) BowlPerformanceStats {
	query := fmt.Sprintf(
		`
		WITH latest_matches AS (
			SELECT 
				m.matchID,
				m.date,
				b.w,
				b.r,
				b.b,
				ROW_NUMBER() OVER (ORDER BY m.date DESC) AS row_num
			FROM 
				matches m
			JOIN 
				bowling b ON m.matchID = b.matchID
			WHERE 
				b.player_id = %d AND b.season = %d
		),
		latest_5 AS (
			SELECT 
				matchID
			FROM 
				latest_matches
			WHERE 
				row_num <= 5
		),
		latest_10 AS (
			SELECT 
				matchID
			FROM 
				latest_matches
			WHERE 
				row_num <= 10
		),
		stats AS (
			SELECT 
				first.w, 
				first.r, 
				first.team, 
				second.team AS against, 
				first.season 
			FROM 
				(SELECT 
					matchID, 
					w, 
					r, 
					team, 
					team_id, 
					season 
				FROM 
					bowling 
				WHERE 
					player_id = %d AND season = %d
				ORDER BY 
					w DESC, 
					r ASC 
				LIMIT 1) AS first
			JOIN 
				(SELECT 
					team, 
					team_id, 
					matchID 
				FROM 
					scores) AS second 
			ON 
				first.matchID = second.matchID 
				AND first.team_id != second.team_id
		),
		stats_5 AS (
			SELECT
				SUM(b.w) AS totalWickets,
				SUM(b.r) AS totalRuns,
				CASE 
					WHEN SUM(b.b) = 0 THEN NULL
					ELSE SUM(b.r) / SUM(b.b) 
				END AS economy,
				CASE 
					WHEN SUM(b.w) = 0 THEN NULL
					ELSE SUM(b.r) / SUM(b.w) 
				END AS bowlingAverage
			FROM 
				bowling b
			JOIN 
				latest_5 lm ON b.matchID = lm.matchID
			WHERE 
				b.player_id = %d AND b.season = %d
		),
		stats_10 AS (
			SELECT
				SUM(b.w) AS totalWickets,
				SUM(b.r) AS totalRuns,
				CASE 
					WHEN SUM(b.b) = 0 THEN NULL
					ELSE SUM(b.r) / SUM(b.b) 
				END AS economy,
				CASE 
					WHEN SUM(b.w) = 0 THEN NULL
					ELSE SUM(b.r) / SUM(b.w) 
				END AS bowlingAverage
			FROM 
				bowling b
			JOIN 
				latest_10 lm ON b.matchID = lm.matchID
			WHERE 
				b.player_id = %d AND b.season = %d
		)
		SELECT 
			s.w,
			s.r,
			s.team,
			s.against,
			s.season,
			s5.totalWickets AS totalWickets_5,
			s5.totalRuns AS totalRuns_5,
			s5.economy AS economy_5,
			s5.bowlingAverage AS bowlingAverage_5,
			s10.totalWickets AS totalWickets_10,
			s10.totalRuns AS totalRuns_10,
			s10.economy AS economy_10,
			s10.bowlingAverage AS bowlingAverage_10
		FROM
			stats s
		JOIN
			stats_5 s5
		ON 1=1 
		JOIN
			stats_10 s10
		ON 1=1;
		`,
		id, season,
		id, season,
		id, season,
		id, season,
	)

	return GetBowlingPerformanceStats(query, db)
}

func GetSeasonPlayerCard(name string, season int) (PlayerCard, error) {
	var playerCard PlayerCard
	var err error

	playerCard.PlayerID, playerCard.PlayerName, err = GetPlayerByName(name, database.DB)

	playerCard.Season = season

	if err != nil {
		return playerCard, err
	}

	teams, err := GetSeasonTeams(playerCard.PlayerID, season, database.DB)

	playerCard.Team = teams[0]
	playerCard.OtherTeams = teams[1:]

	if err != nil {
		return playerCard, err
	}

	playerCard.CareerHistory, err = GetSeasonHistory(playerCard.PlayerID, season, database.DB)

	if err != nil {
		return playerCard, err
	}

	playerCard.Batting.PerformanceStats = GetSeasonBattingPerformance(playerCard.PlayerID, season, database.DB)

	playerCard.Batting.CareerStats = GetSeasonBattingStats(playerCard.PlayerID, season, database.DB)

	playerCard.Bowling.CareerStats = GetSeasonBowlingStats(playerCard.PlayerID, season, database.DB)

	playerCard.Bowling.PerformanceStats = GetSeasonBowlingPerformance(playerCard.PlayerID, season, database.DB)

	return playerCard, err
}
