package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// App inject dependencies
type App struct {
	Cache *Cache
}

// Message the request sent from client
type Message struct {
	Key  string `json:"key,omitempty"`
	Data string `json:"data,omitempty"`
}

// Ping sends "PONG" back for server status
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))
}

// WriteHandler updates or add a new value to the cache
func (a *App) WriteHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	a.Cache.Put(msg.Key, msg.Data)
	w.WriteHeader(201)
}

// ReadHandler fetches a value with the given key in the cache
func (a *App) ReadHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	value, err := a.Cache.Get(msg.Key)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	if str, ok := value.(string); ok {
		msg = Message{msg.Key, str}
	} else {
		http.Error(w, errors.New("Could not convert value to a string").Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// BuildServer build out server with handlers already defined
func BuildServer(a *App) *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/ping", Ping).Methods("GET")
	router.HandleFunc("/", a.ReadHandler).Methods("GET")
	router.HandleFunc("/", a.WriteHandler).Methods("POST")

	return &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
}
