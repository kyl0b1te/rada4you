package rada4you

import (
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
	// Try get all peoples
	res, err := CLI.GetAllPeoples()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.True(t, len(res) > 0)

	for _, dep := range res {
		assert.NotZero(t, dep.ID)
	}
}

func TestClient_GetPeopleById(t *testing.T) {
	// Try invalid ID
	_, err := CLI.GetPeopleById(0)
	assert.NotNil(t, err)
	assert.Equal(t, "Not Found", err.Message)

	// Get deputy details by API id
	all, _ := CLI.GetAllPeoples()
	res, err := CLI.GetPeopleById(all[0].ID)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotZero(t, res.ID)
}

func TestClient_GetAllPolitics(t *testing.T) {
	// Try get all politics
	res, err := CLI.GetAllPolitics()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.True(t, len(res) > 0)

	if len(res) > 0 {
		pol := res[rand.Intn(len(res))]
		assert.NotZero(t, pol.ID)
		assert.NotZero(t, pol.Description)
		assert.NotZero(t, pol.Name)
	}
}
