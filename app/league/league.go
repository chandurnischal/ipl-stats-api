package league

import (
	"ipl-api/database"
)

func GetBattingRecord(season int) Batting {
	var bat Batting

	db := database.DB

	bat.Season = season
	bat.MostRuns = GetMostRuns(season, db)
	bat.MostFours = GetMostFours(season, db)
	bat.MostSixes = GetMostSixes(season, db)
	bat.Most50s = GetMostFifties(season, db)
	bat.Most100s = GetMostHundrends(season, db)
	bat.BestStrikeRate = GetBestStrikeRate(season, db)
	bat.BestBattingAverages = GetBestBattingAverage(season, db)
	bat.Most4sPerInnings = GetMost4sPerInnings(season, db)
	bat.Most6sPerInnings = GetMost6sPerInnings(season, db)
	bat.Fastest50 = GetFastest50(season, db)
	bat.Fastest100 = GetFastest100(season, db)
	bat.HighestScore = GetHighestScore(season, db)

	return bat
}

func GetBowlingRecord(season int) Bowling {
	var bowl Bowling

	db := database.DB

	bowl.Season = season
	bowl.MostWickets = GetMostWickets(season, db)
	bowl.MostMaidens = GetMostMaidens(season, db)
	bowl.MostDotBalls = GetMostDots(season, db)
	bowl.BestBowlingAvg = GetBestBowlingAverage(season, db)
	bowl.BestEconomy = GetBestEconomy(season, db)
	bowl.BestSR = GetBestBowlStrikeRate(season, db)
	bowl.MostDotsPerInnings = GetMostDotsPerInnings(season, db)
	bowl.BestEconomyPerInnings = GetBestEconomyPerInnings(season, db)
	bowl.BestSRPerInnings = GetBestBowlStrikeRatePerInnings(season, db)
	bowl.MostConcededRuns = GetMostConcededRuns(season, db)
	bowl.BestBowlingFigures = GetBestBowlingFigures(season, db)

	return bowl
}
