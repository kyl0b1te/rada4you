package rada4you

// Person struct represent current deputy model
type Person struct {
	ID           int `json:"id"`
	LatestMember `json:"latest_member"`
}

// LatestMember struct represent deputy details
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

// Division struct represent current division
type Division struct {
	ID              int    `json:"id"`
	AyeVotes        int    `json:"aye_votes"`
	ClockTime       string `json:"clock_time"`
	Date            string `json:"date"`
	Edited          bool   `json:"edited"`
	House           string `json:"house"`
	Name            string `json:"name"`
	NoVotes         int    `json:"no_votes"`
	Number          int    `json:"number"`
	PossibleTurnout int    `json:"possible_turnout"`
	Rebellions      int    `json:"rebellions"`
}

type Bill struct {
	ID         int    `json:"id"`
	OfficialID string `json:"official_id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
}

type Member struct {
	ID         int    `json:"id"`
	Electorate string `json:"electorate"`
	House      string `json:"house"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Party      string `json:"party"`
	Person     struct {
		ID int `json:"id"`
	} `json:"person"`
}

type Policy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Provisional bool   `json:"provisional"`
}

type PolicyComparisons struct {
	Comparisons
	Policy `json:"policy"`
}

type PeopleComparisons struct {
	Comparisons
	Person `json:"person"`
}

type Comparisons struct {
	Agreement string `json:"agreement"`
	Voted     bool   `json:"voted"`
}
