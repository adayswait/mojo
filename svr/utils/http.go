package utils

import (
	"github.com/valyala/fasthttp"
)

func httpRequest(uri, method, contentType, body string) ([]byte, error) {
	req := &fasthttp.Request{}
	req.SetRequestURI(uri)
	req.SetBody([]byte(body))

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType(contentType)
	req.Header.SetMethod(method)

	res := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, res); err != nil {
		return nil, err
	}
	return res.Body(), nil
}

func HttpRequestRaw(uri, method, contentType string, body []byte) ([]byte, error) {
	req := &fasthttp.Request{}
	req.SetRequestURI(uri)
	req.SetBody(body)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType(contentType)
	req.Header.SetMethod(method)

	res := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, res); err != nil {
		return nil, err
	}
	return res.Body(), nil
}

func HttpGet(uri, body string) ([]byte, error) {
	return httpRequest(uri, "GET", "application/json", body)
}
func HttpPost(uri, body string) ([]byte, error) {
	return httpRequest(uri, "POST", "application/json", body)
}
