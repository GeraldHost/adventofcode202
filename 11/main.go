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
  data, err := ioutil.ReadFile("test.txt")
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
  directions := []int{1,-1}

  // diagonals
  x1 := x
  y1 := y
  for _, d := range directions {
    for x1 < len(m[x]) && x1 >= 0 && y1 < len(m) && y1 >= 0 {
      if x1 == x && y1 == y {
        x1+=d
        y1+=d
        continue
      }
      if m[y1][x1] != "." {
        if m[y1][x1] == "#" {
          count++ 
        }
        break
      }
      x1+=d
      y1+=d
    }
  }

  fmt.Println("Count", count)

  // X
  x2 := x
  y2 := y
  for _, dx := range directions {
    for x2 < len(m) && x2 >= 0 {
      if x2 == x {
        x2+=dx
        continue
      }
      if m[y2][x2] != "." {
        if m[y2][x2] == "#" {
          count++
        }
        break
      }
      x2+=dx
    }
  }
  fmt.Println("Count", count)

  // Y
  x3 := x
  y3 := y
  for _, dy := range directions {
    for y3 < len(m[y]) && y3 >= 0 {
      if y3 == y {
        y3+=dy
        continue
      }
      if m[y3][x3] != "." {
        if m[y3][x3] == "#" {
          count++
        }
        break
      }
      y3+=dy
    }
  }
  fmt.Println("Count", count)

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
      fmt.Println(ca)
      if ca >= 5 && pos == "#" {
        change = true
        mtx[i][j] = "L"
      } else if ca <= 0 && pos == "L" {
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
  
  ans, _  := pass(m)
  //ans, _ = pass(ans)
  debug(ans)
  n := countInView(ans, 0, 3)
  fmt.Println(n)
  //if change { 
  //  for {
  //    ans, change = pass(ans)
  //    if !change {
  //      break;
  //    }
  //  }
  //}
  //debug(ans)
  //n := countTotalOcc(ans)
  //fmt.Println(n)
}
