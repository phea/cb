// Package cb is a Go client library for the CrunchBase API v2.
package cb

type ProductResponse struct {
	ProductData `json:"data"`
	Metadata    interface{} `json:"metadata"`
}

type ProductData struct {
	*Product      `json:"properties"`
	Relationships interface{} `json:"relationships"`
}

type Product struct {
	Name              string      `json:"name"`
	Permalink         string      `json:"permalink"`
	ShortDescription  string      `json:"short_description"`
	Description       string      `json:"description"`
	URL               string      `json:"homepage_url"`
	Launched          string      `json:"launched_on"`
	LaunchedDay       int64       `json:"launched_on_day"`
	LaunchedMonth     int64       `json:"launched_on_month"`
	LaunchedYear      int64       `json:"launched_on_year"`
	LaunchedTrustCode int64       `json:"launched_on_trust_code"`
	Lifecycle_stage   string      `json:"lifecycle_stage"`
	NumCustomersMax   int64       `json:"num_customers_max"`
	NumCustomersMin   int64       `json:"num_customers_min"`
	NumCustomersRange interface{} `json:"num_customers_range"`
	Created           int64       `json:"created_at"`
	Updated           int64       `json:"updated_at"`
}

func (c *Client) GetProduct(name string) (*Product, error) {
	var data *ProductResponse
	err := c.Call("/product/"+name, nil, &data)
	if err != nil {
		return nil, err
	}

	return data.Product, nil

}
