package data

import (
	"time"
)

type Thread struct {
	Id       int
	Uuid     string
	Topic    string
	UserId   int
	CreatedAt time.Time
}

type Post struct {
	Id       int
	Uuid     string
	Body     string
	UserId   int
	Thread   int
	CreatedAt time.Time
}

func Threads() (threads []Thread,err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	return
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id=$1", thread.Id)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	return
}