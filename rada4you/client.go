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

// New function for create a new instance of Rada4You client
func New(key string) Client {
	return Client{key}
}

// GetAllPeoples function for retrieve the list of current deputies
func (c *Client) GetAllPeoples() (*[]GetAllPeoplesResponse, *ErrorResponse) {
	res := new([]GetAllPeoplesResponse)
	if fail := c.sendRequest("people", res); fail != nil {
		return &[]GetAllPeoplesResponse{}, fail
	}
	return res, nil
}

func (c *Client) GetPeopleById(id int) (*GetPeopleByIdResponse, *ErrorResponse) {
	res := new(GetPeopleByIdResponse)
	url := fmt.Sprintf("people/%d", id)
	if fail := c.sendRequest(url, res); fail != nil {
		return &GetPeopleByIdResponse{}, fail
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
	// Try to parse target response
	err := json.Unmarshal(res, target)
	if err == nil {
		return nil
	}

	// Try to parse error response
	fail := ErrorResponse{}
	err = json.Unmarshal(res, &fail)
	if err != nil {
		// Cannot parse API response
		panic(err)
	}

	return &fail
}
