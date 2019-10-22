package main

import (
	"encoding/json"
	spec "github.com/Tayeb-Ali/rabbit/spec"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// keep waiting channels for reply messages from rabbit
var rchans = make(map[string](chan spec.CreateDocumentReply))

func initApi() {
	// router
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/documents", apiDocument).Methods("POST")

	log.Printf("INFO: init http api")

	// start server
	err := http.ListenAndServe(":7654", r)
	if err != nil {
		log.Printf("ERROR: fail init http server, %s", err.Error)
		os.Exit(1)
	}
}

func apiDocument(w http.ResponseWriter, r *http.Request) {
	// read body
	data, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	log.Printf("INFO: document request: %v", string(data))

	// unmarshal and create docMsg
	doc := &spec.Document{}
	err := json.Unmarshal(data, doc)
	if err != nil {
		log.Printf("ERROR: fail unmarshl: %s", err.Error)
		response(w, "Invalid request json", 400)
	}
	docMsg := &spec.CreateDocumentMessage{
		Uid:      uid(),
		Document: doc,
		ReplyTo:  "gateway",
	}

	log.Printf("INFO: document message: %v", docMsg)

	// create channel and add to rchans with uid
	rchan := make(chan spec.CreateDocumentReply)
	rchans[docMsg.Uid] = rchan

	// publish rabbit message
	msg := RabbitMsg{
		QueueName: "storage",
		Message:   *docMsg,
	}
	pchan <- msg

	// wait for reply
	waitReply(docMsg.Uid, rchan, w)
}

func waitReply(uid string, rchan chan spec.CreateDocumentReply, w http.ResponseWriter) {
	for {
		select {
		case docReply := <-rchan:
			// responses received
			log.Printf("INFO: received reply: %v uid: %s", docReply, uid)

			// send response back to client
			response(w, "Created", 201)

			// remove channel from rchans
			delete(rchans, uid)
			return
		case <-time.After(10 * time.Second):
			// timeout
			log.Printf("ERROR: request timeout uid: %s", uid)

			// send response back to client
			response(w, "Timeout", 408)

			// remove channel from rchans
			delete(rchans, uid)
			return
		}
	}
}

func response(w http.ResponseWriter, resp string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, resp)
}
