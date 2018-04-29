package common

import (
	"io"
	"net/url"
)

func PostCommand(path string) (io.ReadCloser, error) {
	// First, get an CSRF token
	token, err := GetCSRFToken(path)
	if err != nil {
		return nil, err
	}
	// Then, call command to get the data
	form := url.Values{}
	form.Add("csrf_token", token)
	form.Add("CMD", "CR")
	body, err := RequestPost(path, form)
	if err != nil {
		return nil, err
	}
	// Return body
	return body, nil
}

func PostCommandSimple(path string) error {
	body, err := PostCommand(path)
	if err != nil {
		return err
	} else {
		return body.Close()
	}
}
