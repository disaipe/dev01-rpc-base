package rpc

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
)

var AppConfig = &Config{}

type ActionFunction func(rpc *Rpc, body io.ReadCloser, appAuth string) (Response, error)

type Rpc struct {
	Action ActionFunction

	worker *Worker
}

func init() {
	parseFlags()
}

func (rpc *Rpc) Run() {
	if AppConfig.IsService() {
		if AppConfig.appUrl == "" {
			flag.PrintDefaults()
			log.Fatal("application hook URL is required")
		}

		runService()
	} else if AppConfig.serve {
		if AppConfig.appUrl == "" {
			flag.PrintDefaults()
			log.Fatal("application hook URL is required")
		}

		rpc.serve(AppConfig.addr)
	}
}

func (rpc *Rpc) serve(addr string) {
	go func() {
		rpc.worker = &Worker{NewQueue("default")}
		rpc.worker.DoWork()
	}()

	http.HandleFunc("/get", rpc.getRequest)
	log.Printf("Listening on %s", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalf("Cannot start http server: %v", err)
	}
}

func (rpc *Rpc) getRequest(w http.ResponseWriter, req *http.Request) {
	secret := req.Header.Get("X-SECRET")
	appAuth := req.Header.Get("X-APP-AUTH")

	if secret != AppConfig.appSecret {
		http.Error(w, "Wrong secret", http.StatusUnauthorized)
		return
	}

	if req.ContentLength == 0 {
		http.Error(w, "Bad data", http.StatusBadRequest)
		return
	}

	requestAcceptedResponse, err := rpc.Action(rpc, req.Body, appAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(requestAcceptedResponse)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
	}
}

func (rpc *Rpc) AddJob(job Job) {
	rpc.worker.Queue.AddJob(job)
}

func (rpc *Rpc) SendResult(response ResultResponse, appAuth string) {
	data, _ := json.Marshal(response)

	appRequest, _ := http.NewRequest("POST", AppConfig.appUrl, bytes.NewBuffer(data))
	appRequest.Header.Set("Content-Type", "application/json")
	appRequest.Header.Set("X-APP-AUTH", appAuth)

	client := &http.Client{}
	_, err := client.Do(appRequest)

	if err != nil {
		log.Printf("Failed to send results: %v", err)
	} else {
		log.Printf("Results sent to %s", AppConfig.appUrl)
	}
}
