package main

import (
	"net/http"
	"casual-talk/routes"
	"time"
	"os"
	"log"
	"encoding/json"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func main() {
	mux := http.NewServeMux()

	// handle static assets
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mapper := map[string]func(http.ResponseWriter, *http.Request){
		// defined in route/index.go
		"/":    routes.Index,
		"/err": routes.Err,

		// defined in route/auth.go
		"/login":          routes.Login,
		"/logout":         routes.Logout,
		"/signup":         routes.Signup,
		"/signup_account": routes.SignupAccount,
		"/authenticate":   routes.Authenticate,

		// defined in route/thread.go
		"/thread/new":    routes.NewThread,
		"/thread/create": routes.CreateThread,
		"/thread/post":   routes.PostThread,
		"/thread/read":   routes.ReadThread,
	}

	for pattern, handler := range mapper {
		mux.HandleFunc(pattern, handler)
	}

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
