// Package cb is a Go client library for the CrunchBase API v2.
package cb

import (
	"errors"
	"net/url"
)

type LocationsResponse struct {
	*LocationsList `json:"data"`
	Metadata       interface{} `json:"metadata"`
}

type LocationsList struct {
	List []*LocationEntry `json:"items"`
	Page `json:"paging"`
}

type LocationEntry struct {
	ID               string `json:"uuid"`
	Name             string `json:"name"`
	Path             string `json:"path"`
	Type             string `json:"type"`
	LocationType     string `json:"location_type"`
	ParentLocationID string `json:"parent_location_uuid"`
	Created          int64  `json:"created_at"`
	Updated          int64  `json:"updated_at"`
}

func (c *Client) GetLocations(params *url.Values) (*LocationsList, error) {
	var data *LocationsResponse
	err := c.Call("/locations", params, &data)
	if err != nil {
		return nil, err
	}

	return data.LocationsList, nil
}

func (ll *LocationsList) Next() (*LocationsList, error) {
	if ll.Page.Current >= ll.Page.TotalPages {
		return nil, errors.New("Next page out of range")
	}

	u, err := url.Parse(ll.Page.NextURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.GetLocations(&params)
}

func (ll *LocationsList) Prev() (*LocationsList, error) {
	if ll.Page.Current <= 1 {
		return nil, errors.New("Previous page out of range")
	}

	u, err := url.Parse(ll.Page.PrevURL)
	if err != nil {
		return nil, err
	}

	c := GetClient()
	params := u.Query()
	return c.GetLocations(&params)
}
