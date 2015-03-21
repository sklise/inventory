package models

import "time"

type Format struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Thing struct {
	Id        int64
	Year      int64
	Title     string
	AuthorId  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Author struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Publisher struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

type AuthorAndThings struct {
	Author Author
	Things []Thing
}
