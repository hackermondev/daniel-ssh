package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	BaseURL    = "https://daniel.is-a.dev"
	BaseAPIURL = fmt.Sprintf("%s/api", BaseURL)

	HTTPTransport = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: false,
	}

	HTTPClient = &http.Client{
    Transport: HTTPTransport,
  }
)

func GetAboutMeDescription() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/about", BaseAPIURL))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
