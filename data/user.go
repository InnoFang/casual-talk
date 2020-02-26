package data

import "time"

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string // 随机生成的唯一Id，会话机制的核心
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (u User) CreateSession() Session {
	return Session{}
}

func (s Session) Check() (bool, error){
	return false, nil
}

func UserByEmail(email string) (User, error) {
	return User{}, nil
}

func Encrypt(password string) string {
	return ""
}