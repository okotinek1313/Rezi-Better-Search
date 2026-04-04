package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Search(query string) {
	url := "https://search.rezi.one/indexes/rezi/search"

	payload := strings.NewReader(fmt.Sprintf(`{"q":"%s","limit":20}`, query))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer e2a1974678b37386fef69bb3638a1fb36263b78a8be244c04795ada0fa250d3d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
