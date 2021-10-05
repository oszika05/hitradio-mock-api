package news

import "time"

type News struct {
	Id      string    `json:"id"`
	Title   string    `json:"title"`
	Picture string    `json:"picture"`
	Date    time.Time `json:"date"`
	Tags    []string  `json:"tags"`
	Content string    `json:"content"`
}
