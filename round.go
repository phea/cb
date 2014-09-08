// Package cb is a Go client library for the CrunchBase API v2.
package cb

type RoundResponse struct {
	RoundData `json:"data"`
	Metadata  interface{} `json:"data"`
}

type RoundData struct {
	*Round        `json:"properties"`
	Relationships interface{} `json:"relationships"`
}

type Round struct {
	Name               string `json:"name"`
	Permalink          string `json:"permalink"`
	Announced          string `json:"announced_on"`
	AnnouncedDay       int64  `json:"announced_on_day"`
	AnnouncedMonth     int64  `json:"announced_on_month"`
	AnnouncedTrustCode int64  `json:"announced_on_trust_code"`
	AnnouncedYear      int64  `json:"announced_on_year"`
	Currency           string `json:"canonical_currency_code"`
	Type               string `json:"funding_type"`
	Amount             int64  `json:"money_raised"`
	CurrencyCode       string `json:"money_raised_currency_code"`
	AmountUSD          int64  `json:"money_raised_usd"`
	PostCurrency       string `json:"post_money_valuation_currency_code"`
	Created            int64  `json:"created_at"`
	Updated            int64  `json:"updated_at"`
}

func (c *Client) GetRound(id string) (*Round, error) {
	var data *RoundResponse
	err := c.Call("/funding-round/", nil, &data)
	if err != nil {
		return nil, err
	}

	return data.Round, nil
}
