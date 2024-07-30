import requests
from bs4 import BeautifulSoup
import json
from tqdm import tqdm

with open("extraction/IPL.json") as file:
    SEASONS = json.load(file)


def retrieveMatchURLs(year: str) -> list[str]:
    scheduleURL = "{}/match-schedule-fixtures-and-results".format(SEASONS[year])
    r = requests.get(scheduleURL)
    html = BeautifulSoup(r.text, "lxml")

    div = html.find("div", class_="ds-mb-4")
    matchURLs = div.find_all("a")
    matchURLS = [matchURL["href"] for matchURL in matchURLs if "href" in matchURL.attrs]

    matchURLS = [
        "https://www.espncricinfo.com{}".format(matchURL)
        for matchURL in matchURLS
        if "full-scorecard" in matchURL
    ]

    return matchURLS


def retrieveAllMatchURLs() -> None:
    urls = []
    for key in tqdm(SEASONS.keys()):
        try:
            urls += retrieveMatchURLs(key)
        except:
            pass

    with open("extraction/matches.txt", "w") as file:
        for url in urls:
            file.write("{}\n".format(url.strip()))

retrieveAllMatchURLs()