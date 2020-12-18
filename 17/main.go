package main

import (
	"io/ioutil"
	"log"
  "fmt"
	"strconv"
	"strings"
)

func readInput() string {
  data, err := ioutil.ReadFile("sample.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

type Loc struct {
  x int
  y int
  z int
}

func main() {
  // fuck this, doing this in pythong because doing n dimentional matrix in go hurts my brain!
  // TODO: come back tomorrow and try this in Go now you understand the problem better dipshit - 18 dec 2020
}
