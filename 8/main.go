package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "log"
  "strconv"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return string(data) 
}

func Run(lines []string) (int, int) {
  global := 0
  currPos := 0
  pastPos := make(map[int]bool)
  for {
    if currPos == len(lines) {
      break
    }
    if _, ok := pastPos[currPos]; ok {
      break
    }
    p := strings.Split(lines[currPos], " ")
    op := p[0]
    n, _ := strconv.Atoi(p[1])
    
    pastPos[currPos] = true
    switch op {
      case "nop":
        // noop
        currPos++
      case "acc":
        global+=n
        currPos++
      case "jmp":
        currPos+=n
    }
  }
  return global, currPos
}

func main() {
  data := readInput()
  lines := strings.Split(strings.TrimSpace(data), "\n")
  // part 1 global, currPos := Run(lines)
  // part 1 fmt.Println(global, currPos)

  for i, line := range lines {
    p := strings.Split(line, " ")
    op := p[0]
    n := p[1]

    testlines := make([]string, len(lines))
    copy(testlines, lines)
  
    if op == "jmp" {
      testlines[i] = strings.Join([]string{"nop",n}, " ")
    } else if op == "nop" {
      testlines[i] = strings.Join([]string{"jmp",n}, " ")
    }

    g, count := Run(testlines)


    if count == len(lines) {
      fmt.Println(g)
      break
    }
 
  }
}
