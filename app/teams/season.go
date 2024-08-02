package teams

import (
	"database/sql"
	"fmt"
	"ipl-api/database"
)

func GetSeasonHistory(id, season int, db *sql.DB) (History, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM
	(SELECT COUNT(DISTINCT season) played FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND season=%d) AS a
	JOIN
	(SELECT COUNT(DISTINCT season) playoffs FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND match_type NOT LIKE '%%match%%' AND season=%d) AS b
	JOIN
	(SELECT COUNT(DISTINCT season) finals FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND match_type LIKE 'Final%%' AND season=%d) AS c
	JOIN
	(SELECT COUNT(DISTINCT season) championships FROM matches WHERE winner_id=%d AND match_type LIKE 'Final%%' AND season=%d) AS d;

	`, id, id, season,
		id, id, season,
		id, id, season,
		id, season,
	)

	return GetTeamHistory(query, db)

}

func GetSeasonMatchRecord(id, season int, db *sql.DB) (MatchRecord, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM 
	(SELECT COUNT(*) played FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND season=%d) AS a
	JOIN    
	(SELECT count(*) won FROM matches WHERE winner_id=%d AND season=%d) AS b
	JOIN    
	(SELECT count(*) lost FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND winner_id != %d AND season=%d) AS c
	JOIN    
	(SELECT count(*) nOResult FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND outcome LIKE '%%abandoned%%' AND season=%d) AS d
	JOIN    
	(SELECT count(*) tied FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND outcome LIKE '%%tied%%' AND season=%d) AS e;
	`, id, id, season,
		id, season,
		id, id, id, season,
		id, id, season,
		id, id, season,
	)

	return GetTeamMatchRecord(query, db)
}

func GetSeasonPerformanceStats(id, season int, db *sql.DB) (PerformanceStats, error) {
	query := fmt.Sprintf(`
	SELECT * FROM 
	(SELECT SUM(total) AS totalRuns, SUM(wickets) AS totalWickets, ROUND(AVG(total)) AS avgRuns, ROUND(AVG(wickets)) AS avgWickets FROM scores WHERE team_id = %d AND season = %d) AS a 
	JOIN 
	(SELECT CONCAT(a.total, '/', a.wickets, ' (', a.overs, ' ov) vs. ', b.team) AS highestScore FROM (SELECT matchID, total, wickets, overs FROM scores WHERE team_id = %d AND season = %d) AS a JOIN (SELECT matchID, team FROM scores WHERE team_id != %d) AS b ON a.matchID = b.matchID ORDER BY a.total DESC, a.wickets ASC LIMIT 1) AS b 
	JOIN 
	(SELECT CONCAT(b.total, '/', b.wickets, ' (', b.overs, ' ov) vs. ', b.team) AS bestBowling FROM (SELECT matchID FROM scores WHERE team_id = %d AND season = %d) AS a JOIN (SELECT matchID, team, total, wickets, overs FROM scores WHERE team_id != %d) AS b ON a.matchID = b.matchID ORDER BY b.total ASC, b.wickets DESC, b.overs ASC LIMIT 1) AS c 
	JOIN 
	(SELECT ROUND(win * 100.0 / total, 1) AS overall FROM (SELECT COUNT(*) AS total FROM scores WHERE team_id = %d AND season = %d) AS a JOIN (SELECT COUNT(*) AS win FROM scores WHERE team_id = %d AND winner = 1 AND season = %d) AS b) AS d
	JOIN 
	(SELECT ROUND(win * 100.0 / total, 1) AS batfirst FROM (SELECT COUNT(*) AS total FROM scores WHERE team_id = %d AND innings = 1 AND season = %d) AS a JOIN (SELECT COUNT(*) AS win FROM scores WHERE team_id = %d AND winner = 1 AND innings = 1 AND season = %d) AS b) AS e
	JOIN 
	(SELECT ROUND(win * 100.0 / total, 1) AS fieldfirst FROM (SELECT COUNT(*) AS total FROM scores WHERE team_id = %d AND innings = 2 AND season = %d) AS a JOIN (SELECT COUNT(*) AS win FROM scores WHERE team_id = %d AND winner = 1 AND innings = 2 AND season = %d) AS b) AS f 
	JOIN 
	(SELECT CONCAT(COUNT(*), '/', 5) AS last5 FROM (SELECT matchID FROM scores WHERE team_id = %d AND winner = 1 AND season = %d) AS a JOIN (SELECT date, matchID FROM matches WHERE (team_1_id = %d OR team_2_id = %d) AND season = %d ORDER BY date DESC LIMIT 5) AS b ON a.matchID = b.matchID) AS g
	JOIN (SELECT CONCAT(COUNT(*), '/', 10) AS last10 FROM (SELECT matchID FROM scores WHERE team_id = %d AND winner = 1 AND season = %d) AS a JOIN (SELECT date, matchID FROM matches WHERE (team_1_id = %d OR team_2_id = %d) AND season = %d ORDER BY date DESC LIMIT 10) AS b ON a.matchID = b.matchID) AS h;
	`, id, season,
		id, season, id,
		id, season, id,
		id, season, id, season,
		id, season, id, season,
		id, season, id, season,
		id, season,
		id, id, season,
		id, season,
		id, id, season,
	)

	return GetPerformancestats(query, db)
}

