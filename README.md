# rada4you [![Build Status](https://travis-ci.org/zhikiri/rada4you.svg?branch=master)](https://travis-ci.org/zhikiri/rada4you)

Ukrainian parliament activity API client build with golang. 
It was made for telegram bot to monitor votes, deputies stats and parliament activity in general.

We are using API `v1`, documentation should be accessible by [url](https://rada4you.org/help/data).

## Documentation

API client library document contain format of function call with response structure.

#### `Client.GetAllPeoples()`

**Description:** function for retrieve the list of all deputies in system.

Response format:

```go
type GetAllPeoplesResponse struct {
	Peoples []Person
}
```

---

#### `Client.GetPeopleByID(int)`

**Description:** function for retrieve deputy details by system database id.

Response format:

```go
type GetPeopleByIDResponse struct {
	Person
	PolicyComparisons []PolicyComparisons
	Rebellions        int
	VotesAttended     int
	VotesPossible     int
}
```

---

#### `Client.GetAllPolicies()`

**Description:** function for retrieve the list of all available policies.

Response format:

```go
type GetAllPoliciesResponse struct {
	Policies []Policy
}
```

---

#### `Client.GetPolicyByID(int)`

**Description:** function for retrieve target policy by system database id.

```go
type GetPolicyByIDResponse struct {
	Policy
	PeopleComparisons
	PolicyDivisions   []struct {
		Division
		Strong   bool
		Vote     string
	}
	Provisional bool
}
```

---

#### `Client.GetAllDivisions(GetAllDivisionsRequest)`

**Description:** function for retrieve the list of all available divisions.

```go
type GetAllDivisionsResponse struct {
	Divisions []Division
}
```

---

#### `Client.GetDivisionByID(int)`

**Description:** function for retrieve target division by system database id.

```go
type GetDivisionByIDResponse struct {
	Division
	Bills           []Bill
	PolicyDivisions []struct {
		Policy Policy
		Strong bool
		Vote   string
	}
	Summary string
	Votes   []struct {
		Member
		Vote  string
	}
}
```


## TODO

Here is the list of next steps:

- [X] complete the tests for "get all" methods
- [X] add tests for private methods
- [ ] add docker configuration files
- [ ] make a project refactoring

## Summary

Any contributions are welcome.
