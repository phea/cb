// Package cb is a Go client library for the CrunchBase API v2.
package cb

type FundRaiseResponse struct {
	FundRaiseData `json:"data"`
	Metadata      interface{} `json:"metadata"`
}

type FundRaiseData struct {
	*FundRaise    `json:"properties"`
	Relationships interface{} `json:"relationships"`
}

type FundRaise struct {
	Name             string `json:"name"`
	Permalink        string `json:"permalink"`
	MoneyRaised      int64  `json:"money_raised"`
	CurrencyCode     string `json:"money_raised_currency_code"`
	Started          string `json:"started_on"`
	StartedDay       int64  `json:"started_on_day"`
	StartedMonth     int64  `json:"started_on_month"`
	StartedTrustCode int64  `json:"started_on_trust_code"`
	StartedYear      int64  `json:"started_on_year"`
	Created          int64  `json:"created_at"`
	Updated          int64  `json:"updated_at"`
}

func (c *Client) GetFundRaise(id string) (*FundRaise, error) {
	var data *FundRaiseResponse
	err := c.Call("/fund-raise/"+id, nil, &data)
	if err != nil {
		return nil, err
	}

	return data.FundRaise, nil
}
