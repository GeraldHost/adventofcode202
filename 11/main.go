package main

import (
	"fmt"
	"io/ioutil"
	"log"
	//"strconv"
	"strings"
  //"sort"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

func debug(m [][]string) {
  fmt.Println("-------- DEBUG --------")
  for _, l := range m {
    fmt.Println(l)
  }
}

func countAdj(m [][]string, x, y int) int {
  count := 0

  for i := y - 1; i <= y + 1; i++ {
    for j := x - 1; j <= x + 1; j++ {
      if i < 0 || i >= len(m) {
        continue
      }
      if j < 0 || j >= len(m[i]) {
        continue
      }
      if i == y && j == x {
        continue
      }
      if m[i][j] == "#" {
        count++
      }
    }
  }
  
  return count
}

func countInView(m [][]string, x, y int) int {
  count := 0
  pos := [][]int{
    {-1,-1},
    {-1,0},
    {-1,1},
    {0,-1},
    {0,1},
    {1,-1},
    {1,0},
    {1,1},
  }

  for _, p := range pos {
    y2, x2 := y+p[0], x+p[1]
    for 0<=y2 && y2<len(m) && 0<=x2 && x2<len(m[0]) {
      if m[y2][x2]=="#" {
        count++
        break
      } else if m[y2][x2] == "L" {
        break
      }
      y2+=p[0]
      x2+=p[1]
    }
  }

  return count
}

func countTotalOcc(m [][]string) int {
  count := 0 
  for _, a := range m {
    for _, c := range a {
      if c == "#" {
        count++
      }
    }
  }
  return count
}

func pass(m [][]string) ([][]string, bool) {
  change := false
  mtx := make([][]string, len(m))
  for i, seats := range m {
    mtx[i] = make([]string, len(seats))
    for j, pos := range seats {
      //ca := countAdj(m, j, i)
      ca := countInView(m, j, i)
      if ca >= 5 && pos == "#" {
        change = true
        mtx[i][j] = "L"
      } else if ca == 0 && pos == "L" {
        change = true
        mtx[i][j] = "#"
      } else {
        mtx[i][j] = pos
      }
    }
  }
  return mtx, change
}

func main() {
  data := readInput()
  rows := strings.Split(data, "\n")

  m := make([][]string, len(rows))

  for i, row := range rows {
    pos := strings.Split(row, "")
    m[i] = make([]string, len(pos))
    for j, col := range pos {
      m[i][j] = col
    }
  }
  
  // ans, _ := pass(m)
  // ans, _ = pass(ans)
  // debug(ans)

  ans, change := pass(m)
  if change { 
    for {
      ans, change = pass(ans)
      if !change {
        break;
      }
    }
  }
  debug(ans)
  n := countTotalOcc(ans)
  fmt.Println(n)
}
