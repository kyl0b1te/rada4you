package rada4you

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var APIKey string

func init() {
	APIKey = os.Getenv("API_KEY")
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
	cli := New(APIKey)
	res, err := cli.GetAllPeoples()

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.True(t, len(*res) > 0)
}
