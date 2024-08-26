package validation

import (
	"net/http"
	"strconv"
)

type SearchQuery struct {
	SearchString string
	Page         int
}

func ValidateSearchQuery(r *http.Request) (SearchQuery, error) {
	searchString := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return SearchQuery{}, err
	}
	return SearchQuery{SearchString: searchString, Page: pageInt}, nil
}
