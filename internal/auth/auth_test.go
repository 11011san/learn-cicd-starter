package auth

import (
	"net/http"
	"testing"
)

func TestBearerToken(t *testing.T) {
	headers := http.Header{}
	bearerToken, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("Should not found a token")
		return
	}
	if bearerToken != "" {
		t.Errorf("Should return empty token: %v", bearerToken)
		return
	}
	token := "1234567890sdfghjh"
	headers.Add("Authorization", token)

	bearerToken, err = GetAPIKey(headers)
	if err == nil {
		t.Errorf("Should not found a token")
		return
	}
	if bearerToken != "" {
		t.Errorf("Should return empty token: %v", bearerToken)
		return
	}

	headers.Set("Authorization", "ApiKey "+token)
	bearerToken, err = GetAPIKey(headers)
	if err != nil {
		t.Errorf("Should have rtrived a token")
		return
	}
	if bearerToken != token {
		t.Errorf("Should return token: %v", bearerToken)
		return
	}
}
