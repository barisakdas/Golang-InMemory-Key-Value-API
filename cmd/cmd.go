package cmd

import (
	"net/http"
	. "ys/Handlers"
)

func Run() {
	http.HandleFunc("/", GetAllEndpoints)
	http.HandleFunc("/get-all-keys", GetAllKeys)
	http.HandleFunc("/get", GetKey)
	http.HandleFunc("/set", SetKeys)
	http.HandleFunc("/delete-all", DeleteAllKeys)

	http.ListenAndServe(":8080", nil)
}
