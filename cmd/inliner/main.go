package main

import (
  "log"
  "fmt"
  "os"
  "bytes"
  "golang.org/x/net/html"
  "io/ioutil"
)

type traverser func (*html.Node) *html.Node

func walk(node *html.Node, cb traverser) {
  // pass
  return
}

func main() {
  var file string
  if len(os.Args) > 1 {
    file = os.Args[1]
  }

  data, err := ioutil.ReadFile(file)
  if err != nil {
    log.Fatal(err)
  }

  tree, err := html.Parse(bytes.NewReader(data))

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(tree.FirstChild.FirstChild.NextSibling.FirstChild.Attr[0].Val)
}
