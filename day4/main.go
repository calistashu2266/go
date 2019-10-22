package main

import (
	"fmt"
	"net/http"
)

type AutotaskRequest struct {
	RequestID string     `json:"requestid"`
	Clone     CloneModel `json:"clone"`
	Push      PushModel  `json:"push"`
}

type CloneModel struct {
	//TODO
	//"Method": string `json:"ceph"`
	RequestID   string `json:"requestid"`
	CallbackURL string `json:"callbackurl"`
}

type PushModel struct {
	RequestID   string `json:"requestiD"`
	CallbackURL string `json:"callbackuRL"`
	IP          string `json:"remoteip"`
	Port        int    `json:"remoteport"`
	User        string `json:"user"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	// str := json.Unmarshal(&AutotaskRequest{})
	http.HandleFunc("/bb", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "weeeee")
	})

	http.ListenAndServe(":3001", nil)
}
