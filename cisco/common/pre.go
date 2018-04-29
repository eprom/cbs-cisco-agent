package common

import (
	"errors"
	"golang.org/x/net/html"
	"strings"
)

type TryParseFunc func(data []byte) (interface{}, error)

func parsePreText(node *html.Node, parser TryParseFunc) (interface{}, error) {
	// Check if we are at the input with parsable data
	if strings.EqualFold(node.Data, "pre") {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				if data, err := parser([]byte(c.Data)); err == nil {
					return data, nil
				}
			}
		}
	}
	// If not, try the children nodes
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if data, err := parsePreText(c, parser); err == nil {
			return data, nil
		}
	}
	// Didn't find it here
	return nil, errors.New("no parsable data found")
}

func GetData(path string, parser TryParseFunc) (interface{}, error) {
	// Post the command
	body, err := PostCommand(path)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	// Parse the HTML
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	// Find the right data
	return parsePreText(doc, parser)
}
