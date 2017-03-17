package main

import (
	"fmt"
	"golang.org/x/net/html"
)

var cached html.Node

func textNodeVisitor(node *html.Node, parent *html.Node) {
		// pass
		// should something be added here?
}

func elementNodeVisitor(node *html.Node, parent *html.Node) {
	//var href string

	if node.Data != "link" {
		return
	}

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			//href = attr.Val
			break
		}
	}

	//patch := StyleNode(read(href))

	for c := node.NextSibling; c != nil; c = c.NextSibling {
		fmt.Println("next", c)
	}

	// fmt.Println(node.PrevSibling.Data)
	// *node, cached = StyleNode(read(href)), *node
	// node.NextSibling = cached.NextSibling
	// node.NextSibling.PrevSibling = node
	// node.Parent = cached.Parent
	// node.PrevSibling = cached.PrevSibling
}
