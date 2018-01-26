# rada4you

[![Build Status](https://travis-ci.org/zhikiri/rada4you.svg?branch=master)](https://travis-ci.org/zhikiri/rada4you)

It's a GO API client for the [API](https://rada4you.org/help/data).

Currently client is working with last API version, `v1`.

## TODO

Here is the list of next steps:

- [ ] complete the tests for "get all" methods
- [ ] add tests for private methods
- [ ] add docker configuration files

## Available methods

### `Client.GetAllPeoples()`

**Description:** function for retrieve the list of all deputies in system.

Response format:

```go
type GetAllPeoplesResponse struct {
	Peoples []Person
}
```

---

### `Client.GetPeopleByID(int)`

**Description:** function for retrieve deputy detail by system database id.

Response format:

```go
type GetPeopleByIDResponse struct {
	Person
	PolicyComparisons []PolicyComparisons `json:"policy_comparisons"`
	Rebellions        int                 `json:"rebellions"`
	VotesAttended     int                 `json:"votes_attended"`
	VotesPossible     int                 `json:"votes_possible"`
}
```

---

### `Client.GetAllPolicies()`

**Description:** function for retrieve the list of all available policies.

Response format:

```go
type GetAllPoliciesResponse struct {
	Policies []Policy
}
```

---

### `Client.GetPolicyByID(int)`

**Description:** function for retrieve target policy by system database id.

```go
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
```

---

### `Client.GetAllDivisions()`

**Description:** function for retrieve the list of all available divisions.

```go
type GetAllDivisionsResponse struct {
	Divisions []Division
}
```

---

### `Client.GetDivisionByID(int)`

**Description:** function for retrieve target division by system database id.

```go
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
```

Any contributions are welcome.

Feel free to create a new issues and PR's