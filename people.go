// Package cb is a Go client library for the CrunchBase API v2.
package cb

import (
	"errors"
	"net/url"
)

type PeopleResponse struct {
	PeopleList `json:"data"`
	Metadata   interface{} `json:"metadata"`
}

type PeopleList struct {
	List []PeopleEntry `json:"items"`
	Page `json:"paging"`
}

type PeopleEntry struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	Created int64  `json:"created_at"`
	Updated int64  `json:"updated_at"`
}

func (c *Client) GetPeople(params *url.Values) (*PeopleList, error) {
	var data *PeopleResponse
	err := c.Call("/people", params, &data)
	if err != nil {
		return nil, err
	}

	return &data.PeopleList, nil
}

func (pl *PeopleList) Next() (*PeopleList, error) {
	if pl.Page.Current >= pl.Page.TotalPages {
		return nil, errors.New("Next page out of range")
	}

	u, err := url.Parse(pl.Page.NextURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.GetPeople(&params)
}

func (pl *PeopleList) Prev() (*PeopleList, error) {
	if pl.Page.Current <= 1 {
		return nil, errors.New("Previous page out of range")
	}

	u, err := url.Parse(pl.Page.PrevURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.GetPeople(&params)
}
