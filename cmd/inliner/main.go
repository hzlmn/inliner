package main

import (
	"fmt"
  "flag"
	"log"
  "bytes"
  "bufio"
	"path/filepath"
  "golang.org/x/net/html"
)

var (
	input = flag.String("i", "", "define entry point to file that should be inlined")
	output = flag.String("o", "out.html", "output file path")
)


// Main entry point
func main() {
	flag.Parse()

	if len(*input) < 1 {
		log.Fatal("-i input flag should be defined")
	}

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

	// Shold be replaced with flag
  write(*output, b.Bytes())

	abspath, err := filepath.Abs(*output)
	handleError(err)

	fmt.Printf("File was written at: %s \n", abspath)
}
