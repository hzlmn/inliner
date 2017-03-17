package main

import (
	"golang.org/x/net/html"
)

func TextNode(data []byte) *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: string(data),
	}
}

func StyleNode(data []byte) html.Node {
	return html.Node{
    Type: html.ElementNode,
    Data: "style",
    FirstChild: TextNode(data),
  }
}
