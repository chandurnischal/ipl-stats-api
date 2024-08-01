import requests
from bs4 import BeautifulSoup
import re
import pandas as pd

def removeNonASCII(text):
    return re.sub(r"[^\x00-\x7F]+", "", text)


def extractMatchID(url: str) -> int:
    try:
        id = ""
        matchNumber = re.compile(r"\-(\d+)/full-scorecard")

        for m in matchNumber.finditer(url):
            id += m.group(1).strip()

        return id
    except:
        return None

def resultDetails(result:str) -> dict:
    result = result.split(', ')
    res = dict()

    res["match_type"] = result[0].strip()
    res["venue"] = result[1].strip()
    res["date"] = result[2].strip()

    index = res["date"].find('-')

    if index != -1:
        res["date"] = res["date"][:index].strip()


    res["season"] = result[3].strip()
    res["date"] = "{}, {}".format(res["date"], res["season"]) 

    return res


def teamsPlayed(teams: BeautifulSoup) -> dict:
    titles = teams.find_all("a", {"title": True})
    titles = [title["title"] for title in titles]

    return {"team_1": titles[0], "team_2": titles[1]}


def getMatchDetails(html: BeautifulSoup) -> dict:
    result = html.find(
        "div", class_="ds-text-tight-m ds-font-regular ds-text-typo-mid3"
    ).text
    teams = html.find(
        "div", class_="ds-flex ds-flex-col ds-mt-3 md:ds-mt-0 ds-mt-0 ds-mb-1"
    )
    outcome = html.find(
        "p", class_="ds-text-tight-s ds-font-medium ds-truncate ds-text-typo"
    ).text

    matchDetails = dict()
    matchDetails.update(resultDetails(result))
    matchDetails.update(teamsPlayed(teams))
    matchDetails["outcome"] = outcome

    return matchDetails


def extractBatting(innings: BeautifulSoup) -> pd.DataFrame:
    batting = innings.find(
        "table",
        class_="ds-w-full ds-table ds-table-md ds-table-auto ci-scorecard-table",
    )

    battingRows = batting.find("tbody").find_all("tr")
    batRows = []

    for t in battingRows:
        tds = t.find_all("td")

        row = [td.text.strip() for td in tds]
        if len(row) > 8:
            row = row[:8]

        batRows.append(row)

        if "TOTAL" in t.text:
            break

    bat = pd.DataFrame(
        batRows, columns=["player", "dismissal_info", "r", "b", "m", "4s", "6s", "sr"]
    )
    bat = bat.dropna(subset=["r"])
    return bat


def extractBowling(innings: BeautifulSoup) -> pd.DataFrame:
    bowling = innings.find(
        "table", class_="ds-w-full ds-table ds-table-md ds-table-auto"
    )
    bowlingRows = bowling.find("tbody").find_all("tr")
    bowlRows = []

    for t in bowlingRows:
        tds = t.find_all("td")

        row = [td.text.strip() for td in tds]

        if len(row) > 11:
            row = row[:11]

        bowlRows.append(row)

    bowl = pd.DataFrame(
        bowlRows,
        columns=["player", "o", "m", "r", "w", "econ", "0s", "4s", "6s", "wd", "nb"],
    )
    bowl = bowl.dropna(subset=["r"])
    return bowl


def extractTotal(batting: pd.DataFrame) -> dict:
    total = batting.iloc[-1]
    oversPattern = re.compile(r"(.+?)Ov", re.I)
    res = dict()

    res["team"] = total["team"]
    for m in oversPattern.finditer(total["dismissal_info"]):
        res["overs"] = m.group(1).strip()

    scorePattern1 = re.compile(r"(\d+)\/(\d+)")
    scorePattern2 = re.compile(r"\d+")

    for m in scorePattern1.finditer(total["r"]):
        res["total"] = m.group(1).strip()
        res["wickets"] = m.group(2).strip()

    if "total" not in res:
        for m in scorePattern2.finditer(total["r"]):
            res["total"] = m.group(0).strip()
            res["wickets"] = "10"

    return res


