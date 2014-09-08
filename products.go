// Package cb is a Go client library for the CrunchBase API v2.
package cb

import (
	"errors"
	"net/url"
)

type ProductsResponse struct {
	ProductsList `json:"data"`
	Metadata     interface{} `json:"metadata"`
}

type ProductsList struct {
	List []*ProductEntry `json:"items"`
	Page `json:"paging"`
}

type ProductEntry struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Type    string `json:"product"`
	Created int64  `json:"created_at"`
	Updated int64  `json:"updated_at"`
}

func (c *Client) GetProducts(params *url.Values) (*ProductsList, error) {
	var data *ProductsResponse
	err := c.Call("/products", params, &data)
	if err != nil {
		return nil, err
	}

	return &data.ProductsList, nil
}

func (pl *ProductsList) Next() (*ProductsList, error) {
	if pl.Page.Current >= pl.Page.TotalPages {
		return nil, errors.New("Next page out of range")
	}

	u, err := url.Parse(pl.Page.NextURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.GetProducts(&params)
}

func (pl *ProductsList) Prev() (*ProductsList, error) {
	if pl.Page.Current <= 1 {
		return nil, errors.New("Previous page out of range")
	}

	u, err := url.Parse(pl.Page.PrevURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.GetProducts(&params)
}
