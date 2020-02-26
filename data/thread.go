package data

import "time"

type Thread struct {
	Id       int
	Uuid     string
	Topic    string
	UserId   int
	CreateAt time.Time
}

type Post struct {
	Id       int
	Uuid     string
	Body     string
	UserId   int
	Thread   int
	CreateAt time.Time
}

func Threads() ([]Thread, error) {
	return nil, nil
}