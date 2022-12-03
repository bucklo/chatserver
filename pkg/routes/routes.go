package routes

import (
	"chatserver/pkg/db"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func logRequest(r *http.Request) {
	log.Printf("%v %v %v %v", r.RemoteAddr, r.Header["User-Agent"][0], r.Method, r.URL)
}

func Register(w http.ResponseWriter, req *http.Request) {
	type Message struct {
		Login    string
		Password string
	}

	var m Message

	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &m)
	if err != nil {
		log.Print(err)
	}

	logRequest(req)

	db.AddUser(m.Login, m.Password)
}

func Login(w http.ResponseWriter, req *http.Request) {
	type Message struct {
		Login    string
		Password string
	}

	var m Message

	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &m)
	if err != nil {
		log.Print(err)
	}

	logRequest(req)

	if db.AuthenticateUser(m.Login, m.Password) {
		log.Printf("User %v authenticated", m.Login)
	} else {
		log.Printf("User %v failed to authenticate", m.Login)
	}
}
