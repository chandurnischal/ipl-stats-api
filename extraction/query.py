import re
import pandas as pd
import json
from tqdm import tqdm
import mysql.connector as mc

def extractQueries(file_path):
    with open(file_path, 'r') as file:
        content = file.read()

    content = re.sub(r'--.*', '', content)
    content = re.sub(r'/\*.*?\*/', '', content, flags=re.DOTALL)

    queries = [query.strip() + ';' for query in content.split(';') if query.strip()]

    return queries

def getAbbreviations(string:str) -> str:
    string = re.sub(r"\(.+", "", string)
    string = string.strip()

    return "".join([word[0] for word in string.split(' ')])

def updateWithAbbrev(row):
    if row["winner"] == None:
        return row["winner"]
    if row["team_1_abbrev"] in row["winner"]:
        return row["team_1"]
    if row["team_2_abbrev"] in row["winner"]:
        return row["team_2"]
    return row["winner"]

def updateWithNames(row):
    if row["winner"] == None:
        return row["winner"]
    if row["winner"] in row["team_1"]:
        return row["team_1"]
    if row["winner"] in row["team_2"]:
        return row["team_2"]
    return row["winner"]

matches = extractQueries("extraction/matches.sql")
batting = extractQueries("extraction/batting.sql")
bowling = extractQueries("extraction/bowling.sql")
others = extractQueries("extraction/others.sql")


with open("creds.json") as file:
    creds = json.load(file)

with mc.connect(**creds) as conn:
    cursor = conn.cursor()

    conn.autocommit = True

    for query in tqdm(matches, desc="Processing matches..."):
        cursor.execute(query)

    cursor.execute("SELECT matchID, team_1, team_2, winner FROM matches")
    rows = cursor.fetchall()

    data = pd.DataFrame(rows, columns=['matchID', 'team_1', 'team_2', 'winner'])

    data['winner'] = data['winner'].str.replace(r'.+\(', '', regex=True).str.strip()
    data = data.replace("Guj Lions", "Gujarat Lions")
    data = data.replace("Supergiants", "Supergiant")
    data = data.replace("PBKS", "Punjab")
    data = data.replace("SRH", "SH")


    data["team_1_abbrev"] = data["team_1"].apply(getAbbreviations)
    data["team_2_abbrev"] = data["team_2"].apply(getAbbreviations)
    data["winner"] = data.apply(updateWithAbbrev, axis=1)
    data["winner"] = data.apply(updateWithNames, axis=1)

    data = data[["matchID", "team_1", "team_2", "winner"]]

    insertStatement = """
    INSERT INTO winners (matchID, team_1, team_2, winner) VALUES (%s, %s, %s, %s)
    """
    
    for index, row in tqdm(data.iterrows(), total=data.shape[0], desc="Processing winners..."):
        cursor.execute(insertStatement, (row['matchID'], row["team_1"], row["team_2"], row["winner"])) 


    for query in tqdm(batting, desc="Processing batting..."):
        cursor.execute(query)

    for query in tqdm(bowling, desc="Processing bowling..."):
        cursor.execute(query)

    for query in tqdm(others, desc="Processing scores and extras"):
        cursor.execute(query)