package rada4you

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllDivisionsRequest_Values(t *testing.T) {
	// Test empty request
	r := GetAllDivisionsRequest{}
	assert.Empty(t, r.Values())

	// Test without dates
	r = GetAllDivisionsRequest{House: "rada"}
	ex := make(map[string]string)
	ex["house"] = "rada"
	assert.Exactly(t, ex, r.Values())

	// Test with one date
	r = GetAllDivisionsRequest{Start: time.Now()}
	assert.Empty(t, r.Values())

	// Test with two dates
	current := time.Now()
	r = GetAllDivisionsRequest{Start: current, End: current}
	ex = make(map[string]string)
	ex["start_date"] = current.Format("2006-01-02")
	ex["end_date"] = current.Format("2006-01-02")
	assert.Exactly(t, ex, r.Values())

	// Test with all parameters
	r = GetAllDivisionsRequest{Start: current, End: current, House: "rada"}
	ex["house"] = "rada"
	assert.Exactly(t, ex, r.Values())
}
