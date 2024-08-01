package batting

type BattingCard struct {
	PlayerID      int               `json:"id"`
	PlayerName    string            `json:"name"`
	CurrentTeam   string            `json:"currentTeam"`
	PreviousTeams []string          `json:"previousTeams"`
	Stats         CareerStats       `json:"stats"`
	Appearances   Appearances       `json:"appearances"`
	Performance   PerformanceInWins `json:"performanceInWins"`
	Last5Games    Last5Games        `json:"last5Games"`
}

type CareerStats struct {
	Season               int     `json:"season"`
	Matches              int     `json:"matches"`
	Innings              int     `json:"innings"`
	TotalRuns            int     `json:"totalRuns"`
	TotalBalls           int     `json:"totalBalls"`
	Centuries            int     `json:"centuries"`
	HalfCenturies        int     `json:"halfCenturies"`
	Fours                int     `json:"fours"`
	Sixes                int     `json:"sixes"`
	StrikeRate           float32 `json:"strikeRate"`
	BattingAverage       float32 `json:"battingAverage"`
	BoundariesPerInnings float32 `json:"boundariesPerInnings"`
	HighestScore         string  `json:"highestScore"`
	Ducks                int     `json:"ducks"`
}

type Appearances struct {
	SeasonsPlayed      int `json:"seasonsPlayed"`
	PlayoffAppearances int `json:"playoffAppearances"`
	FinalAppearances   int `json:"finalAppearances"`
	ChampionshipsWon   int `json:"championshipsWon"`
}

type PerformanceInWins struct {
	RunsInWins       int     `json:"runs"`
	AverageInWins    int     `json:"average"`
	StrikeRateInWins float32 `json:"strikeRate"`
}

type Last5Games struct {
	Runs           int     `json:"runs"`
	Balls          int     `json:"balls"`
	Boundaries     int     `json:"boundaries"`
	StrikeRate     float32 `json:"strikeRate"`
	BattingAverage float32 `json:"battingAverage"`
}
