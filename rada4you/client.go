package rada4you

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	if fail := c.sendRequest("people", res); fail.IsOccur() {
		return nil, fail
	}
	return &GetAllPeoplesResponse{Peoples: *res}, nil
}

func (c *Client) GetPeopleByID(id int) (*GetPeopleByIdResponse, *ErrorResponse) {
	res := new(GetPeopleByIdResponse)
	url := fmt.Sprintf("people/%d", id)
	if fail := c.sendRequest(url, res); fail.IsOccur() {
		return nil, fail
	}
	return res, nil
}

func (c *Client) GetAllPolicies() (*GetAllPolicies, *ErrorResponse) {
	res := new([]Policy)
	if fail := c.sendRequest("policies", res); fail.IsOccur() {
		return nil, fail
	}
	return &GetAllPolicies{Policies: *res}, nil
}

func (c *Client) GetPolicyByID(id int) (*GetPolicyByIdResponse, *ErrorResponse) {
	res := new(GetPolicyByIdResponse)
	url := fmt.Sprintf("policies/%d", id)
	if fail := c.sendRequest(url, res); fail.IsOccur() {
		return nil, fail
	}
	return res, nil
}

func (c *Client) sendRequest(path string, target interface{}) *ErrorResponse {
	url := c.getRequestURL(path)
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

func (c *Client) getRequestURL(path string) string {
	return fmt.Sprintf("%s%s.json?key=%s", apiHost, path, c.APIKey)
}

func (c *Client) parseResponse(target *interface{}, res []byte) *ErrorResponse {
	fail := new(ErrorResponse)
	// Try to parse target response
	json.Unmarshal(res, target)
	// Try to parse error response
	json.Unmarshal(res, fail)
	return fail
}
