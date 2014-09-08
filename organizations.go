// Package cb is a Go client library for the CrunchBase API v2.
package cb

import (
	"errors"
	"net/url"
)

type OrgListResponse struct {
	*OrgList `json:"data"`
	Metadata interface{} `json:"metadata"`
}

type OrgList struct {
	List []*OrgEntry `json:"items"`
	Page `json:"paging"`
}

type OrgEntry struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Type    string `json"type"`
	Created int64  `json:"created_at"`
	Updated int64  `json:"updated_at"`
}

func (c *Client) OrgSearch(params *url.Values) (*OrgList, error) {
	var data *OrgListResponse
	err := c.Call("/organizations", params, &data)
	if err != nil {
		return nil, err
	}

	return data.OrgList, nil
}

func (ol *OrgList) Next() (*OrgList, error) {
	if ol.Page.Current >= ol.Page.TotalPages {
		return nil, errors.New("Next page out of range")
	}

	u, err := url.Parse(ol.Page.NextURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.OrgSearch(&params)
}

func (ol *OrgList) Prev() (*OrgList, error) {
	if ol.Page.Current <= 1 {
		return nil, errors.New("Previous page out of range")
	}

	u, err := url.Parse(ol.Page.PrevURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.OrgSearch(&params)
}
