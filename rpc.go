package rpc

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
)

var Config = &config{}

type Rpc struct {
	worker *Worker
}

func init() {
	parseFlags()
}

func (rpc *Rpc) Run() {
	if Config.IsService() {
		if Config.appUrl == "" {
			flag.PrintDefaults()
			Logger.Fatal().Msgf("application hook URL is required")
		}

		runService()
	} else if Config.serve {
		if Config.appUrl == "" {
			flag.PrintDefaults()
			Logger.Fatal().Msgf("application hook URL is required")
		}

		rpc.serve(Config.addr)
	}
}

func (rpc *Rpc) serve(addr string) {
	go func() {
		rpc.worker = &Worker{NewQueue("default")}
		rpc.worker.DoWork()
	}()

	actions := Config.GetActions()

	for uri, _ := range actions {
		http.HandleFunc(uri, rpc.getRequest)
	}

	Logger.Info().Msgf("Listening on %s", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		Logger.Fatal().Msgf("Cannot start http server: %v", err)
	}
}

func (rpc *Rpc) getRequest(w http.ResponseWriter, req *http.Request) {
	secret := req.Header.Get("X-SECRET")
	appAuth := req.Header.Get("X-APP-AUTH")

	if secret != Config.appSecret {
		Logger.Warn().Msgf("Request declined - wrong secret")
		http.Error(w, "Wrong secret", http.StatusUnauthorized)
		return
	}

	if req.ContentLength == 0 {
		Logger.Warn().Msgf("Request declined - bad data")
		http.Error(w, "Bad data", http.StatusBadRequest)
		return
	}

	action := Config.GetAction(req.RequestURI)
	if action == nil {
		Logger.Warn().Msgf("Request declined - action not found")
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	Logger.Info().Msgf("Request %s accepted (%v)", req.RequestURI, req.Body)

	requestAcceptedResponse, err := (*action)(rpc, req.Body, appAuth)
	if err != nil {
		Logger.Warn().Msgf("Request declined - bad action result: %v", err)
		http.Error(w, "Bad result", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(requestAcceptedResponse)
	if err != nil {
		Logger.Error().Msgf("Failed to send request: %v", err)
	}
}

func (rpc *Rpc) AddJob(job Job) {
	rpc.worker.Queue.AddJob(job)
}

func (rpc *Rpc) SendResult(response ResultResponse, appAuth string) {
	data, _ := json.Marshal(response)

	appRequest, _ := http.NewRequest("POST", Config.appUrl, bytes.NewBuffer(data))
	appRequest.Header.Set("Content-Type", "application/json")
	appRequest.Header.Set("X-APP-AUTH", appAuth)

	client := &http.Client{}
	_, err := client.Do(appRequest)

	if err != nil {
		Logger.Warn().Msgf("Failed to send results: %v", err)
	} else {
		Logger.Info().Msgf("Results sent to %s", Config.appUrl)
	}
}
