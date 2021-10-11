package program

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Episode struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Date        time.Time   `json:"date"`
	Description null.String `json:"description"`
	Tags        []string    `json:"tags"`
	Program     Program     `json:"program"`

	AudioUrl string `json:"audioUrl"`

	Hosts  []Person `json:"hosts"`
	Guests []Person `json:"guests"`
}

func (e Episode) Equals(other Episode) bool {
	return e.Id == other.Id
}

func (e Episode) NotEquals(other Episode) bool {
	return e.Id != other.Id
}
