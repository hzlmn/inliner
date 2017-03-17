package main

import (
  "fmt"
  "flag"
  "bytes"
  "bufio"
  "golang.org/x/net/html"
)

var (
	input = flag.String("i", "", "define entry point to file that should be inlined")
	output = flag.String("o", "bundle.html", "output file path")
)


// Main entry point
func main() {
	flag.Parse()

  fmt.Println(*input, *output)

  data := read(*input)

  tree, err := html.Parse(bytes.NewReader(data))

  handleError(err)

  traverser := &Traverser{
    visitors: VMap{
      html.TextNode: textNodeVisitor,
      html.ElementNode: elementNodeVisitor,
		},
  }

	traverser.traverse(tree)

  var b bytes.Buffer
  w := bufio.NewWriter(&b)
  html.Render(w, tree)

  w.Flush()

  fmt.Println(b.String())
}
