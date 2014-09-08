// Package cb is a Go client library for the CrunchBase API v2.
package cb

type IPOResponse struct {
	IPOData  `json:"data"`
	Metadata interface{} `json:"metadata"`
}

type IPOData struct {
	*IPO          `json:"properties"`
	Relationships interface{} `json:"relationships"`
}

type IPO struct {
	Name                         string      `json:"name"`
	Permalink                    string      `json:"permalink"`
	StockSymbol                  string      `json:"stock_symbol"`
	StockExchangeSymbol          interface{} `json:"stock_exchange_symbol"`
	CurrencyCode                 string      `json:"canonical_currency_code"`
	MoneyRaisedCurrencyCode      string      `json:"money_raised_currency_code"`
	MoneyRaisedUSD               interface{} `json:"money_raised_usd"`
	OpeningPriceCurrencyCode     string      `json:"opening_share_price_currency_code"`
	OpeningSharePriceUSD         interface{} `json:"opening_share_price_usd"`
	OpeningValuationCurrencyCode string      `json:"opening_valuation_currency_code"`
	OpeningValuationUSD          interface{} `json:"opening_valuation_usd"`
	WentPublic                   string      `json:"went_public_on"`
	WentPublicDay                int64       `json:"went_public_on_day"`
	WentPublicMonth              int64       `json:"went_public_on_month"`
	WentPublicTrustCode          int64       `json:"went_public_on_trust_code"`
	WentPublicYear               int64       `json:"went_public_on_year"`
	Updated                      int64       `json:"updated_at"`
	Created_at                   int64       `json:"created_at"`
}

func (c *Client) GetIPO(id string) (*IPO, error) {
	var data *IPOResponse
	err := c.Call("/ipo/"+id, nil, &data)
	if err != nil {
		return nil, err
	}

	return data.IPO, nil
}
