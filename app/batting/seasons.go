package batting

import (
	"database/sql"
	"fmt"
	"ipl-api/database"
)

func GetSeasonPlayerHistory(id, season int, db *sql.DB) (History, error) {

	query := fmt.Sprintf(
		`
		SELECT *
		FROM
		(SELECT COUNT(*) matches FROM batting WHERE player_id=%d) AS a
		JOIN
		(SELECT COUNT(DISTINCT season) seasons FROM batting WHERE player_id=%d) AS b
		JOIN
		(SELECT COUNT(DISTINCT first.season) finals FROM (SELECT DISTINCT team_id, season FROM batting WHERE player_id=%d) AS first JOIN (SELECT team_1_id, team_2_id, season FROM matches WHERE match_type not like '%%match%%') AS second ON (first.team_id=second.team_1_id OR first.team_id=second.team_2_id) AND first.season=second.season) AS c
		JOIN
		(SELECT COUNT(*) finals FROM (SELECT DISTINCT team_id, season FROM batting WHERE player_id=%d) AS first JOIN (SELECT team_1_id, team_2_id, season FROM matches WHERE match_type like 'Final%%') AS second ON (first.team_id=second.team_1_id OR first.team_id=second.team_2_id) AND first.season=second.season) AS d
		JOIN
		(SELECT COUNT(*) championships FROM (SELECT DISTINCT team_id, season FROM batting WHERE player_id=%d) AS first JOIN (SELECT winner_id, season FROM matches WHERE match_type like 'Final%%') AS second ON first.team_id=second.winner_id AND first.season=second.season) AS e;
		`, id,
		id,
		id,
		id,
		id,
	)

	return GetPlayerHistory(query, db)
}

func GetSeasonBattingCard(name string, season int) (BattingCard, error) {
	var batting BattingCard
	var err error

	batting.PlayerID, batting.PlayerName, err = GetPlayerID(name, database.DB)

	if err != nil {
		return batting, err

	}

	batting.TournamentHistory, err = GetSeasonPlayerHistory(batting.PlayerID, season, database.DB)

	return batting, err
}
