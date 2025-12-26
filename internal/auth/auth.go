package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}
	headerContent := strings.Split(authHeader, " ")

	if len(headerContent) != 2 {
		return "", errors.New("malformed authorization header")

	}
	if headerContent[0] != "ApiKey" {
		return "", errors.New("malformed first param authorization header")
	}

	return headerContent[1], nil
}
