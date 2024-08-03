UPDATE bowling 
    SET team = 'Rising Pune Supergiant(s)' 
        WHERE 
            team LIKE '%Rising%';

UPDATE bowling
    SET team = 'Royal Challengers Bengaluru (Bangalore)' 
        WHERE 
            team LIKE '%Bangalore%' OR team LIKE '%Bengaluru%'; 

UPDATE bowling 
    SET team = 'Delhi Capitals (Delhi Daredevils)' 
        WHERE 
            team LIKE '%Delhi%';

UPDATE bowling
    SET team = 'Punjab Kings (Kings XI Punjab)' 
        WHERE 
            team LIKE '%Punjab%';

UPDATE bowling SET o = NULL WHERE o LIKE '%-%';
UPDATE bowling SET m = NULL WHERE m LIKE '%-%';
UPDATE bowling SET r = NULL WHERE r LIKE '%-%';
UPDATE bowling SET w = NULL WHERE w LIKE '%-%';
UPDATE bowling SET econ = NULL WHERE econ LIKE '%-%';
UPDATE bowling SET 0s = NULL WHERE 0s LIKE '%-%';
UPDATE bowling SET 4s = NULL WHERE 4s LIKE '%-%';
UPDATE bowling SET 6s = NULL WHERE 6s LIKE '%-%';
UPDATE bowling SET wd = NULL WHERE wd LIKE '%-%';
UPDATE bowling SET nb = NULL WHERE nb LIKE '%-%';
UPDATE bowling SET player = REGEXP_REPLACE(player, '[^a-zA-Z0-9 ]', '');



ALTER TABLE bowling
    MODIFY COLUMN o DECIMAL(10, 2),
    MODIFY COLUMN m INT,
    MODIFY COLUMN r INT,
    MODIFY COLUMN w INT,
    MODIFY COLUMN econ DECIMAL(10, 2),
    MODIFY COLUMN 0s INT,
    MODIFY COLUMN 4s INT,
    MODIFY COLUMN 6s INT,
    MODIFY COLUMN wd INT,
    MODIFY COLUMN nb INT,
    MODIFY COLUMN innings INT,
    MODIFY COLUMN matchID BIGINT,
    ADD COLUMN team_id INT,
    ADD COLUMN season INT,
    ADD COLUMN player_id INT,
    ADD COLUMN b INT;

UPDATE bowling SET b = (FLOOR(o) * 6) + (o - FLOOR(o));
UPDATE bowling
    JOIN teams ON bowling.team = teams.team_name
        SET bowling.team_id = teams.team_id;

CREATE INDEX idx_matchID ON bowling (matchID);

UPDATE bowling
    JOIN matches ON bowling.matchID = matches.matchID
        SET bowling.season = YEAR(matches.date);



DROP TABLE IF EXISTS players;

CREATE TABLE players (
    player_id INT AUTO_INCREMENT PRIMARY KEY,
    player_name TEXT NOT NULL
);


INSERT INTO players (player_name)
    SELECT DISTINCT player FROM batting
        UNION
            SELECT DISTINCT player FROM bowling;

UPDATE batting 
    JOIN players ON batting.player = players.player_name
        SET batting.player_id = players.player_id;

UPDATE bowling 
    JOIN players ON bowling.player = players.player_name
        SET bowling.player_id = players.player_id;