// Package cb is a Go client library for the CrunchBase API v2.
package cb

type Page struct {
	TotalItems   int64  `json:"total_items"`
	Current      int64  `json:"current_page"`
	TotalPages   int64  `json:"number_of_pages"`
	ItemsPerPage int64  `json:"items_per_page"`
	SortOrder    string `json:"sort_order"`
	NextURL      string `json:"next_page_url"`
	PrevURL      string `json:"prev_page_url"`
}
