package rada4you

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const apiHost string = "https://rada4you.org/api/v1/"

// Client struct that contain all API methods
type Client struct {
	APIKey string
}

// New function for create a new instance of Rada4You client
func New(key string) Client {
	return Client{key}
}

func (c *Client) sendRequest(path string, target interface{}) error {
	url := c.getRequestURL(path)
	res, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "API request sending failed")
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func (c *Client) getRequestURL(path string) string {
	return fmt.Sprintf("%s%s.json?key=%s", apiHost, path, c.APIKey)
}
