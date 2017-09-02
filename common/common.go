package common

import (
	"errors"
	"net/http"
)

func GetCSRF() (csrfToken string, err error) {
	response, err := http.Get(URL_KRASINFOCOM)
	if err != nil {
		return "", err
	}

	cookies := response.Cookies()
	for i := range cookies {
		if cookies[i].Name == "csrf_token" {
			return cookies[i].Value, nil
		}
	}

	return "", errors.New("CSRF Token not found")
}
