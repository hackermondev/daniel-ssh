package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
  "encoding/json"
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

type Blog struct{
  Id int `json:id`
  Title string `json:title`
  Teaser string `json:teaser`
  Data string `json:html_content`
  Slug string `json:slug`
  ViewCount int `json:viewCount`
  PublishedAt string `json:publishedAt`
}

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

func GetBlogs() ([]Blog, error){
 resp, err := http.Get(fmt.Sprintf("%s/blog/api", BaseURL))

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