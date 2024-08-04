package league

type League struct {
	Season      int     `json:"season"`
	BatRecords  Batting `json:"battingRecords"`
	BowlRecords Bowling `json:"bowlingRecords"`
}

type Batting struct {
	MostRuns            BatOverall `json:"mostRuns"`
	MostFours           BatOverall `json:"mostFours"`
	MostSixes           BatOverall `json:"mostSixes"`
	Most50s             BatOverall `json:"most50s"`
	Most100s            BatOverall `json:"most100s"`
	BestStrikeRate      BatOverall `json:"bestSR"`
	BestBattingAverages BatOverall `json:"bestBattingAvg"`
	Most4sPerInnings    BatInnings `json:"most4sPerInnings"`
	Most6sPerInnings    BatInnings `json:"most6sPerInnings"`
	Fastest50           BatInnings `json:"fastest50"`
	Fastest100          BatInnings `json:"fastest100"`
	HighestScore        BatInnings `json:"highestScore"`
}

type BatOverall struct {
	Player        string  `json:"player"`
	Team          string  `json:"team"`
	Innings       int     `json:"innings"`
	Runs          int     `json:"runs"`
	Balls         int     `json:"balls"`
	Fours         int     `json:"fours"`
	Sixes         int     `json:"sixes"`
	Average       float32 `json:"battingAverage"`
	StrikeRate    float32 `json:"strikeRate"`
	HighestScore  string  `json:"highestScore"`
	HalfCenturies int     `json:"halfCenturies"`
	Centuries     int     `json:"centuries"`
}

type BatInnings struct {
	Player     string  `json:"player"`
	Team       string  `json:"team"`
	Against    string  `json:"against"`
	Year       int     `json:"season"`
	Runs       int     `json:"runs"`
	Balls      int     `json:"balls"`
	StrikeRate float32 `json:"strikeRate"`
	Fours      int     `json:"fours"`
	Sixes      int     `json:"sixes"`
}

type Bowling struct {
	MostWickets           BowlOverall `json:"mostWickets"`
	MostMaidens           BowlOverall `json:"mostMaidens"`
	MostDotBalls          BowlOverall `json:"mostDots"`
	BestBowlingAvg        BowlOverall `json:"bestBowlingAvg"`
	BestEconomy           BowlOverall `json:"bestEconomy"`
	BestSR                BowlOverall `json:"bestSR"`
	MostDotsPerInnings    BowlInnings `json:"mostDotsPerInnings"`
	BestEconomyPerInnings BowlInnings `json:"bestEconomyPerInnings"`
	BestSRPerInnings      BowlInnings `json:"bestSRPerInnings"`
	MostConcededRuns      BowlInnings `json:"mostConcededRuns"`
	BestBowlingFigures    BowlInnings `json:"bestBowlingFigures"`
}

type BowlOverall struct {
	Player         string  `json:"player"`
	Team           string  `json:"team"`
	Innings        int     `json:"innings"`
	Overs          float32 `json:"overs"`
	Runs           int     `json:"runs"`
	Wickets        int     `json:"wickets"`
	Dots           int     `json:"dots"`
	Maidens        int     `json:"maidens"`
	BestBowling    string  `json:"bestBowling"`
	Average        float32 `json:"bowlingAverage"`
	Economy        float32 `json:"economy"`
	StrikeRate     float32 `json:"strikeRate"`
	FourWicketHaul int     `json:"fourWicketHaul"`
	FiveWicketHaul int     `json:"fiveWicketHaul"`
}

type BowlInnings struct {
	Player     string  `json:"player"`
	Team       string  `json:"team"`
	Against    string  `json:"against"`
	Overs      float32 `json:"overs"`
	Runs       int     `json:"runs"`
	Wickets    int     `json:"wickets"`
	Dots       int     `json:"dots"`
	Economy    float32 `json:"economy"`
	StrikeRate float32 `json:"strikeRate"`
}
