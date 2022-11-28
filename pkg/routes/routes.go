package routes

import (
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
}
