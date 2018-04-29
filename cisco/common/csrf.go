package common

import (
	"errors"
	"golang.org/x/net/html"
	"io"
	"strings"
)

func findAttributeValue(attributes []html.Attribute, key string) (string, bool) {
	for _, attribute := range attributes {
		if strings.EqualFold(attribute.Key, key) {
			return attribute.Val, true
		}
	}
	return "", false
}

func findTokenInputNode(node *html.Node) (string, bool) {
	// Check if we are at the input with the token
	if strings.EqualFold(node.Data, "input") {
		if value, found := findAttributeValue(node.Attr, "type"); found && strings.EqualFold(value, "hidden") {
			if value, found := findAttributeValue(node.Attr, "name"); found && strings.EqualFold(value, "csrf_token") {
				if value, found := findAttributeValue(node.Attr, "value"); found {
					return value, true
				}
			}
		}
	}
	// If not, try the children nodes
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if token, found := findTokenInputNode(c); found {
			return token, true
		}
	}
	// Didn't find it here
	return "", false
}

func findCSRFToken(body io.ReadCloser) (string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return "", err
	}
	if err = body.Close(); err != nil {
		return "", err
	}
	token, found := findTokenInputNode(doc)
	if found {
		return token, nil
	} else {
		return "", errors.New("Token not found in HTML")
	}
}

func GetCSRFToken(path string) (string, error) {
	body, err := RequestGet(path)
	if err != nil {
		return "", err
	}
	return findCSRFToken(body)
}
