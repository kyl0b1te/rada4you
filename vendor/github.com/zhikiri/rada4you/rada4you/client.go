package rada4you

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	nurl "net/url"
)

const apiHost = "https://rada4you.org/api/v1/"

// Client struct that contain all API methods
type Client struct {
	APIKey string
}

func New(key string) Client {
	return Client{key}
}

func (c *Client) GetAllPeoples() (*GetAllPeoplesResponse, *ErrorResponse) {
	res := new([]Person)
	if fail := c.sendRequest("people", res, make(map[string]string)); fail.IsOccur() {
		return nil, fail
	}
	return &GetAllPeoplesResponse{Peoples: *res}, nil
}

func (c *Client) GetPeopleByID(id int) (*GetPeopleByIDResponse, *ErrorResponse) {
	res := new(GetPeopleByIDResponse)
	url := fmt.Sprintf("people/%d", id)
	if fail := c.sendRequest(url, res, make(map[string]string)); fail.IsOccur() {
		return nil, fail
	}
	return res, nil
}

func (c *Client) GetAllPolicies() (*GetAllPoliciesResponse, *ErrorResponse) {
	res := new([]Policy)
	if fail := c.sendRequest("policies", res, make(map[string]string)); fail.IsOccur() {
		return nil, fail
	}
	return &GetAllPoliciesResponse{Policies: *res}, nil
}

func (c *Client) GetPolicyByID(id int) (*GetPolicyByIDResponse, *ErrorResponse) {
	res := new(GetPolicyByIDResponse)
	url := fmt.Sprintf("policies/%d", id)
	if fail := c.sendRequest(url, res, make(map[string]string)); fail.IsOccur() {
		return nil, fail
	}
	return res, nil
}

func (c *Client) GetAllDivisions(r GetAllDivisionsRequest) (*GetAllDivisionsResponse, *ErrorResponse) {
	res := new([]Division)
	if fail := c.sendRequest("divisions", res, r.Values()); fail.IsOccur() {
		return nil, fail
	}
	return &GetAllDivisionsResponse{Divisions: *res}, nil
}

func (c *Client) GetDivisionByID(id int) (*GetDivisionByIDResponse, *ErrorResponse) {
	res := new(GetDivisionByIDResponse)
	url := fmt.Sprintf("divisions/%d", id)
	if fail := c.sendRequest(url, res, make(map[string]string)); fail.IsOccur() {
		return nil, fail
	}
	return res, nil
}

func (c *Client) getQueryString(params map[string]string) string {
	values := nurl.Values{}
	for key, val := range params {
		values.Add(key, val)
	}
	return values.Encode()
}

func (c *Client) sendRequest(path string, target interface{}, params map[string]string) *ErrorResponse {
	url := c.getRequestURL(path, params)
	res, err := http.Get(url)
	if err != nil {
		panic("API request sending failed")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Response body reading failed")
	}

	if fail := c.parseResponse(&target, body); fail != nil {
		return fail
	}

	return nil
}

func (c *Client) getRequestURL(path string, params map[string]string) string {
	params["key"] = c.APIKey
	return fmt.Sprintf("%s%s.json?%s", apiHost, path, c.getQueryString(params))
}

func (c *Client) parseResponse(target *interface{}, res []byte) *ErrorResponse {
	fail := new(ErrorResponse)
	// Try to parse target response
	json.Unmarshal(res, target)
	// Try to parse error response
	json.Unmarshal(res, fail)
	return fail
}
