package rada4you

// GetAllPeoplesResponse struct represent API /peoples response
type GetAllPeoplesResponse struct {
	Peoples []Person
}

// GetPeopleByIdResponse struct represent API /peoples/X response
type GetPeopleByIdResponse struct {
	Person
	PolicyComparisons []PolicyComparisons `json:"policy_comparisons"`
	Rebellions        int                 `json:"rebellions"`
	VotesAttended     int                 `json:"votes_attended"`
	VotesPossible     int                 `json:"votes_possible"`
}

// GetAllPolicies struct represent API /policies response
type GetAllPolicies struct {
	Policies []Policy
}

// GetPolicyByIdResponse struct represent /policies/X.json response
type GetPolicyByIdResponse struct {
	Policy
	PeopleComparisons `json:"people_comparisons"`
	PolicyDivisions   []struct {
		Division `json:"division"`
		Strong   bool   `json:"strong"`
		Vote     string `json:"aye"`
	} `json:"policy_divisions"`
	Provisional bool `json:"provisional"`
}

// ErrorResponse struct represent error API response
type ErrorResponse struct {
	Message string `json:"error"`
}

func (e *ErrorResponse) IsOccur() bool {
	return e.Message != ""
}
