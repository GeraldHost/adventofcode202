package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readInput() string {
  data, err := ioutil.ReadFile("sample.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

func zeros(y, x int) [][]string {
  ret := make([][]string, y)
  for i := 0; i < y; i++ {
    ret[i] = make([]string, x)
    for j := 0; j < x; j++ {
      ret[i][j] = "0"
    }
  }
  return ret
}

func debug(tile [][]string, name string) {
  fmt.Printf("--------------------- DEBUG: %s -------------------------\n", name)
  for _, t := range tile {
    fmt.Println(t)
  }
} 

func rot(tile [][]string) [][]string {
  cols := len(tile[0])
  rows := len(tile)

  nt := zeros(cols, rows)

  for i := 0; i < rows; i++ {
    for j := 0; j < cols; j++ {
      nt[j][i] = tile[rows-i-1][j]
    }
  }
  return nt
}

func flip(tile [][]string) [][]string {
  cols := len(tile[0])
  rows := len(tile)

  nt := zeros(cols, rows)

  for i := 0; i < rows; i++ {
    for j := cols-1; j >= 0; j-- {
      nt[i][j] = tile[i][cols-j-1]
    }
  }
  return nt
}

// returns an array of top, right, bottom, left
func edges(tile [][]string) [][]string {
  top := tile[0]
  bottom := tile[len(tile)-1]
  left := make([]string, len(tile))
  right := make([]string, len(tile))

  for i := 0; i < len(tile); i++ {
    left[i] = tile[i][0]
    right[i] = tile[i][len(tile[i])-1]
  }

  return [][]string{top,right,bottom,left}
}

func main() {
  data := readInput()
  tiles := strings.Split(data, "\n\n")
  m := make(map[string][][]string)

  for _, tile := range tiles {
    lns := strings.Split(tile, "\n")
    n := lns[0]
    m[n] = make([][]string, len(lns[1:]))
    for i, lyr := range lns[1:] {
      m[n][i] = make([]string, len(lyr))
      for j, pix := range strings.Split(lyr, "") {
        m[n][i][j] = pix
      }
    }
  }

  for k, v := range m {
    debug(v, k)
    a := edges(v)
    fmt.Println(a)
  }
}
