package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error){

	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentaction found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authentaction")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("apiKey Not Found")
	}
	return vals[1], nil
}