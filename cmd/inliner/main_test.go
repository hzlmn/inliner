package main

import (
	"testing"
)

// testing for node structures
// should returns corresponding one

func TestTextNode(t *testing.T) {
	str := "TEST STRING"
	testData := []byte(str)
	node := TextNode(testData)
	if node.Data != str {
		t.Log("Data should be equal")
		t.Fatal()
	}
}

func TestStyleNode(t *testing.T) {
	str := "some data"
	node := StyleNode([]byte(str))
	t.Log(node)
}
