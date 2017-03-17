package main

import (
	"log"
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

// write helper
func write(file string, data []byte) {
	err := ioutil.WriteFile(file, data, 0666)
	handleError(err)
}
