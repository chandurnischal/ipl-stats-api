SELECT 
    a.player, 
    a.team, 
    b.team AS against, 
    a.season,
    a.r, 
    a.b, 
    a.r * 100 / a.b AS strikerate, 
    a.4s, 
    a.6s 
FROM 
(
    SELECT 
        matchID, 
        team, 
        player,
        season, 
        r, 
        b, 
        4s, 
        6s 
    FROM 
        batting WHERE season = 2024 
    ORDER BY 
        r DESC, b ASC
    LIMIT 1
) AS a 
JOIN 
(
    SELECT 
        matchID, 
        team 
    FROM 
        scores
) AS b 
ON 
    a.matchID = b.matchID 
    AND a.team != b.team;