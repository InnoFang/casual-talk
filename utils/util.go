package utils

import (
	"net/http"
	"casual-talk/data"
	"errors"
)

func Session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); ok {
			err = errors.New("invalid session")
		}
	}
	return
}

