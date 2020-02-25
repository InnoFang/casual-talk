package route

import "net/http"

// GET /threads/new
// show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {

}

// POST /signup
// create the user account
func CreateThread(writer http.ResponseWriter, request *http.Request) {

}

// GET /thread/read
// show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {

}

// POST /thread/post
// create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {

}