package earthquake

import (
	"strings"
)

type MockClient struct {
	urlWhenExecute   string
	returnJsonResult string
}

type MockClientFunc interface {
	Get(url string) string
	NoContainsExecutedUrl(url string) bool
	UrlWhenExecute() string
}


func (mc *MockClient) Get(url string) string {
	mc.urlWhenExecute = url
	return mc.returnJsonResult
}

func (mc MockClient) NoContainsExecutedUrl(url string) bool {
	return !strings.Contains(mc.urlWhenExecute, url)
}

func (mc MockClient) UrlWhenExecute() string {
	return mc.urlWhenExecute
}


func CreateMockClient(jsonResult string) MockClientFunc {
	result := new(MockClient)
	result.returnJsonResult = jsonResult
	return result
}
