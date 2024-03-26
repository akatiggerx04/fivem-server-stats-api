package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"time"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type VanityCodes struct {
    VanityCodes map[string]string `yaml:"vanity_codes"`
}

var vanityCodes VanityCodes

func loadVanityCodes(filename string) error {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }

    if err := yaml.Unmarshal(data, &vanityCodes); err != nil {
        return err
    }

    return nil
}

func main() {
	if err := loadVanityCodes("vanity.yaml"); err != nil {
        fmt.Println("Error loading server vanity codes:", err)
        return
    }

    r := mux.NewRouter()
    r.HandleFunc("/{server_code}", func(w http.ResponseWriter, r *http.Request) {
		GetServerInfo(w, r, false)
	}).Methods("GET")

	r.HandleFunc("/original/{server_code}", func(w http.ResponseWriter, r *http.Request) {
		GetServerInfo(w, r, true)
	}).Methods("GET")

    http.Handle("/", r)

	log.Println("FiveM Servers API by @akatiggerx04 | Starting API on :7747...")
    log.Fatal(http.ListenAndServe(":7747", r))
}

func ServerisOnline(utcTimeString string) bool {
    utcTime, err := time.Parse(time.RFC3339Nano, utcTimeString)
    if err != nil {
		fmt.Println("Error parsing UTC time:", err)
        return false
    }

    currentTime := time.Now().UTC()
    timeDifference := currentTime.Sub(utcTime)

	return timeDifference <= time.Minute * 5
}