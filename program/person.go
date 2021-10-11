package program

import "gopkg.in/guregu/null.v4"

type PersonType string

type Person struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Type        PersonType  `json:"type"`
	Picture     null.String `json:"picture"`
	Description null.String `json:"description"` // ?
}

const (
	PersonTypeHost   PersonType = "host"
	PersonTypeGuest  PersonType = "guest"
	PersonType_Empty PersonType = ""
)
