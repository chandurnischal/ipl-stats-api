import re
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

matches = extractQueries("extraction/matches.sql")
batting = extractQueries("extraction/batting.sql")
bowling = extractQueries("extraction/bowling.sql")
others = extractQueries("extraction/others.sql")


with open("creds.json") as file:
    creds = json.load(file)

with mc.connect(**creds) as conn:
    cursor = conn.cursor()

    conn.autocommit = True

    print("Processing matches...")
    for query in tqdm(matches):
        cursor.execute(query)

    print("Processing batting...")
    for query in tqdm(batting):
        cursor.execute(query)

    print("Processing bowling...")
    for query in tqdm(bowling):
        cursor.execute(query)

    print("Processing scores and extras...")
    for query in tqdm(others):
        cursor.execute(query)
