package route

import (
	"net/http"
	"casual-talk/data"
)

// GET /login
// show the login page
func Login(writer http.ResponseWriter, request *http.Request) {

}

// GET /signup
// show the signup page
func Signup(writer http.ResponseWriter, request *http.Request) {

}

// POST /signup
// create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {

}

// POST /authenticate
// authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	user, _ := data.UserByEmail(request.PostFormValue("email"))
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

// GET /logout
// logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {

}
