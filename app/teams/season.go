package teams

import (
	"database/sql"
	"fmt"
	"ipl-api/database"
)

func GetSeasonMatches(id, season int, db *sql.DB) (Matches, error) {
	query := fmt.Sprintf(`SELECT *, ROUND((won * 100) / played, 2) AS winPerc
	FROM
	(SELECT COUNT(*) played FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND season=%d) AS a
	JOIN
	(SELECT COUNT(*) won FROM matches WHERE winner_id=%d AND season=%d) AS b
	JOIN
	(SELECT COUNT(*) lost FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND winner_id!=%d AND winner_id !=-1 AND winner_id IS NOT NULL AND season=%d) AS c
	JOIN
	(SELECT COUNT(*) tied FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND winner_id=-1 AND season=%d) AS d
	JOIN
	(SELECT COUNT(*) nr FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND winner_id IS NULL AND season=%d) AS e
	JOIN
	(SELECT ROUND(won * 100 / (won + lost), 2) firstBatWinPerc FROM (SELECT COUNT(*) AS won FROM scores WHERE innings=1 AND team_id=%d AND winner=1 AND season=%d) AS w JOIN (SELECT COUNT(*) AS lost FROM scores WHERE innings=1 AND team_id=%d AND winner=0 AND season=%d) AS l) AS f
	JOIN
	(SELECT ROUND(won * 100 / (won + lost), 2) firstFieldWinPerc FROM (SELECT COUNT(*) AS won FROM scores WHERE innings=2 AND team_id=%d AND winner=1 AND season=%d) AS w JOIN (SELECT COUNT(*) AS lost FROM scores WHERE innings=2 AND team_id=%d AND winner=0 AND season=%d) AS l) AS g`,
		id, id, season,
		id, season,
		id, id, id, season,
		id, id, season,
		id, id, season,
		id, season, id, season,
		id, season, id, season,
	)

	return GetMatches(query, db)

}

func GetSeasonAppearances(id, season int, db *sql.DB) (Appearances, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM	
	(SELECT COUNT(DISTINCT season) played FROM matches WHERE team_1_id=%d AND season=%d) as a
	JOIN
	(SELECT COUNT(DISTINCT season) appearances FROM matches WHERE match_type NOT LIKE '%%match%%' AND ((team_1_id=%d) OR (team_2_id=%d)) AND season=%d) AS b
	JOIN
	(SELECT COUNT(DISTINCT season) finals FROM matches WHERE match_type LIKE 'Final%%' AND ((team_1_id=%d) OR (team_2_id=%d)) AND season=%d) AS c
	JOIN
	(SELECT COUNT(DISTINCT season) championships FROM matches WHERE match_type LIKE 'Final%%' AND winner_id=%d AND season=%d) as d;
	`, id, season,
		id, id, season,
		id, id, season,
		id, season)

	return GetAppearances(query, db)
}

func GetSeasonIndividualPerformance(id, season int, db *sql.DB) (Indiviudal, error) {
	query := fmt.Sprintf(`
	SELECT *
	FROM
	(SELECT player AS batsman, SUM(r) AS mostRuns FROM batting WHERE team_id=%d AND season=%d GROUP BY player ORDER BY mostRuns DESC LIMIT 1) AS a
	JOIN
	(SELECT player AS bowler, SUM(w) as mostWickets FROM bowling WHERE team_id=%d AND season=%d GROUP BY player ORDER BY mostWickets DESC LIMIT 1) AS b
	JOIN
	(SELECT CONCAT(CONVERT(r, CHAR), '/', CONVERT(b, CHAR), ' (', player, ')') AS highestScore FROM batting WHERE r=(SELECT MAX(r) FROM batting WHERE team_id=%d AND season=%d ORDER BY b ASC) AND team_id=%d AND season=%d LIMIT 1) AS c
	JOIN
	(SELECT CONCAT(CONVERT(w, CHAR), '/', CONVERT(r, CHAR), ' (', player, ')') bestBowling FROM bowling WHERE w=(SELECT MAX(w) FROM bowling WHERE team_id=%d AND season=%d ORDER BY r) AND team_id=%d AND season=%d LIMIT 1) AS d;
	`, id, season,
		id, season,
		id, season, id, season,
		id, season, id, season,
	)

	return GetIndividualPerformance(query, db)

}

func GetSeasonStats(id, season int, db *sql.DB) (Stats, error) {
	query := fmt.Sprintf(`
	SELECT *
	FROM
	(SELECT CONCAT(CONVERT(total, CHAR), '/', CONVERT(wickets, CHAR)) AS highestScore FROM scores WHERE team_id=%d AND season=%d AND total=(SELECT MAX(total) FROM scores WHERE team_id=%d AND season=%d)) AS a
	JOIN
	(SELECT CONCAT(CONVERT(total, CHAR), '/', CONVERT(wickets, CHAR)) AS lowestScore FROM scores WHERE team_id=%d AND season=%d AND total=(SELECT MIN(total) FROM scores WHERE team_id=%d AND season=%d)) AS b
	JOIN
	(SELECT ROUND(AVG(total)) averageScore FROM scores WHERE team_id=%d AND season=%d) AS c
	JOIN
	(SELECT ROUND(AVG(wickets)) averageWickets FROM scores WHERE team_id=%d AND season=%d) AS d
	JOIN
	(SELECT SUM(total) totalRuns FROM scores WHERE team_id=%d AND season=%d) AS e
	JOIN
	(SELECT SUM(wickets) totalWickets FROM scores WHERE team_id=%d AND season=%d) AS f
	`, id, season, id, season,
		id, season, id, season,
		id, season,
		id, season,
		id, season,
		id, season,
	)
	return GetTeamStats(query, db)
}

func GetSeasonData(name string, season int) (Team, error) {
	var team Team
	var err error

	team.TeamID, team.TeamName, err = GetTeamID(name, database.DB)
	team.Season = season
	if err != nil {
		return team, err
	}

	team.Matches, err = GetSeasonMatches(team.TeamID, season, database.DB)

	if err != nil {
		return team, err
	}

	team.Appearances, err = GetSeasonAppearances(team.TeamID, season, database.DB)

	if err != nil {
		return team, err
	}

	team.IndiviudalPerformance, err = GetSeasonIndividualPerformance(team.TeamID, season, database.DB)

	if err != nil {
		return team, err
	}

	team.Stats, err = GetSeasonStats(team.TeamID, season, database.DB)

	return team, err

}
