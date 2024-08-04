package league

import (
	"database/sql"
	"ipl-api/database"
)

func GetBattingRecord(season int, db *sql.DB) Batting {
	var bat Batting

	bat.MostRuns = GetMostRuns(season, db)
	bat.MostFours = GetMostFours(season, db)
	bat.MostSixes = GetMostSixes(season, db)
	bat.Most50s = GetMost50s(season, db)
	bat.Most100s = GetMost100s(season, db)
	bat.BestStrikeRate = GetBestStrikeRate(season, db)
	bat.BestBattingAverages = GetBestBattingAverage(season, db)
	bat.Most4sPerInnings = GetMost4sPerInnings(season, db)
	bat.Most6sPerInnings = GetMost6sPerInnings(season, db)
	bat.Fastest50 = GetFastest50(season, db)
	bat.Fastest100 = GetFastest100(season, db)
	bat.HighestScore = GetHighestScore(season, db)

	return bat
}

func GetLeagueRecords(season int) League {
	var record League

	record.Season = season
	record.BatRecords = GetBattingRecord(season, database.DB)

	return record
}
