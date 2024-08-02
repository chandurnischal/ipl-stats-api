import json
import pandas as pd
from sqlalchemy import create_engine
import re
from tqdm import tqdm 
from random import randint
import scorecard as s
from time import sleep


def getID(url):
    res = ""
    pattern = re.compile(r"(\d+)/full-scorecard", re.I)

    for m in pattern.finditer(url):
        res = m.group(1).strip()

    return res



with open("creds.json") as file:
    creds = json.load(file)

engine = create_engine("mysql+mysqlconnector://{}:{}@{}/{}".format(creds["username"], creds["password"], creds["host"], creds["database"]))

def pushToDB(data:pd.DataFrame, tablename:str) -> None:
    data.to_sql(name=tablename, con=engine, if_exists="append", index=False)


with open("extraction/matches.txt") as file:
    urls = file.readlines()


with open("extraction/errors.csv", "a") as error:
    error.write("id,url,error\n")
    for url in tqdm(urls):
        url = url.strip()
        id = getID(url)

        try:
            scorecard = s.extractScorecard(url)
        except:
            error.write("{},{}, failed to extract scorecard\n".format(id, url))

        try:
            pushToDB(scorecard["match_details"], tablename="matches")
            pushToDB(scorecard["batting"], tablename="batting")
            pushToDB(scorecard["bowling"], tablename="bowling")
            pushToDB(scorecard["total"], tablename="scores")
            pushToDB(scorecard["extras"], tablename="extras")
            sleep(randint(5, 10))
            error.write("{},{}, None\n".format(id, url))
        except Exception as e:
            error.write("{},{},{}\n".format(id, url, e.__class__)) 
