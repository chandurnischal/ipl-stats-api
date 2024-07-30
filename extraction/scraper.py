import json
import pandas as pd
from sqlalchemy import create_engine
import scorecard as s
from time import sleep
from random import randint
from tqdm import tqdm



with open("creds.json") as file:
    creds = json.load(file)

engine = create_engine("mysql+mysqlconnector://{}:{}@{}/{}".format(creds["username"], creds["password"], creds["host"], creds["database"]))

def pushToDB(data:pd.DataFrame, tablename:str) -> None:
    data.to_sql(name=tablename, con=engine, if_exists="append", index=False)

with open("extraction/matches.txt") as file:
    urls = file.readlines()


with open("extraction/errors.csv", "a") as error:
    for url in tqdm(urls):
        url = url.strip()
        try:
            scorecard = s.extractScorecard(url)
        except:
            error.write("{}, {}\n".format(url, "failed to extract scorecard"))
        try:


            pushToDB(scorecard["match_details"], tablename="matches")
            pushToDB(scorecard["batting"], tablename="batting")
            pushToDB(scorecard["bowling"], tablename="bowling")
            pushToDB(scorecard["total"], tablename="total")
            pushToDB(scorecard["extras"], tablename="extras")
            sleep(randint(2, 4))

        except Exception as e:
            error.write("{}, {}\n".format(url, e.__class__))