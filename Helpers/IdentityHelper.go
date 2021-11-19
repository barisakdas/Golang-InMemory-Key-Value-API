package Helpers

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var mySigningKey = []byte("a65ed4d2-e7ba-4a18-8891-91d76cb41cdd")

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func GetAccessToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "client"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
	}

	tokenResp := `{"Started Time":`+time.Now().Local().String()+`,"End Time":`+time.Now().Local().Add(time.Minute*30).String()+`,"Expire Time": "30 min","Token": "` + tokenString + `"}`
	w.Write([]byte(tokenResp))
}
