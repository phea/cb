// Package cb is a Go client library for the CrunchBase API v2.
package cb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// CrunchBase does not set response headers with the code so we have to unmarshal
// the json error message.
type ErrorResponse struct {
	Data struct {
		Error struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
		Response bool `json:"response"`
	} `json:"data"`
}

// apiUrl is the location to the CrunchBase API"
const baseURL = "http://api.crunchbase.com/v/2"

// Client struct
type Client struct {
	key        string
	httpClient *http.Client
}

// NewClient
func NewClient(k string) *Client {
	client = &Client{
		key:        k,
		httpClient: &http.Client{},
	}

	return client
}

func GetClient() *Client {
	return client
}

var client *Client

// Call
func (c *Client) Call(path string, params *url.Values, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if params == nil {
		path += "?user_key=" + c.key
	} else {
		params.Set("user_key", "6d68d412562ee027999dc31f66570bc2")
		path += "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", baseURL+path, nil)
	if err != nil {
		log.Printf("Cannot create an API request: %v", err)
		return err
	}

	req.Header.Set("User-Agent", "cb +http://github.com/phea/cb")
	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("Request to API failed: %v\n", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Cannot parse API response: %v\n", err)
		return err
	}

	// Check if CrunchBase API returns an error message
	var errResponse ErrorResponse
	json.Unmarshal(resBody, &errResponse)
	if errResponse.Data.Error.Code > 0 {
		return errors.New(errResponse.Data.Error.Message)
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}
