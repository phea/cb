// Package cb is a Go client library for the CrunchBase API v2.
package cb

type OrgResponse struct {
	OrgData  `json:"data"`
	Metadata interface{} `json:"metadata"`
}

type OrgData struct {
	*Organization `json:"properties"`
	Relationships interface{} `json:"relationships"`
}

type Organization struct {
	Name                string `json:"name"`
	Homepage            string `json:"homepage_url"`
	Permalink           string `json:"permalink"`
	Role                string `json:"primary_role"`
	IsClosed            bool   `json:"is_closed"`
	IsCompany           bool   `json:"role_company"`
	Description         string `json:"short_description"`
	NumberOfInvestments int64  `json:"number_of_investments"`
	TotalFunding        int64  `json:"total_funding_usd"`
	Created             int64  `json:"created_at"`
	Updated             int64  `json:"updated_at"`
}

func (c *Client) GetOrganization(name string) (*Organization, error) {
	var data *OrgResponse
	err := c.Call("/organization/"+name, nil, &data)
	if err != nil {
		return nil, err
	}

	return data.Organization, nil
}
