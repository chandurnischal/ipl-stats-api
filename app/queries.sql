SELECT *
FROM 
(SELECT COUNT(DISTINCT b.season) AS playoffAppearances FROM (SELECT matchID FROM batting WHERE player_id = 20 UNION SELECT matchID FROM bowling WHERE player_id = 20) AS a JOIN (SELECT matchID, season FROM matches WHERE match_type NOT LIKE '%match%') AS b ON a.matchID = b.matchID) AS b
JOIN
(SELECT COUNT(DISTINCT b.season) AS finalAppearances FROM (SELECT matchID FROM batting WHERE player_id = 20 UNION SELECT matchID FROM bowling WHERE player_id = 20) AS a JOIN (SELECT matchID, season FROM matches WHERE match_type LIKE 'Final%') AS b ON a.matchID = b.matchID) AS c
JOIN
(SELECT DISTINCT b.season AS championshipsWon FROM (SELECT matchID, team_id FROM batting WHERE player_id = 20 UNION SELECT matchID, team_id FROM bowling WHERE player_id = 20) AS a JOIN (SELECT matchID, season, winner_id FROM matches WHERE match_type LIKE 'Final%') AS b ON a.matchID = b.matchID AND b.winner_id=a.team_id) AS d;