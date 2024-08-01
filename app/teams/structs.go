package teams

type Team struct {
	TeamID                int         `json:"id"`
	TeamName              string      `json:"team"`
	Season                int         `json:"season"`
	Matches               Matches     `json:"matches"`
	Appearances           Appearances `json:"playoffs"`
	Stats                 Stats       `json:"stats"`
	IndiviudalPerformance Indiviudal  `json:"individualPerformance"`
}

type Matches struct {
	Played            int     `json:"played"`
	Won               int     `json:"won"`
	Lost              int     `json:"lost"`
	Tied              int     `json:"tied"`
	NoResult          int     `json:"noResult"`
	BattingFirstPerc  float32 `json:"firstBatWinPerc"`
	FieldingFirstPerc float32 `json:"firstFieldWinPerc"`
	WinPercentage     float32 `json:"winPerc"`
}

type Appearances struct {
	Played        int `json:"seasonPlayed"`
	Appearances   int `json:"playoffAppearances"`
	Finals        int `json:"finalAppearances"`
	Championships int `json:"championships"`
}

type Stats struct {
	HighestScore   string  `json:"highestScore"`
	LowestScore    string  `json:"lowestScore"`
	AverageRuns    float32 `json:"averageRuns"`
	AverageWickets float32 `json:"averageWickets"`
	TotalRuns      int     `json:"totalRuns"`
	TotalWickets   int     `json:"totalWickets"`
}

type TopRunScorer struct {
	Name string `json:"name"`
	Runs int    `json:"runs"`
}

type TopWicketTaker struct {
	Name    string `json:"name"`
	Wickets int    `json:"wickets"`
}

type Indiviudal struct {
	BestBatting    string         `json:"bestBatting"`
	BestBowling    string         `json:"bestBowling"`
	TopRunScorer   TopRunScorer   `json:"topRunScorer"`
	TopWicketTaker TopWicketTaker `json:"topWicketTaker"`
}
