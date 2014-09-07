// Package cb is a Go client library for the CrunchBase API v2.
package cb

import "encoding/json"

// Person Response struct is the general struct to hold the response
type PersonResponse struct {
	PersonData `json:"data"`
	Metadata   interface{} `json:"metadata"`
}

type PersonData struct {
	Properties    *Person                `json:"properties"`
	Relationships map[string]interface{} `json:"relationships"`
}

// Person struct holds the info in the Properties tag.
type Person struct {
	Id         string `json:"uuid"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Permalink  string `json:"permalink"`
	Born       string `json:"born_on"`
	Bio        string `json:"bio"`
	IsInvestor bool   `json:"role_investor"`
	Created    int64  `json:"created_at"`
	Updated    int64  `json:"updated_at"`
	LocationId string `json:"location_uuid"`
}

// GetPerson takes the person's permalink as id and returns a Person object.
// Note: 404 will return an error
func (c *Client) GetPerson(id string) (*Person, error) {
	var data *PersonResponse
	err := c.Call("/person/"+id, nil, &data)
	if err != nil {
		return nil, err
	}

	person := data.Properties
	return person, nil
}

// GetDegree returns a list of the person's degree.
func (p *Person) GetDegree() ([]*Degree, error) {
	c := GetClient()
	degrees, err := c.GetDegree(p.Permalink)
	if err != nil {
		return nil, err
	}

	return degrees, nil
}

func (p *Person) UnmarshalJSON(data []byte) error {
	type person Person
	var pp person
	err := json.Unmarshal(data, &pp)
	if err != nil {
		return err
	}

	*p = Person(pp)
	return nil
}
