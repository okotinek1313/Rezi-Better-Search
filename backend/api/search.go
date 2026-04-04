package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

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

func Search(query string) (SearchResult, error) {
	url := "https://search.rezi.one/indexes/rezi/search"

	payload := strings.NewReader(fmt.Sprintf(`{"q":"%s","limit":20}`, query))

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer e2a1974678b37386fef69bb3638a1fb36263b78a8be244c04795ada0fa250d3d")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result SearchResult
	err := json.Unmarshal(body, &result)

	return result, err
}
