package main

import (
  "log"
  "fmt"
  "os"
  "bufio"
  "bytes"
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

func patch(data []byte) *html.Node {

  textNode := &html.Node{
    Type: html.TextNode,
    Data: string(data),
  }

  styleNode := &html.Node{
    Type: html.ElementNode,
    Data: "style",
    FirstChild: textNode,
  }

  return styleNode
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
              ctx := patch(read(attr.Val))
              parent.RemoveChild(node)
              //fmt.Println("new node", ctx.FirstChild)
              parent.AppendChild(ctx)
              return
            }
          }
        } 
      },
    },
  }

  traverser.traverse(tree)

  //fmt.Println("Child", tree.FirstChild.NextSibling.FirstChild.FirstChild.Type)
  var b bytes.Buffer
  w := bufio.NewWriter(&b)
  html.Render(w, tree)

  w.Flush()

  fmt.Println(b.String())
  
}
