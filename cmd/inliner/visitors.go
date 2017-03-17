package main

import (
	"fmt"
	"golang.org/x/net/html"
)

var Cached *html.Node

func textNodeVisitor(node *html.Node, parent *html.Node) {
		// pass
}

func elementNodeVisitor(node *html.Node, parent *html.Node) {
	if node.Data == "link" && node.Attr != nil {
		fmt.Println(node.Attr)
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				parent.RemoveChild(node)
				//fmt.Println(node.Data, attr.Key, attr.Val)
			}
		}
	}
	// if node.Attr != nil {
	// 	for _, attr := range node.Attr {
	// 		if attr.Key == "href" {
	// 			parent.RemoveChild(node)

	// 			if Cached != nil {
	// 				fmt.Println("cache not null")
	// 				Cached.AppendChild(TextNode(read(attr.Val)))
	// 				return
	// 			}

	// 			Cached = StyleNode(read(attr.Val))
	// 			parent.AppendChild(Cached)

	// 			return
	// 		}
	// 	}
}
