package main

import (
	"crypto"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/golang-jwt/jwt/v5"
)

var signingKey crypto.PrivateKey

func init() {
	_, key, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	signingKey = key
}

func generateToken() (string, error) {
	t := jwt.NewWithClaims(&jwt.SigningMethodEd25519{},
		jwt.MapClaims{
			"iss": "testIssuer",
			"sub": "john",
			"foo": 2,
		})
	return t.SignedString(signingKey)
}

func tokenDump(token string) (TokenInspectResponse, error) {
	recorder := httptest.NewRecorder()
	http.HandleFunc("/", TokenDump)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	http.DefaultServeMux.ServeHTTP(recorder, req)
	if recorder.Code == http.StatusInternalServerError || recorder.Code == http.StatusUnauthorized {
		return TokenInspectResponse{}, fmt.Errorf("unexpected status code: %d", recorder.Code)
	}

	var resp TokenInspectResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		return TokenInspectResponse{}, err
	}
	return resp, nil
}
