package teams

import (
	"database/sql"
	"fmt"
	"ipl-api/database"
)

func GetAllTimeHistory(id int, db *sql.DB) (History, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM
	(SELECT COUNT(DISTINCT season) played FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d))) AS a
	JOIN
	(SELECT COUNT(DISTINCT season) playoffs FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND match_type NOT LIKE '%%match%%') AS b
	JOIN
	(SELECT COUNT(DISTINCT season) finals FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND match_type LIKE 'Final%%') AS c
	JOIN
	(SELECT COUNT(DISTINCT season) championships FROM matches WHERE winner_id=%d AND match_type LIKE 'Final%%') AS d;
	`, id, id,
		id, id,
		id, id,
		id,
	)

	return GetTeamHistory(query, db)

}

func GetAllTimeMatchRecord(id int, db *sql.DB) (MatchRecord, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM 
	(SELECT COUNT(*) played FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d))) AS a
	JOIN    
	(SELECT count(*) won FROM matches WHERE winner_id=%d) AS b
	JOIN    
	(SELECT count(*) lost FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND winner_id != %d) AS c
	JOIN    
	(SELECT count(*) nOResult FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND outcome LIKE '%%abandoned%%') AS d
	JOIN    
	(SELECT count(*) tied FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) AND outcome LIKE '%%tied%%') AS e;
	`, id, id,
		id,
		id, id, id,
		id, id,
		id, id,
	)

	return GetTeamMatchRecord(query, db)
}

func GetAllTimePerformanceStats(id int, db *sql.DB) (PerformanceStats, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM
	(SELECT SUM(total) totalRuns, SUM(wickets) totalWickets, ROUND(AVG(total)) avgRuns, ROUND(AVG(wickets)) avgWickets FROM scores WHERE team_id=%d) AS a
	JOIN
	(SELECT CONCAT(a.total, '/', a.wickets, ' (', a.overs, ' ov)', ' vs.', b.team) highestScore FROM (SELECT matchID, total, wickets, overs FROM scores WHERE team_id=%d) AS a JOIN (SELECT matchID, team FROM scores WHERE team_id != %d) AS b on a.matchID=b.matchID ORDER BY a.total DESC, a.wickets ASC LIMIT 1) AS b
	JOIN
	(SELECT CONCAT(b.total, '/', b.wickets, ' (', b.overs, ' ov)', ' vs. ', b.team) bestBowling FROM (SELECT matchID, team_id FROM scores WHERE team_id = %d) AS a JOIN (SELECT matchID, team, total, wickets, overs FROM scores WHERE team_id != %d) AS b ON a.matchID = b.matchID ORDER BY b.total ASC, b.wickets DESC, b.overs ASC LIMIT 1) AS c
	JOIN
	(SELECT ROUND(win * 100 / total, 1) overall FROM (SELECT COUNT(*) total FROM scores WHERE team_id=%d) AS a JOIN (SELECT COUNT(*) win FROM scores WHERE team_id=%d AND winner=1) AS b) AS d
	JOIN
	(SELECT ROUND(win * 100 / total, 1) batfirst FROM (SELECT COUNT(*) total FROM scores WHERE team_id=%d AND innings=1) AS a JOIN (SELECT COUNT(*) win FROM scores WHERE team_id=%d AND winner=1 AND innings=1) AS b) AS e
	JOIN
	(SELECT ROUND(win * 100 / total, 1) fieldfirst FROM (SELECT COUNT(*) total FROM scores WHERE team_id=%d AND innings=2) AS a JOIN (SELECT COUNT(*) win FROM scores WHERE team_id=%d AND winner=1 AND innings=2) AS b) AS f
	JOIN
	(SELECT CONCAT(COUNT(*), '/', 5) AS last5 FROM (SELECT matchID FROM scores WHERE team_id=%d AND winner=1) AS a JOIN (SELECT date, matchID FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) ORDER BY date DESC LIMIT 5) b on a.matchID=b.matchID) AS g
	JOIN
	(SELECT CONCAT(COUNT(*), '/', 10) AS last10 FROM (SELECT matchID FROM scores WHERE team_id=%d AND winner=1) AS a JOIN (SELECT date, matchID FROM matches WHERE ((team_1_id=%d) OR (team_2_id=%d)) ORDER BY date DESC LIMIT 10) b on a.matchID=b.matchID) AS h;
	`, id,
		id, id,
		id, id,
		id, id,
		id, id,
		id, id,
		id, id, id,
		id, id, id,
	)

	return GetPerformancestats(query, db)
}

func GetAllTimePlayerAchievements(id int, db *sql.DB) (PlayerAchievements, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM
	(SELECT a.player, CONCAT(a.r, '/', a.b) AS score, CONCAT('vs. ', b.team) AS against, a.season FROM (SELECT matchID, team, player, r, b, season FROM batting WHERE r = (SELECT max(r) FROM batting WHERE team_id=%d)) AS a JOIN (SELECT matchID, team FROM scores WHERE team_id != %d) AS b ON a.matchID=b.matchID) AS a
	JOIN
	(SELECT a.player, CONCAT(a.w, '/', a.r, ' (', a.o, ' ov)') AS figure, CONCAT('vs. ', b.team) AS against, a.season FROM (SELECT matchID, team, team_id, player, w, r, o, season FROM bowling WHERE team_id = %d AND w = (SELECT MAX(w) FROM bowling WHERE team_id = %d ORDER BY r LIMIT 1)) AS a JOIN (SELECT matchID, team FROM scores WHERE team_id != %d) AS b ON a.matchID = b.matchID LIMIT 1) AS b
	JOIN
	(SELECT player, count(*) centuries FROM batting WHERE r >= 100 AND team_id=%d GROUP BY player ORDER BY centuries DESC LIMIT 1) AS c
	JOIN
	(SELECT player, SUM(w) wickets FROM bowling WHERE team_id=%d GROUP BY player ORDER BY SUM(w) DESC LIMIT 1) AS d;
	`, id, id,
		id, id, id,
		id,
		id,
	)

	return GetPlayerAchievements(query, db)
}

func GetAllTimeStats(name string) (Team, error) {
	var team Team
	var err error

	team.TeamID, team.TeamName, err = GetTeamID(name, database.DB)

	if err != nil {
		return team, err
	}

	team.TeamHistory, err = GetAllTimeHistory(team.TeamID, database.DB)

	if err != nil {
		return team, err
	}

	team.TeamMatchRecord, err = GetAllTimeMatchRecord(team.TeamID, database.DB)

	if err != nil {
		return team, err
	}

	team.TeamPerformanceStats, err = GetAllTimePerformanceStats(team.TeamID, database.DB)

	if err != nil {
		return team, err
	}

	team.PlayerAchievements, err = GetAllTimePlayerAchievements(team.TeamID, database.DB)

	return team, err
}