def extractExtras(batting: pd.DataFrame) -> dict:

    extra = batting.iloc[-2]

    extras = dict()
    extras["team"] = extra["team"]
    byesPattern = re.compile(r"\(b\s*(\d+)", re.I)
    legByesPattern = re.compile(r"lb\s*(\d+)", re.I)
    widePattern = re.compile(r"w\s*(\d+)", re.I)
    noballPattern = re.compile(r"nb\s*(\d+)", re.I)

    for m in byesPattern.finditer(extra["dismissal_info"]):
        extras["b"] = m.group(1).strip()
        break

    if "b" not in extras:
        extras["b"] = "0"

    for m in legByesPattern.finditer(extra["dismissal_info"]):
        extras["lb"] = m.group(1).strip()

    if "lb" not in extras:
        extras["lb"] = "0"

    for m in widePattern.finditer(extra["dismissal_info"]):
        extras["w"] = m.group(1).strip()

    if "w" not in extras:
        extras["w"] = "0"

    for m in noballPattern.finditer(extra["dismissal_info"]):
        extras["nb"] = m.group(1).strip()

    if "nb" not in extras:
        extras["nb"] = "0"

    return extras


def extractInnings(bat1: pd.DataFrame, bat2: pd.DataFrame) -> pd.DataFrame:
    total = pd.DataFrame([extractTotal(bat1), extractTotal(bat2)])
    total.index = [1, 2]
    total = total.reset_index(names="innings")

    extras = pd.DataFrame([extractExtras(bat1), extractExtras(bat2)])
    extras.index = [1, 2]
    extras = extras.reset_index(names="innings")

    return total, extras


def extractScorecard(url: str) -> dict:
    matchID = extractMatchID(url)

    r = requests.get(url)
    html = BeautifulSoup(r.text, "lxml")

    matchDetails = getMatchDetails(html=html)
    matchDetails["matchID"] = matchID
    matchDetails = pd.DataFrame([matchDetails])
    matchDetails["date"] = pd.to_datetime(matchDetails["date"], format="%B %d, %Y")

    if "abandoned" in matchDetails["outcome"].iloc[0] or "No result" in matchDetails["outcome"].iloc[0]:
        return {
                "match_details": matchDetails,
                "batting": None,
                "bowling": None,
                "total": None,
                "extras": None,
            }
    tables = html.find_all("div", class_="ds-rounded-lg ds-mt-2")
    firstBatting = (
        tables[0]
        .find("span", class_="ds-text-title-xs ds-font-bold ds-capitalize")
        .text.strip()
    )

    secondBatting = (
        matchDetails["team_1"].iloc[0]
        if matchDetails["team_1"].iloc[0] != firstBatting
        else matchDetails["team_2"].iloc[0]
    )

    bat1, bat2 = extractBatting(tables[0]), extractBatting(tables[1])
    bat1["innings"], bat2["innings"] = 1, 2
    bat1["team"], bat2["team"] = firstBatting, secondBatting
    bat = pd.concat([bat1, bat2])
    bat["player"] = bat["player"].apply(removeNonASCII)
    bat["dismissal_info"] = bat["dismissal_info"].apply(removeNonASCII)
    bat["player"] = bat["player"].str.replace("(c)", "")
    bat["matchID"] = matchID

    bowl1, bowl2 = extractBowling(tables[0]), extractBowling(tables[1])
    bowl1["innings"], bowl2["innings"] = 1, 2
    bowl1["team"], bowl2["team"] = secondBatting, firstBatting
    bowl = pd.concat([bowl1, bowl2]).reset_index(drop=True)
    bowl["matchID"] = matchID

    total, extras = extractInnings(bat1, bat2)
    bat = bat[
        ~bat["player"].str.contains("TOTAL|Extras", case=False, na=False)
    ].reset_index(drop=True)
    total["matchID"] = matchID
    extras["matchID"] = matchID

    return {
        "match_details": matchDetails,
        "batting": bat,
        "bowling": bowl,
        "total": total,
        "extras": extras,
    }


def testScorecard(url:str) -> None:
    scorecard = extractScorecard(url)

    for s in scorecard:
        print(scorecard[s])
