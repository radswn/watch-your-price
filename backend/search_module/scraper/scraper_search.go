package scraper

// SearchResult represent the response from the website search
type SearchResult struct {
	Phrase     string            `json:"phrase" binding:"required"`
	Page       int               `json:"page" binding:"required"`
	NumOfPages int               `json:"numOfPages" binding:"required"`
	Results    map[string]string `json:"results" binding:"required"`
}

// SearchRequest represent query to the specific website search
type SearchRequest struct {
	Phrase  string      `json:"phrase" binding:"required"`
	Page    int         `json:"page"`
	Website WebsiteType `json:"website" binding:"required"`
}
