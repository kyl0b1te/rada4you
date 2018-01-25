package rada4you

// GetAllPeoplesResponse struct that represent general deputy information with API id
type GetAllPeoplesResponse struct {
	ID           int `json:"id"`
	LatestMember `json:"latest_member"`
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
