package web

import (
	"io/ioutil"
	"net/http"
)

type Client interface {
	Get(url string) string
}

type HttpClient struct {
}

func (hc HttpClient) Get(url string) string {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return ""
	}

	byteArray, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return ""
	}

	return string(byteArray)
}

func CreateClient() Client {
	return HttpClient{}
}