func GetSeasonPlayerAchievements(id, season int, db *sql.DB) (PlayerAchievements, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM
	(SELECT a.player, CONCAT(a.r, '/', a.b) AS score, CONCAT('vs. ', b.team) AS against, a.season FROM (SELECT matchID, team, player, r, b, season FROM batting WHERE r = (SELECT max(r) FROM batting WHERE team_id=%d AND season=%d)AND season=%d) AS a JOIN (SELECT matchID, team FROM scores WHERE team_id != %d) AS b ON a.matchID=b.matchID) AS a
	JOIN
	(SELECT a.player, CONCAT(a.w, '/', a.r, ' (', a.o, ' ov)') AS figure, CONCAT('vs. ', b.team) AS against, a.season FROM (SELECT matchID, team, team_id, player, w, r, o, season FROM bowling WHERE team_id = %d AND w = (SELECT MAX(w) FROM bowling WHERE team_id = %d AND season=%d ORDER BY r LIMIT 1) AND season=%d) AS a JOIN (SELECT matchID, team FROM scores WHERE team_id != %d) AS b ON a.matchID = b.matchID LIMIT 1) AS b
	JOIN
	(SELECT player, count(*) centuries FROM batting WHERE r >= 100 AND team_id=%d AND season=%d GROUP BY player ORDER BY centuries DESC LIMIT 1) AS c
	JOIN
	(SELECT player, SUM(w) wickets FROM bowling WHERE team_id=%d AND season=%d GROUP BY player ORDER BY SUM(w) DESC LIMIT 1) AS d;
	`, id, season, season, id,
		id, id, season, season, id,
		id, season,
		id, season,
	)

	return GetPlayerAchievements(query, db)
}

func GetSeasonStats(name string, season int) (Team, error) {
	var team Team
	var err error

	team.TeamID, team.TeamName, err = GetTeamID(name, database.DB)

	if err != nil {
		return team, err
	}

	team.TeamHistory, err = GetSeasonHistory(team.TeamID, season, database.DB)

	if err != nil {
		return team, err
	}

	team.TeamMatchRecord, err = GetSeasonMatchRecord(team.TeamID, season, database.DB)

	if err != nil {
		return team, err
	}

	team.TeamPerformanceStats, err = GetSeasonPerformanceStats(team.TeamID, season, database.DB)

	if err != nil {
		return team, err
	}

	team.PlayerAchievements, err = GetSeasonPlayerAchievements(team.TeamID, season, database.DB)

	team.Season = season
	return team, err
}
