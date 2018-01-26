package rada4you

// GetAllPeoplesResponse struct represent API /peoples response
type GetAllPeoplesResponse struct {
	Peoples []Person
}

// GetAllPoliciesResponse struct represent API /policies response
type GetAllPoliciesResponse struct {
	Policies []Policy
}

// GetAllDivisionsResponse struct represent API /divisions response
type GetAllDivisionsResponse struct {
	Divisions []Division
}

// GetPeopleByIDResponse struct represent API /peoples/X response
type GetPeopleByIDResponse struct {
	Person
	PolicyComparisons []PolicyComparisons `json:"policy_comparisons"`
	Rebellions        int                 `json:"rebellions"`
	VotesAttended     int                 `json:"votes_attended"`
	VotesPossible     int                 `json:"votes_possible"`
}

// GetPolicyByIDResponse struct represent /policies/X.json response
type GetPolicyByIDResponse struct {
	Policy
	PeopleComparisons `json:"people_comparisons"`
	PolicyDivisions   []struct {
		Division `json:"division"`
		Strong   bool   `json:"strong"`
		Vote     string `json:"aye"`
	} `json:"policy_divisions"`
	Provisional bool `json:"provisional"`
}

type GetDivisionByIDResponse struct {
	Division
	Bills           []Bill `json:"bills"`
	PolicyDivisions []struct {
		Policy Policy `json:"policy"`
		Strong bool   `json:"strong"`
		Vote   string `json:"vote"`
	} `json:"policy_divisions"`
	Summary string `json:"summary"`
	Votes   []struct {
		Member `json:"member"`
		Vote   string `json:"vote"`
	}
}

// ErrorResponse struct represent error API response
type ErrorResponse struct {
	Message string `json:"error"`
}

func (e *ErrorResponse) IsOccur() bool {
	return e.Message != ""
}
