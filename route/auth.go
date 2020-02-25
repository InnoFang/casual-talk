package route

import "net/http"

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

}

// GET /logout
// logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {

}