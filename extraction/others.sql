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

UPDATE matches
    JOIN winners ON matches.matchID=winners.matchID
        SET matches.winner=winners.winner;

UPDATE matches
    JOIN teams ON matches.winner=teams.team_name
        SET matches.winner_id=teams.team_id;