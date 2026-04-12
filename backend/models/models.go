package models

type Hit struct {
	Id       string      `json:"id"`
	Bios     interface{} `json:"bios"`
	Core     interface{} `json:"core"`
	Icon     string      `json:"icon"`
	Link     string      `json:"link"`
	Playable string      `json:"playable"`
	System   string      `json:"system"`
	Title    string      `json:"title"`
	IgdbUrl  string      `json:"igdb_url"`
	Site     string      `json:"site"`
}

type SearchResult struct {
	Hits               []Hit  `json:"hits"`
	Query              string `json:"query"`
	ProcessingTimeMs   int    `json:"processingTimeMs"`
	Limit              int    `json:"limit"`
	Offset             int    `json:"offset"`
	EstimatedTotalHits int    `json:"estimatedTotalHits"`
}
