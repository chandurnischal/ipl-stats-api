DROP TABLE IF EXISTS scores;
RENAME TABLE total TO scores;

UPDATE scores 
    SET team = 'Rising Pune Supergiant(s)' 
        WHERE 
            team LIKE '%Rising%';

UPDATE scores
    SET team = 'Royal Challengers Bengaluru (Bangalore)' 
        WHERE 
            team LIKE '%Bangalore%' OR team LIKE '%Bengaluru%'; 

UPDATE scores 
    SET team = 'Delhi Capitals (Delhi Daredevils)' 
        WHERE 
            team LIKE '%Delhi%';

UPDATE scores
    SET team = 'Punjab Kings (Kings XI Punjab)' 
        WHERE 
            team LIKE '%Punjab%';

ALTER TABLE scores
    MODIFY COLUMN innings TINYINT(1),
    MODIFY COLUMN overs DECIMAL(10, 2),
    MODIFY COLUMN total INT, 
    MODIFY COLUMN wickets INT,
    MODIFY COLUMN matchID BIGINT,
    ADD COLUMN team_id INT,
    ADD COLUMN season INT;
    
CREATE INDEX idx_matchID ON scores (matchID);

UPDATE scores
    JOIN teams ON scores.team = teams.team_name
        SET scores.team_id = teams.team_id;

UPDATE scores
    JOIN matches ON scores.matchID = matches.matchID
        SET scores.season = YEAR(matches.date);


UPDATE extras 
    SET team = 'Rising Pune Supergiant(s)' 
        WHERE 
            team LIKE '%Rising%';

UPDATE extras
    SET team = 'Royal Challengers Bengaluru (Bangalore)' 
        WHERE 
            team LIKE '%Bangalore%' OR team LIKE '%Bengaluru%'; 

UPDATE extras 
    SET team = 'Delhi Capitals (Delhi Daredevils)' 
        WHERE 
            team LIKE '%Delhi%';

UPDATE extras
    SET team = 'Punjab Kings (Kings XI Punjab)' 
        WHERE 
            team LIKE '%Punjab%';

ALTER TABLE extras
    MODIFY COLUMN innings TINYINT(1),
    MODIFY COLUMN b INT,
    MODIFY COLUMN lb INT,
    MODIFY COLUMN w INT,
    MODIFY COLUMN nb INT,
    MODIFY COLUMN matchID BIGINT,
    ADD COLUMN team_id INT,
    ADD COLUMN season INT;

CREATE INDEX idx_matchID ON extras (matchID);

UPDATE extras
    JOIN teams ON extras.team = teams.team_name
        SET extras.team_id = teams.team_id;

UPDATE extras
    JOIN matches ON extras.matchID = matches.matchID
        SET extras.season = YEAR(matches.date);

ALTER TABLE scores ADD COLUMN winner TINYINT(1);

UPDATE 
    scores AS s
JOIN 
    (
        SELECT 
            team_id,
            matchID,
            CASE
                WHEN total = MAX(total) OVER (PARTITION BY matchID) THEN 1
                ELSE 0
            END AS is_winner
        FROM 
            scores
    ) AS ranked
ON 
    s.team_id = ranked.team_id
    AND s.matchID = ranked.matchID
SET 
    s.winner = ranked.is_winner;

UPDATE matches
    JOIN scores ON scores.matchID=matches.matchID
        SET matches.winner_id=scores.team_id WHERE scores.winner=1; 

UPDATE matches 
    JOIN teams ON matches.winner_id=teams.team_id
        SET matches.winner=teams.team_name;

UPDATE matches SET winner="Tied" WHERE outcome LIKE '%tied%';
UPDATE matches SET winner_id=-1 WHERE winner="Tied";


UPDATE scores 
    JOIN (SELECT matchID, winner_id FROM matches WHERE winner_id=-1) AS a ON scores.matchID=a.matchID
        SET scores.winner=-1;