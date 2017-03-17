package main

import (
	"golang.org/x/net/html"
)

var cached html.Node

func textNodeVisitor(node *html.Node, parent *html.Node) {
		// pass
		// should something be added here?
}

func elementNodeVisitor(node *html.Node, parent *html.Node) {
	var href string

	if node.Data != "link" {
		return
	}

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			href = attr.Val
			break
		}
	}

	*node, cached = StyleNode(read(href)), *node
	node.NextSibling = cached.NextSibling
	node.Parent = cached.Parent
	node.PrevSibling = cached.PrevSibling

	// Maybe wrap with new lines

	// raw := TextNode([]byte{10,32,32})
	// fmt.Println(raw)

	// for c := node.NextSibling; c != nil; c = c.NextSibling {
	// 	if c.Type == html.TextNode {
	// 		fmt.Println("test", []byte(c.Data))
	// 	}
	// }

}
