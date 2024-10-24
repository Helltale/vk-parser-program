package fetcher

import "time"

type Url struct {
	Link string
	Time time.Time
}

func NewUrl(link string) *Url {
	return &Url{
		Link: link,
		Time: time.Now(),
	}
}
