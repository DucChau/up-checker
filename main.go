package main

import (
        "encoding/json"
        "log"
        "net/http"
        "time"
)

type In struct {
        URL string `json:"url"`
}

type Out struct {
        Time   int64  `json:"time"`
        URL    string `json:"url"`
        Code   int    `json:"code"`
        Status string `json:"status"`
}

func checkHealth(rw http.ResponseWriter, req *http.Request) {
        // get the epoch
        now := time.Now()
        epoch := now.Unix()

        var in In

        // decode json body
        err := json.NewDecoder(req.Body).Decode(&in)

        // barf on error
        if err != nil {
                log.Print(err.Error())
                http.Error(rw, err.Error(), http.StatusBadRequest)
                return
        }

        log.Print(in.URL)

        // execute get on URL
        checkResponse, err := http.Get(in.URL)

        // get failed, return error message with code of 500
        if err != nil {
                log.Print(err.Error())
                var checkError = Out{epoch, in.URL, 500, err.Error()}
                r, _ := json.Marshal(checkError)
                rw.WriteHeader(http.StatusOK)
                rw.Write([]byte(r))
                return
        }

        // if we got here URL check worked
        var out = Out{epoch, in.URL, checkResponse.StatusCode, checkResponse.Status}

        // convert to JSON
        r, err := json.Marshal(out)

        rw.Header().Set("Content-Type", "application/json")

        // return results
        if err != nil {
                log.Print(err.Error())
                rw.WriteHeader(http.StatusInternalServerError)
                rw.Write([]byte("fail"))
        } else {
                rw.WriteHeader(http.StatusOK)
                rw.Write([]byte(r))
        }
}

func main() {
        mux := http.NewServeMux()
        mux.HandleFunc("/v1/health", checkHealth)
        err := http.ListenAndServe(":8000", mux)
        log.Fatal(err)
}
