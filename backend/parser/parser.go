package parser

import "rezi-better-search/models"

type ParsedHits struct {
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	System   string `json:"system"`
	Source   string `json:"site"`
	Download string `json:"link"`
	IgdbUrl  string `json:"igdb_url"`
}

type ParsedResult struct {
	Hits         []ParsedHits `json:"hits"`
	TotalResults int          `json:"estimatedTotalHits"`
}

func Parse(result models.SearchResult) ParsedResult {
	var hits []ParsedHits

	for _, hit := range result.Hits {
		hits = append(hits, ParsedHits{
			Title:    hit.Title,
			Icon:     hit.Icon,
			System:   hit.System,
			Source:   hit.Site,
			Download: hit.Link,
			IgdbUrl:  hit.IgdbUrl,
		})
	}

	return ParsedResult{
		Hits:         hits,
		TotalResults: result.EstimatedTotalHits,
	}
}
