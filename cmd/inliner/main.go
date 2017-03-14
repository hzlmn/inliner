package main

import (
  "log"
  "fmt"
  "os"
  "bytes"
  "golang.org/x/net/html"
  "io/ioutil"
)

func Walk(node html.Node, cb func(html.Node)) {
  cb(node)
  for c := node.FirstChild; c != nil; c = c.NextSibling {
    Walk(*c, cb)
  }
}

func fetch(file string) ([]byte, error) {
  // loading filesting
  return make([]byte, 10), nil
}

func Traverse(node html.Node) {
  if node.Type == html.ElementNode && len(node.Attr) >= 1 {
    for _, atr := range node.Attr {
      //fmt.Println(atr)
      if atr.Key == "href" {
        content, _ := fetch(atr.Val)
        fmt.Println("content", string(content))
        fmt.Println("node has href", node.Data)
      }
    }
  }
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

  Walk(*tree, Traverse)
}
