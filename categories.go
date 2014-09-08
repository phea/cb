// Package cb is a Go client library for the CrunchBase API v2.
package cb

import (
	"errors"
	"net/url"
)

type CategoriesResponse struct {
	*CategoriesList `json:"data"`
	Metadata        interface{} `json:"metadata"`
}

type CategoriesList struct {
	List []*CategoryEntry `json:"items"`
	Page `json:"paging"`
}

type CategoryEntry struct {
	ID                    string `json:"uuid"`
	Name                  string `json:"name"`
	Path                  string `json:"path"`
	Type                  string `json:"type"`
	NumberOfOrganizations int64  `json:"number_of_organizations"`
	Created               int64  `json:"created_at"`
	Updated               int64  `json:"updated_at"`
}

func (c *Client) GetCategories(params *url.Values) (*CategoriesList, error) {
	var data *CategoriesResponse
	err := c.Call("/categories", params, &data)
	if err != nil {
		return nil, err
	}

	return data.CategoriesList, nil
}

func (cl *CategoriesList) Next() (*CategoriesList, error) {
	if cl.Page.Current >= cl.Page.TotalPages {
		return nil, errors.New("Next page out of range")
	}

	u, err := url.Parse(cl.Page.NextURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.GetCategories(&params)
}

func (cl *CategoriesList) Prev() (*CategoriesList, error) {
	if cl.Page.Current <= 1 {
		return nil, errors.New("Previous page out of range")
	}

	u, err := url.Parse(cl.Page.PrevURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.GetCategories(&params)
}
