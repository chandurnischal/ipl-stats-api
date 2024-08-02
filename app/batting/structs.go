package batting

type BattingCard struct {
	PlayerID          int      `json:"id"`
	PlayerName        string   `json:"player"`
	Team              string   `json:"team"`
	OtherTeams        []string `json:"otherTeams"`
	TournamentHistory History  `json:"tournamentHistory"`
	Stats             Stats    `json:"stats"`
}

type History struct {
	Matches       int `json:"matches"`
	Seasons       int `json:"season"`
	Playoffs      int `json:"playoffs"`
	Finals        int `json:"finals"`
	Championships int `json:"championships"`
}

type Stats struct {
	Innings        int     `json:"innings"`
	TotalRuns      int     `json:"totalRuns"`
	TotalBalls     int     `json:"totalBalls"`
	Centuries      int     `json:"centuries"`
	HalfCenturies  int     `json:"halfCenturies"`
	Fours          int     `json:"fours"`
	Sixes          int     `json:"sixes"`
	StrikeRate     float32 `json:"strikeRate"`
	BattingAverage float32 `json:"battingAverage"`
	HighestScore   struct {
		Runs    int    `json:"runs"`
		Balls   int    `json:"balls"`
		Team    string `json:"team"`
		Against string `json:"against"`
		Year    int    `json:"int"`
	} `json:"highestScore"`
	Ducks        int `json:"ducks"`
	Last5Matches struct {
		Runs           int     `json:"runs"`
		Boundaries     int     `json:"boundaries"`
		StrikeRate     float32 `json:"strikeRate"`
		BattingAverage float32 `json:"battingAverage"`
	} `json:"last5Matches"`
	Last10Matches struct {
		Runs           int     `json:"runs"`
		Boundaries     int     `json:"boundaries"`
		StrikeRate     float32 `json:"strikeRate"`
		BattingAverage float32 `json:"battingAverage"`
	} `json:"last10Matches"`
}
