package rada4you

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
)

var APIKey string
var CLI Client

func init() {
	APIKey = os.Getenv("API_KEY")
	CLI = New(APIKey)
}

func TestNew(t *testing.T) {
	assert.Equal(t, APIKey, New(APIKey).APIKey)
}

func TestGetRequestURL(t *testing.T) {
	// Test without query string parameters
	generated := CLI.getRequestURL("entity", make(map[string]string))
	calculated := fmt.Sprintf("%sentity.json?key=%s", apiHost, CLI.APIKey)
	assert.Equal(t, calculated, generated)

	// Test with query string parameters
	queryParameters := make(map[string]string)
	queryParameters["a"] = "1"
	queryParameters["b"] = "2"
	queryParameters["c"] = "3"

	queryString := url.Values{}
	queryString.Add("key", CLI.APIKey)
	for k, v := range queryParameters {
		queryString.Add(k, v)
	}
	generated = CLI.getRequestURL("entity", queryParameters)
	calculated = fmt.Sprintf("%sentity.json?%s", apiHost, queryString.Encode())
	assert.Equal(t, calculated, generated)
}

func TestGetQueryString(t *testing.T) {
	queryParameters := make(map[string]string)
	queryParameters["a"] = "1"
	queryParameters["b"] = "2"
	queryParameters["c"] = "3"

	queryString := url.Values{}
	for k, v := range queryParameters {
		queryString.Add(k, v)
	}
	assert.Equal(t, queryString.Encode(), CLI.getQueryString(queryParameters))
}

func TestInvalidApiKey(t *testing.T) {
	cli := New("secret")
	msg := "You need a valid api key. Sign up for an account on Вони голосують для тебе to get one."
	if _, err := cli.GetAllPeoples(); err != nil {
		assert.Equal(t, msg, err.Message)
	}
}

func TestClient_GetAllPeoples(t *testing.T) {
	mock := new([]Person)
	setResponseMock(CLI.getRequestURL("people", make(map[string]string)), "GetAllPeoplesResponseMock", mock)
	defer unsetResponseMock()

	// Try get all peoples
	res, err := CLI.GetAllPeoples()
	assert.Nil(t, err)
	assert.NotNil(t, res.Peoples)

	assert.Exactly(t, &GetAllPeoplesResponse{Peoples: *mock}, res)
}

func TestClient_GetAllPolicies(t *testing.T) {
	mock := new([]Policy)
	setResponseMock(CLI.getRequestURL("policies", make(map[string]string)), "GetAllPoliciesResponseMock", mock)
	defer unsetResponseMock()

	res, err := CLI.GetAllPolicies()
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Exactly(t, &GetAllPoliciesResponse{Policies: *mock}, res)
}

func TestClient_GetAllDivisions(t *testing.T) {
	mock := new([]Division)
	setResponseMock(CLI.getRequestURL("divisions", make(map[string]string)), "GetAllDivisionsResponseMock", mock)
	defer unsetResponseMock()

	res, err := CLI.GetAllDivisions(GetAllDivisionsRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Exactly(t, &GetAllDivisionsResponse{Divisions: *mock}, res)
}

func TestClient_GetPeopleByID(t *testing.T) {
	mock := new(GetPeopleByIDResponse)
	setResponseMock(CLI.getRequestURL("people/0", make(map[string]string)), "ErrorResponseMock", mock)
	defer unsetResponseMock()

	// Try invalid ID
	_, err := CLI.GetPeopleByID(0)
	assert.NotNil(t, err)
	assert.Equal(t, "Not Found", err.Message)

	setResponseMock(CLI.getRequestURL("people/386", make(map[string]string)), "GetPeopleByIdResponseMock", mock)
	defer unsetResponseMock()

	// Get deputy details by API id
	res, err := CLI.GetPeopleByID(386)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Exactly(t, mock, res)
}

func TestClient_GetPolicyByID(t *testing.T) {
	mock := new(GetPolicyByIDResponse)
	setResponseMock(CLI.getRequestURL("policies/0", make(map[string]string)), "ErrorResponseMock", mock)
	defer unsetResponseMock()

	// Try invalid ID
	_, err := CLI.GetPolicyByID(0)
	assert.NotNil(t, err)
	assert.Equal(t, "Not Found", err.Message)

	setResponseMock(CLI.getRequestURL("policies/1", make(map[string]string)), "GetPeopleByIdResponseMock", mock)
	defer unsetResponseMock()

	// Get policy details by API id
	res, err := CLI.GetPolicyByID(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Exactly(t, mock, res)
}

func TestClient_GetDivisionByID(t *testing.T) {
	mock := new(GetDivisionByIDResponse)
	setResponseMock(CLI.getRequestURL("divisions/0", make(map[string]string)), "ErrorResponseMock", mock)
	defer unsetResponseMock()

	// Try invalid ID
	_, err := CLI.GetDivisionByID(0)
	assert.NotNil(t, err)
	assert.Equal(t, "Not Found", err.Message)

	setResponseMock(CLI.getRequestURL("divisions/4896", make(map[string]string)), "GetPeopleByIdResponseMock", mock)
	defer unsetResponseMock()

	// Get policy details by API id
	res, err := CLI.GetDivisionByID(4896)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Exactly(t, mock, res)
}

func setResponseMock(path string, fileName string, target interface{}) {
	dat, err := ioutil.ReadFile(fmt.Sprintf("../tests/%s.json", fileName))
	if err != nil {
		panic(err)
	}
	httpmock.Activate()
	httpmock.RegisterResponder("GET", path, httpmock.NewStringResponder(200, string(dat)))
	json.Unmarshal(dat, target)
}

func unsetResponseMock() {
	httpmock.DeactivateAndReset()
}
