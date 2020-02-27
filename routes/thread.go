package routes

import (
	"net/http"
	"casual-talk/data"
	"fmt"
	"casual-talk/utils"
)

// GET /threads/new
// show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := data.SessionCheck(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// create the user account
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := data.SessionCheck(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			utils.Danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	uuid := request.URL.Query().Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot read thread")
	} else {
		_, err := data.SessionCheck(writer, request)
		if err != nil {
			utils.GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			utils.GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := data.SessionCheck(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			utils.ErrorMessage(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			utils.Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}