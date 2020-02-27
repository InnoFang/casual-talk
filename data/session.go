package data

import (
	"time"
	"net/http"
	"errors"
)

type Session struct {
	Id        int
	Uuid      string // 随机生成的唯一Id，会话机制的核心
	Email     string
	UserId    int
	CreatedAt time.Time
}

// check if session is valid in the database
func (session *Session) Valid() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid=?", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// delete session from database
func (session *Session) DeleteByUUID() (err error) {
	stmt, err := Db.Prepare("DELETE FROM sessions WHERE uuid=?")
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(session.Uuid)
	return
}

// get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id=?", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// delete all sessions from database
func SessionDeleteAll() (err error) {
	_, err = Db.Exec("DELETE FROM sessions")
	return
}

// checks if the user is logged in and has a session, if not err is not nil
func SessionCheck(writer http.ResponseWriter, request *http.Request) (sess Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = Session{Uuid: cookie.Value}
		if ok, _ := sess.Valid(); ok {
			err = errors.New("invalid session")
		}
	}
	return
}