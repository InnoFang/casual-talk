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


// create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	stmt, err := Db.Prepare(
		"INSERT INTO threads(uuid, topic, user_id, created_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	stmt, err := Db.Prepare(
		"INSERT INTO posts(uuid, body, user_id, thread_id, created_at) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}

// create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	stmt, err := Db.Prepare(
		"INSERT INTO sessions(uuid, email, user_id, created_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id=?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// create a new user, save user info into the database
func (user *User) Create() (err error) {
	stmt, err := Db.Prepare(
		"INSERT INTO users(uuid, name, email, password, created_at) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).
				Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// delete user from database
func (user *User) Delete() (err error) {
	stmt, err := Db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	return
}

// update user information in the database
func (user *User) Update() (err error) {
	statement := "UPDATE users SET name=?, email=? where id=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return
}

// delete all users from database
func UserDeleteAll() (err error) {
	_, err = Db.Exec("DELETE FROM users")
	return
}

// get all users in the database and returns it
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

// get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email=?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid=?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
