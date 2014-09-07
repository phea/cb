// Package cb is a Go client library for the CrunchBase API v2.
package cb

type DegreeResponse struct {
	DegreeData `json:"data"`
	Metadata   interface{} `json:"metadata"`
}

type DegreeData struct {
	Items  []*Degree   `json:"items"`
	Paging interface{} `json:"paging"`
}

// Degree struct
type Degree struct {
	Started          string `json:"started_on"`
	Completed        string `json:"completed_on"`
	Type             string `json:"type"`
	DegreeType       string `json:"degree_type_name"`
	Subject          string `json:"degree_subject"`
	Organization     string `json:"organization_name"`
	OrganizationPath string `json:"organization_path"`
}

// GetDegree takes a person's permalink and returns their education experience.
func (c *Client) GetDegree(name string) ([]*Degree, error) {
	var data *DegreeResponse
	err := c.Call("/person/"+name+"/degrees", nil, &data)
	if err != nil {
		return nil, err
	}

	degreeList := data.Items
	return degreeList, nil
}
