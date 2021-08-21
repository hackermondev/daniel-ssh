package api

import (
	"encoding/json"
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

type Blog struct {
	Id          int    `json:id`
	Title       string `json:title`
	Teaser      string `json:teaser`
	Data        string `json:data`
	Slug        string `json:slug`
	ViewCount   int    `json:viewCount`
	PublishedAt string `json:publishedAt`
}

func GetAboutMeDescription() (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/about", BaseAPIURL), nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "daniel-ssh")

	resp, err := HTTPClient.Do(req)

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

func GetBlogs() ([]Blog, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blog/api", BaseURL), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "daniel-ssh")

	resp, err := HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data []Blog

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
