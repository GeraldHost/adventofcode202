package main

import (
  "fmt"
  "log"
  "strings"
  "io/ioutil"
  "strconv"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return string(data) 
}

func isValid(pwd, letter string, min, max int) bool {
  count := 0
  for _, l := range pwd {
    if string(l) == letter {
      count += 1
    }
  }
  return count >= min && count <= max
}

func isValidAlt(pwd, letter string, pos1, pos2 int) bool {
  b1 := string(pwd[pos1 - 1]) == letter
  b2 := pos2 <= len(pwd) && string(pwd[pos2 - 1]) == letter
  if b1 && b2 {
    return false
  }
  return b1 || b2
}

func main() {
  valid := 0
  data := readInput()
  for _, line := range strings.Split(data, "\n") {
    parts := strings.Split(line, " ")
    if len(parts) == 3 {
      r, l, p := parts[0], string(parts[1][0]), parts[2]
      parts = strings.Split(r, "-")
      min, _ := strconv.Atoi(parts[0])
      max, _ := strconv.Atoi(parts[1])
      if isValidAlt(p, l, min, max) {
        valid += 1
      }
    }
  }
  fmt.Println(valid)
}
