package program

import "gopkg.in/guregu/null.v4"

type Program struct {
	Id          string
	Name        string
	Picture     string
	Description null.String
}
