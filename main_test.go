package main

import (
	"testing"
)

func TestTokenDump(t *testing.T) {

	token, err := generateToken()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := tokenDump(token)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Claims) != 3 {
		t.Fatalf("expected 3 claims, got %d", len(resp.Claims))
	}
}
