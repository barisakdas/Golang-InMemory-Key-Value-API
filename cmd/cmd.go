package cmd

import (
	"net/http"
	. "ys/Handlers"
	. "ys/Helpers"
)

func Run() {
	http.HandleFunc("/", GetAllEndpoints)
	http.HandleFunc("/get-access-token", GetAccessToken)
	http.Handle("/get-all-keys", IsAuthorized(GetAllKeys))
	http.Handle("/get", IsAuthorized(GetKey))
	http.Handle("/set", IsAuthorized(SetKeys))
	http.Handle("/delete-all", IsAuthorized(DeleteAllKeys))

	http.ListenAndServe(":8080", nil)
}
