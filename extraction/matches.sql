/*

To dump:
mysqldump --default-character-set=utf8 -u root -p ipl > backup.sql


To load: 
Get-Content backup.sql | mysql --user=root --password=root --binary-mode=1 ipl
*/

ALTER TABLE matches 
    MODIFY COLUMN date DATE,
    MODIFY COLUMN season INT,
    MODIFY COLUMN matchID BIGINT,
    ADD COLUMN winner TEXT,
    ADD COLUMN winner_id INT;

UPDATE matches 
    SET team_1 = 'Rising Pune Supergiant(s)' 
        WHERE 
            team_1 LIKE '%Rising%';

UPDATE matches
    SET team_2 = 'Rising Pune Supergiant(s)' 
        WHERE 
            team_2 LIKE '%Rising%'; 

UPDATE matches
    SET team_1 = 'Royal Challengers Bengaluru (Bangalore)' 
        WHERE 
            team_1 LIKE '%Bangalore%' OR team_1 LIKE '%Bengaluru%'; 

UPDATE matches
    SET team_2 = 'Royal Challengers Bengaluru (Bangalore)' 
        WHERE 
            team_2 LIKE '%Bangalore%' OR team_2 LIKE '%Bengaluru%';

UPDATE matches 
    SET team_1 = 'Delhi Capitals (Delhi Daredevils)' 
        WHERE 
            team_1 LIKE '%Delhi%';

UPDATE matches
    SET team_2 = 'Delhi Capitals (Delhi Daredevils)' 
        WHERE 
            team_2 LIKE '%Delhi%'; 

UPDATE matches
    SET team_1 = 'Punjab Kings (Kings XI Punjab)' 
        WHERE 
            team_1 LIKE '%Punjab%';

UPDATE matches
    SET team_2 = 'Punjab Kings (Kings XI Punjab)' 
        WHERE 
            team_2 LIKE '%Punjab%';


DROP TABLE IF EXISTS teams;

CREATE TABLE teams (
    team_id INT AUTO_INCREMENT PRIMARY KEY,
    team_name TEXT NOT NULL
);

INSERT INTO teams (team_name)
    SELECT DISTINCT team_1
        FROM matches
            ORDER BY team_1;

ALTER TABLE matches 
    ADD COLUMN team_1_id INT, 
    ADD COLUMN team_2_id INT;

UPDATE matches
    JOIN teams ON matches.team_1 = teams.team_name
        SET matches.team_1_id = teams.team_id;

UPDATE matches
    JOIN teams ON matches.team_2 = teams.team_name
        SET matches.team_2_id = teams.team_id;

CREATE INDEX idx_matchID ON matches (matchID);
