package main

import (
	"net/http"
	"os"
	"log"
	"encoding/json"
	"casual-talk/route"
	"time"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

func init() {
	loadConfig()
	file, err := os.OpenFile("casual-talk.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
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
		"/":    route.Index,
		"/err": route.Err,

		// defined in route/auth.go
		"/login":          route.Login,
		"/logout":         route.Logout,
		"/signup":         route.Signup,
		"/signup_account": route.SignupAccount,
		"/authenticate":   route.Authenticate,

		// defined in route/thread.go
		"/thread/new":    route.NewThread,
		"/thread/create": route.CreateThread,
		"/thread/post":   route.PostThread,
		"/thread/read":   route.ReadThread,
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
