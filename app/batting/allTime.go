package batting

import (
	"database/sql"
	"fmt"
	"ipl-api/database"
)

func GetAllTimeCareerStats(id int, db *sql.DB) (CareerStats, error) {
	query := fmt.Sprintf(`
	SELECT matchesPlayed, innings, totalRuns, totalBalls, centuries, halfCenturies, fours, sixes, ROUND((totalRuns * 100) / totalBalls, 2) AS strikeRate, ROUND(totalRuns / dismissals, 2) battingAverage, ROUND((fours + sixes) / innings) AS boundaryPerInnings, highestScore, ducks
	FROM
	(SELECT COUNT(*) matchesPlayed FROM (SELECT matchID, team_id FROM batting WHERE player_id = %d UNION SELECT matchID, team_id FROM bowling WHERE player_id = %d) AS a JOIN (SELECT matchID, team_id FROM scores) AS b ON a.matchID = b.matchID AND a.team_id = b.team_id) AS a
	JOIN
	(SELECT COUNT(*) innings FROM batting WHERE player_id=%d) AS b
	JOIN
	(SELECT SUM(r) totalRuns FROM batting WHERE player_id=%d) AS c
	JOIN
	(SELECT SUM(b) totalBalls FROM batting WHERE player_id=%d) AS d
	JOIN
	(SELECT COUNT(*) centuries FROM batting WHERE r >= 100 AND player_id=%d) AS e
	JOIN
	(SELECT COUNT(*) halfCenturies FROM batting WHERE r >= 50 AND r < 100 AND player_id=%d) AS f
	JOIN
	(SELECT SUM(4s) fours FROM batting WHERE player_id=%d) AS g
	JOIN
	(SELECT SUM(6s) sixes FROM batting WHERE player_id=%d) AS h
	JOIN
	(SELECT COUNT(*) dismissals FROM batting WHERE player_id=%d AND dismissal_info NOT LIKE '%%not out%%') AS i
	JOIN
	(SELECT CONCAT(CONVERT(r, CHAR), '/', CONVERT(b, CHAR)) highestScore FROM batting WHERE player_id=%d AND r = (SELECT MAX(r) FROM batting WHERE player_id=%d) ORDER BY b LIMIT 1) AS j
	JOIN
	(SELECT COUNT(*) ducks FROM batting WHERE player_id=%d AND r=0) AS k;
	`, id, id,
		id,
		id,
		id,
		id,
		id,
		id,
		id,
		id,
		id, id,
		id,
	)

	return GetCareerStats(query, db)
}

func GetAllTimeBattingStats(name string) (BattingCard, error) {
	var batting BattingCard
	var err error
	batting.PlayerID, batting.PlayerName, err = GetPlayerID(name, database.DB)

	if err != nil {
		return batting, err
	}

	teams, err := GetTeams(batting.PlayerID, database.DB)

	if err != nil {
		return batting, err
	}

	batting.CurrentTeam = teams[0]
	batting.PreviousTeams = teams[1:]

	batting.Stats, err = GetAllTimeCareerStats(batting.PlayerID, database.DB)

	if err != nil {
		return batting, err
	}

	return batting, err

}
