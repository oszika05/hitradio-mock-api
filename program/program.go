package program

import "gopkg.in/guregu/null.v4"

type Program struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Picture     string      `json:"picture"`
	Description null.String `json:"description"`
}
