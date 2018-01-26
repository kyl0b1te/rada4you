package rada4you

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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

func TestInvalidApiKey(t *testing.T) {
	cli := New("secret")
	msg := "You need a valid api key. Sign up for an account on Вони голосують для тебе to get one."
	if _, err := cli.GetAllPeoples(); err != nil {
		assert.Equal(t, msg, err.Message)
	}
}

func TestClient_GetAllPeoples(t *testing.T) {
	mock := new(GetAllPeoplesResponse)
	setResponseMock(CLI.getRequestURL("people"), "GetAllPeoplesResponseMock", mock)
	defer unsetResponseMock()

	// Try get all peoples
	res, err := CLI.GetAllPeoples()
	assert.Nil(t, err)
	assert.NotNil(t, res.Peoples)
	assert.True(t, len(res.Peoples) > 0)

	for _, dep := range res.Peoples {
		assert.NotZero(t, dep.ID)
	}
}

func TestClient_GetAllPolicies(t *testing.T) {
	mock := new(GetAllPoliciesResponse)
	setResponseMock(CLI.getRequestURL("policies"), "GetAllPoliciesResponseMock", mock)
	defer unsetResponseMock()

	// Try get all politics
	res, err := CLI.GetAllPolicies()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.True(t, len(res.Policies) > 0)

	if len(res.Policies) > 0 {
		pol := res.Policies[rand.Intn(len(res.Policies))]
		assert.NotZero(t, pol.ID)
		assert.NotZero(t, pol.Description)
		assert.NotZero(t, pol.Name)
	}
}

func TestClient_GetAllDivisions(t *testing.T) {
	mock := new(GetAllDivisionsResponse)
	setResponseMock(CLI.getRequestURL("divisions"), "GetAllDivisionsResponseMock", mock)
	defer unsetResponseMock()

	// Try get all divisions
	res, err := CLI.GetAllDivisions()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.True(t, len(res.Divisions) > 0)

	if len(res.Divisions) > 0 {
		pol := res.Divisions[rand.Intn(len(res.Divisions))]
		assert.NotZero(t, pol.ID)
		assert.NotZero(t, pol.Name)
	}
}

func TestClient_GetPeopleByID(t *testing.T) {
	mock := new(GetPeopleByIDResponse)
	setResponseMock(CLI.getRequestURL("people/0"), "ErrorResponseMock", mock)
	defer unsetResponseMock()

	// Try invalid ID
	_, err := CLI.GetPeopleByID(0)
	assert.NotNil(t, err)
	assert.Equal(t, "Not Found", err.Message)

	setResponseMock(CLI.getRequestURL("people/386"), "GetPeopleByIdResponseMock", mock)
	defer unsetResponseMock()

	// Get deputy details by API id
	res, err := CLI.GetPeopleByID(386)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Exactly(t, mock, res)
}

func TestClient_GetPolicyByID(t *testing.T) {
	mock := new(GetPolicyByIDResponse)
	setResponseMock(CLI.getRequestURL("policies/0"), "ErrorResponseMock", mock)
	defer unsetResponseMock()

	// Try invalid ID
	_, err := CLI.GetPolicyByID(0)
	assert.NotNil(t, err)
	assert.Equal(t, "Not Found", err.Message)

	setResponseMock(CLI.getRequestURL("policies/1"), "GetPeopleByIdResponseMock", mock)
	defer unsetResponseMock()

	// Get policy details by API id
	res, err := CLI.GetPolicyByID(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Exactly(t, mock, res)
}

func TestClient_GetDivisionByID(t *testing.T) {
	mock := new(GetDivisionByIDResponse)
	setResponseMock(CLI.getRequestURL("divisions/0"), "ErrorResponseMock", mock)
	defer unsetResponseMock()

	// Try invalid ID
	_, err := CLI.GetDivisionByID(0)
	assert.NotNil(t, err)
	assert.Equal(t, "Not Found", err.Message)

	setResponseMock(CLI.getRequestURL("divisions/4896"), "GetPeopleByIdResponseMock", mock)
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
