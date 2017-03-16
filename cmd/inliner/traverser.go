package main

import "golang.org/x/net/html"

// type definition for single visitor
type Visitor func(*html.Node, *html.Node)

// type definition for map of visitors
type VMap map[html.NodeType]Visitor

type Traverser struct {
  visitors VMap
}

func (t *Traverser) traverse(node *html.Node) {
  visitor := t.visitors[node.Type]

  if visitor != nil {
    visitor(node, node.Parent)
  }

  for c := node.FirstChild; c != nil; c = c.NextSibling {
    t.traverse(c)
  }
}