package data

import "time"

type Session struct {
	Id        int
	Uuid      string // 随机生成的唯一Id，会话机制的核心
	Email     string
	UserId    int
	CreatedAt time.Time
}

// check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid=$1", session.Uuid).
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
	stmt, err := Db.Prepare("DELETE FROM sessions WHERE uuid=$1")
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
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id=$1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// delete all sessions from database
func SessionDeleteAll() (err error) {
	_, err = Db.Exec("DELETE FROM sessions")
	return
}