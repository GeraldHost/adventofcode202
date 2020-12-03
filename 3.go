package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
)

func readInput(n string) string {
  data, err := ioutil.ReadFile(fmt.Sprintf("%s.txt", n))
  if err != nil {
    log.Fatal("failed to read file")
  }
  return string(data) 
}

func Count(hsize, lsize int, lines []string) int {
  h := 0
  l := 0
  count := 0

  for h < len(lines) - 2 {
    h = h + hsize
    l = (l + lsize) % len(lines[1])

    s := string(lines[h][l])
    if s == "#" {
      count += 1
    }
  }

  return count
}

func main() {
  data := readInput("3")
  lines := strings.Split(data, "\n")

  c1 := Count(1, 1, lines)
  c2 := Count(1, 3, lines)
  c3 := Count(1, 5, lines)
  c4 := Count(1, 7, lines)
  c5 := Count(2, 1, lines)

  fmt.Println(c1 * c2 * c3 * c4 * c5)
}
