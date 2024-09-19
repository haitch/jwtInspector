package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	http.HandleFunc("/", TokenDump)
	http.ListenAndServe(":8080", nil)
}

type TokenInspectResponse struct {
	Claims map[string]interface{} `json:"claims"`
}

func TokenDump(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" || len(bearerToken) < 7 || strings.ToLower(bearerToken[0:7]) != "bearer " {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 Unauthorized"))
		return
	}

	bearerToken = bearerToken[7:]
	// TODO, handle error
	token, _ := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		// TODO: Return the secret signing key
		return []byte(""), nil
	})

	if token == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 Unauthorized"))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		resp := TokenInspectResponse{}
		resp.Claims = make(map[string]interface{})
		for key, val := range claims {
			resp.Claims[key] = val
		}

		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 Unauthorized"))
	}
}
