UPDATE batting 
    SET team = 'Rising Pune Supergiant(s)' 
        WHERE 
            team LIKE '%Rising%';

UPDATE batting
    SET team = 'Royal Challengers Bengaluru (Bangalore)' 
        WHERE 
            team LIKE '%Bangalore%' OR team LIKE '%Bengaluru%'; 

UPDATE batting 
    SET team = 'Delhi Capitals (Delhi Daredevils)' 
        WHERE 
            team LIKE '%Delhi%';

UPDATE batting
    SET team = 'Punjab Kings (Kings XI Punjab)' 
        WHERE 
            team LIKE '%Punjab%';

UPDATE batting SET r = NULL WHERE r LIKE '%-%';
UPDATE batting SET b = NULL WHERE b LIKE '%-%';
UPDATE batting SET m = NULL WHERE m LIKE '%-%';
UPDATE batting SET 4s = NULL WHERE 4s LIKE '%-%';
UPDATE batting SET 6s = NULL WHERE 6s LIKE '%-%';
UPDATE batting SET sr = NULL WHERE sr LIKE '%-%';
UPDATE batting SET player = REGEXP_REPLACE(player, '[^a-zA-Z0-9 ]', '');


ALTER TABLE batting
    MODIFY COLUMN r INT,
    MODIFY COLUMN b INT,
    MODIFY COLUMN m INT,
    MODIFY COLUMN 4s INT,
    MODIFY COLUMN 6s INT,
    MODIFY COLUMN sr DECIMAL(10, 2),
    MODIFY COLUMN matchID BIGINT,
    MODIFY COLUMN innings INT,
    ADD COLUMN team_id INT,
    ADD COLUMN season INT,
    ADD COLUMN player_id INT;

UPDATE batting
    JOIN teams ON batting.team = teams.team_name
        SET batting.team_id = teams.team_id;

CREATE INDEX idx_matchID ON batting (matchID);

UPDATE batting
    JOIN matches ON batting.matchID = matches.matchID
        SET batting.season = YEAR(matches.date);
