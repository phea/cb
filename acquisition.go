// Package cb is a Go client library for the CrunchBase API v2.
package cb

type AcquisitionResponse struct {
	AcquisitionData `json:"data"`
	Metadata        interface{} `json:"metadata"`
}

type AcquisitionData struct {
	*Acquisition  `json:"properties"`
	Relationships interface{} `json:"relationships"`
}

type Acquisition struct {
	Name                string `json:"name"`
	Permalink           string `json:"permalink"`
	Price               int64  `json:"price"`
	Price_currency_code string `json:"price_currency_code"`
	Type                string `json:"payment_type"`
	Announced           string `json:"announced_on"`
	AnnouncedDay        int64  `json:"announced_on_day"`
	AnnouncedMonth      int64  `json:"announced_on_month"`
	AnnouncedTrustCode  int64  `json:"announced_on_trust_code"`
	AnnouncedYear       int64  `json:"announced_on_year"`
	Created             int64  `json:"created_at"`
	Updated             int64  `json:"updated_at"`
}

func (c *Client) GetAcquisition(id string) (*Acquisition, error) {
	var data *AcquisitionResponse
	err := c.Call("/acquisition/"+id, nil, &data)
	if err != nil {
		return nil, err
	}

	return data.Acquisition, nil
}
