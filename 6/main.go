package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "log"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return string(data) 
}

func Count(group string) int {
  m := make(map[string]int)
  s := strings.ReplaceAll(group, "\n", "")
  for _, letter := range strings.Split(s, "") {
    if _, ok := m[letter]; ok {
      m[letter] += 1
    } else {
      m[letter] = 1
    }
  }
  // part1 return len(m)
  fmt.Println(m)
  count := 0
  for _, v := range m {
    if v >= len(strings.Split(group, "\n")) {
      count++
    }
  }
  return count
}

func main() {
  data := readInput()
  groups := strings.Split(data, "\n\n")
  total := 0
  for _, group := range groups {
    count := Count(group)
    total += count
  }
  fmt.Println(total)
}
