package model

import "time"

type Link struct {
	Id        string    `json:"string"`
	ShortLink string    `json:"string"`
	LongLink  string    `json:"string"`
	CreatedAt time.Time `json:"created_at"`
}
