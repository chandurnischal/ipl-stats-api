package teams

type Team struct {
	TeamID               int                `json:"id"`
	TeamName             string             `json:"name"`
	Season               int                `json:"season"`
	TeamHistory          History            `json:"history"`
	TeamMatchRecord      MatchRecord        `json:"matchRecord"`
	TeamPerformanceStats PerformanceStats   `json:"performanceStats"`
	PlayerAchievements   PlayerAchievements `json:"playerAchievements"`
}

type History struct {
	Seasons       int `json:"seasons"`
	Playoffs      int `json:"playoffs"`
	Finals        int `json:"finals"`
	Championships int `json:"championships"`
}

type MatchRecord struct {
	Played   int `json:"played"`
	Won      int `json:"won"`
	Lost     int `json:"lost"`
	NoResult int `json:"noresult"`
	Tied     int `json:"tied"`
}

type PerformanceStats struct {
	TotalRuns      int           `json:"totalRuns"`
	TotalWickets   int           `json:"totalWickets"`
	AverageRuns    int           `json:"avgRuns"`
	AverageWickets int           `json:"avgWickets"`
	HighestScore   string        `json:"highestScore"`
	BestBowling    string        `json:"bestBowling"`
	WinPerc        WinPercentage `json:"winPercentage"`
}

type WinPercentage struct {
	Overall       float32 `json:"overall"`
	BatFirst      float32 `json:"batFirst"`
	FieldFirst    float32 `json:"fieldFirst"`
	Last5Matches  string  `json:"lastFiveMatches"`
	Last10Matches string  `json:"lastTenMatches"`
}

type PlayerAchievements struct {
	TopScorer struct {
		Player  string `json:"player"`
		Score   string `json:"score"`
		Against string `json:"against"`
		Year    int    `json:"year"`
	} `json:"topScorerInAMatch"`
	TopBowler struct {
		Player  string `json:"player"`
		Figure  string `json:"bowlingFigure"`
		Against string `json:"against"`
		Year    int    `json:"year"`
	} `json:"topBowlerInAMatch"`
	MostCenturies struct {
		Player    string `json:"player"`
		Centuries int    `json:"centuries"`
	} `json:"mostCenturies"`
	MostWickets struct {
		Player  string `json:"player"`
		Wickets int    `json:"wickets"`
	} `json:"mostWickets"`
}
