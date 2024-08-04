package players

type PlayerCard struct {
	PlayerID      int         `json:"id"`
	PlayerName    string      `json:"player"`
	Season        int         `json:"season"`
	Team          string      `json:"team"`
	OtherTeams    []string    `json:"otherTeams"`
	CareerHistory History     `json:"tournamentHistory"`
	Batting       BattingCard `json:"battingCard,omitempty"`
	Bowling       BowlingCard `json:"bowlingCard,omitempty"`
}

type History struct {
	Matches       int `json:"matches"`
	Seasons       int `json:"season"`
	Playoffs      int `json:"playoffs"`
	Finals        int `json:"finals"`
	Championships int `json:"championships"`
}

type BattingCard struct {
	CareerStats      BatStats            `json:"careerStats,omitempty"`
	PerformanceStats BatPerformanceStats `json:"performanceStats,omitempty"`
}

type BatStats struct {
	Innings        int     `json:"innings"`
	TotalRuns      int     `json:"totalRuns"`
	TotalBalls     int     `json:"totalBalls"`
	Centuries      int     `json:"centuries"`
	HalfCenturies  int     `json:"halfCenturies"`
	Fours          int     `json:"fours"`
	Sixes          int     `json:"sixes"`
	StrikeRate     float32 `json:"strikeRate"`
	BattingAverage float32 `json:"battingAverage"`
	Ducks          int     `json:"ducks"`
}

type BatPerformanceStats struct {
	HighestScore struct {
		Runs    int    `json:"runs"`
		Balls   int    `json:"balls"`
		Team    string `json:"team"`
		Against string `json:"against"`
		Year    int    `json:"season"`
	} `json:"highestScore"`
	Last5Innings struct {
		Runs           int     `json:"runs"`
		Boundaries     int     `json:"boundaries"`
		StrikeRate     float32 `json:"strikeRate"`
		BattingAverage float32 `json:"battingAverage"`
	} `json:"Last5Innings"`
	Last10Innings struct {
		Runs           int     `json:"runs"`
		Boundaries     int     `json:"boundaries"`
		StrikeRate     float32 `json:"strikeRate"`
		BattingAverage float32 `json:"battingAverage"`
	} `json:"Last10Innings"`
}

type BowlingCard struct {
	CareerStats      BowlStats            `json:"careerStats"`
	PerformanceStats BowlPerformanceStats `json:"performanceStats"`
}

type BowlStats struct {
	Innings         int     `json:"innings"`
	TotalRuns       int     `json:"totalRuns"`
	TotalWickets    int     `json:"totalWickets"`
	Dots            int     `json:"dots"`
	BowlingAverage  float32 `json:"bowlingAverage"`
	Economy         float32 `json:"economy"`
	StrikeRate      float32 `json:"strikeRate"`
	FourWicketHauls int     `json:"FourWicketHauls"`
	FiveWicketHauls int     `json:"FiveWicketHauls"`
	Maidens         int     `json:"maidens"`
}

type BowlPerformanceStats struct {
	BestBowling struct {
		Wickets int    `json:"wickets"`
		Runs    int    `json:"runs"`
		Team    string `json:"team"`
		Against string `json:"against"`
		Year    int    `json:"season"`
	} `json:"bestBowling"`
	Last5Innings struct {
		Wickets        int     `json:"wickets"`
		Runs           int     `json:"runs"`
		Economy        float32 `json:"economy"`
		BowlingAverage float32 `json:"bowlingAverage"`
	} `json:"Last5Innings"`
	Last10Innings struct {
		Wickets        int     `json:"wickets"`
		Runs           int     `json:"runs"`
		Economy        float32 `json:"economy"`
		BowlingAverage float32 `json:"bowlingAverage"`
	} `json:"Last10Innings"`
}
