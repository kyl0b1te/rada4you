package rada4you

// GetAllPeoplesResponse struct that represent general deputy information with API id
type GetAllPeoplesResponse struct {
	ID           int `json:"id"`
	LatestMember `json:"latest_member"`
}

// GetPeopleByIdResponse struct that represent deputy details by API id
type GetPeopleByIdResponse struct {
	ID                int `json:"id"`
	LatestMember      `json:"latest_member"`
	Offices           []string            `json:"-"`
	PolicyComparisons []PolicyComparisons `json:"policy_comparisons"`
	Rebellions        int                 `json:"rebellions"`
	VotesAttended     int                 `json:"votes_attended"`
	VotesPossible     int                 `json:"votes_possible"`
}

// LatestMember struct that represent general deputy information
type LatestMember struct {
	ID         int    `json:"id"`
	Electorate string `json:"electorate"`
	House      string `json:"house"`
	Name       struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Party string `json:"party"`
}

type PolicyComparisons struct {
	Agreement string `json:"agreement"`
	Policy    `json:"policy"`
	Voted     bool `json:"voted"`
}

// Policy struct that represent the policy information
type Policy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Provisional bool   `json:"provisional"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}
