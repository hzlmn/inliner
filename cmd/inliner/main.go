package main

import (
  "log"
  "fmt"
  "os"
  "bytes",
  "bufio"
  "golang.org/x/net/html"
  "io/ioutil"
)


// error handling helper
func handleError(err error) {
  if err != nil {
    log.Fatal(err)
  }  
}

// file reading helper
func read(file string) (data []byte) {
  data, err := ioutil.ReadFile(file)
  handleError(err)
  return
}

func patch(old *html.Node, data []byte) {
  textNode := &html.Node{
    Type: html.TextNode,
    Data: string(data),
  }

  styleNode := &html.Node{
    Type: html.ElementNode,
    Data: "style",
    FirstChild: textNode,
  }

  old = styleNode
}

// Main entry point
func main() {
  var file string

  if len(os.Args) > 1 {
    file = os.Args[1]
  }

  data := read(file)

  tree, err := html.Parse(bytes.NewReader(data))

  handleError(err)

  traverser := &Traverser{
    visitors: VMap{
      html.TextNode: func (node *html.Node, parent *html.Node) {
        // pass
      },

      html.ElementNode: func (node *html.Node, parent *html.Node) {
        if node.Data != "link" {
          return
        }

        if node.Attr != nil {
          for _, attr := range node.Attr {
            if attr.Key == "href" {
              //ctx := read(attr.Val)
              patch(node,read(attr.Val))
              fmt.Println(string(read(attr.Val)))
              return
            }
          }
        } 
      },
    },
  }

  traverser.traverse(tree)

  
}
