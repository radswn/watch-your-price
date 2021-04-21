package scraper

// CheckResult represent the response from the website price check
type CheckResult struct {
	Price string `json:"price" binding:"required"`
}

// CheckRequest represent query to the specific website price check
type CheckRequest struct {
	Url     string      `json:"phrase" binding:"required"`
	Website WebsiteType `json:"website" binding:"required"`
}
