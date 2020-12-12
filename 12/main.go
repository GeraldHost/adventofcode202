package main

import (
	"fmt"
	"io/ioutil"
	"log"
  "strconv"
	"strings"
  "math"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

func parseDirection(a string, n, dx, dy int) (int, int) {
  if a == "N" {
    return 0, 1
  } else if a == "S" {
    return 0, -1
  } else if a == "E" {
    return 1, 0
  } else if a == "W" {
    return -1, 0
  } else if (a == "L" && n == 270) || (a == "R" && n == 90) {
    // right turn
    return dy, -dx    
  } else if (a == "R" && n == 270) || (a == "L" && n == 90) {
    // left turn
    return -dy, dx
  } else if (a == "R" && n == 180) || (a == "L" && n == 180) {
    // turn back
    return -dx, -dy
  } else {
    // Forward
    return dx, dy
  }
}

func turning(a string) bool {
  return a == "L" || a == "R"
}

func main() {
  data := readInput()
  lines := strings.Split(data, "\n")
  
  // D  X  Y
  // E  1  0
  // W -1  0
  // N  0  1
  // S  0 -1
  dx, dy := 1, 0
  x, y := 0, 0
  
  for _, l := range lines {
    action := l[:1]
    n, _ := strconv.Atoi(l[1:])
    pdx, pdy := parseDirection(action, n, dx, dy)
    if turning(action) {
      dx, dy = pdx, pdy
      continue
    } else {
      x = x + (n * pdx)
      y = y + (n * pdy)
    }
  }

  fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
