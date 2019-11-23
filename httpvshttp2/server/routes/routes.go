package routes

import (
	"encoding/json"
	"httpvshttp2/server/defn"
	"io"
	"log"
	"net/http"
	"os"
)

func ServeData(w http.ResponseWriter, r *http.Request) {
	incomingRequest := &Request{}
	logger := &defn.LoggerDefn{}
	err := json.NewDecoder(r.Body).Decode(incomingRequest)
	if err != nil {
		if err == io.EOF {
			http.Error(w, "NoContent", http.StatusNoContent)
			return
		}
		log.Print(logger.SetLogggerDefn(err, "Error while Decoding Request", "ServeData"))
		http.Error(w, err.Error(), http.StatusPartialContent)
	}
	defer r.Body.Close()

	file, err := os.Open("example.json")
	if err != nil {
		log.Print(logger.SetLogggerDefn(err, "Error while opening file", "ServeData"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	exampleMapObj := make(map[int]string, 0)
	err = json.NewDecoder(file).Decode(&exampleMapObj)
	if err != nil {
		log.Print(logger.SetLogggerDefn(err, "Error while decoding data from file", "ServeData"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseObj := &Response{}
	for key, val := range exampleMapObj {
		if key == incomingRequest.Key {
			ResponseObj.Value = val
		} else {
			ResponseObj.Value = ""
		}
	}

	err = json.NewEncoder(w).Encode(ResponseObj)
	if err != nil {
		log.Print(err, "Error while encoding response to responsewriter", "ServeData")
		return
	}

}
