package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

func main() {
  data := readInput()
  numbers := strings.Split(data, ",")

  mem := make(map[int][]int)

  si := 1
  ls := -1

  for _, n := range numbers {
    nn, _ := strconv.Atoi(n)
    mem[nn] = append(mem[nn], si)
    ls = nn
    si++
  }

  for i := si; i <= 30000000; i++ {
    if v, ok := mem[ls]; ok && len(v) > 1 {
      ls = v[len(v)-1] - v[len(v)-2]
    } else {
      ls = 0
    }
    mem[ls] = append(mem[ls], si)
    si++
  }

  fmt.Println(ls)
}
