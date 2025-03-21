package util

import (
	"github.com/go-resty/resty/v2"
)

var restClient *resty.Client

func init() {
	restClient = resty.New()
}

func Get(url string) *resty.Response {
	resp, err := restClient.R().Get(url)
	if err != nil {
		return nil
	}

	if resp.StatusCode() != 200 {
		return nil
	}

	return resp
}

func Post(url string, body any) *resty.Response {
	resp, err := restClient.R().SetBody(body).Post(url)
	if err != nil {
		return nil
	}

	if resp.StatusCode() != 200 {
		return nil
	}

	return resp
}
