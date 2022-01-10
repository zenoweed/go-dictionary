package main

import (
	"fmt"
  "io/ioutil"
  "log"
)

func main() {
  readfile("dict.txt")
 
}
func readfile(filename string) {

  data, err := ioutil.ReadFile(filename)
  if err != nil {
        log.Panicf("failed reading data from file: %s", err)
  }

  fmt.Println("File Name:", filename)
  fmt.Println("Size: ", len(data), "bytes")
 }