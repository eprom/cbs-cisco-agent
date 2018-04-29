package common

import (
	"errors"
	"fmt"
	"github.com/jakhog/cbs-cisco-agent/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func performAuthenticatedRequest(method, path, contentType string, body io.Reader) (io.ReadCloser, error) {
	cfg := config.GetConfig().Cisco
	// Prepare the request
	url := fmt.Sprintf("http://%v:%v/level/15/exec/-%v", cfg.Server, cfg.Port, path)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(cfg.Username, cfg.Password)
	if method == "POST" {
		request.Header.Set("Content-Type", contentType)
	}
	// Perform request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	// Make sure it was successful
	if response.StatusCode != 200 {
		response.Body.Close()
		return nil, errors.New("Non 200 response code")
	}
	return response.Body, nil
}

func RequestGet(path string) (io.ReadCloser, error) {
	return performAuthenticatedRequest("GET", path, "", nil)
}

func RequestPost(path string, form url.Values) (io.ReadCloser, error) {
	return performAuthenticatedRequest("POST", path, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
}
