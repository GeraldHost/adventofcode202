package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "log"
  "sort"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return string(data) 
}

func makeRange(max int) []int {
    a := make([]int, max+1)
    for i := range a {
      a[i] = i
    }
    return a
}

func Upper(arr []int) []int {
  return arr[len(arr)/2:]
}

func Lower(arr []int) []int {
  return arr[:len(arr)/2]
}

func SeatLoc(line string) ([]int, []int) {
  row := makeRange(127)
  col := makeRange(7)
  for _, c := range line {
    switch c {
      case 'F':
        row = Lower(row)
      case 'B':
        row = Upper(row)
      case 'L':
        col = Lower(col)
      case 'R':
        col = Upper(col)
    }
  }
  return row, col
}

func GetMySeat(arr []int) int {
  for i, n := range arr {
    if n != 0 && arr[i+1] != n + 1 {
      return n + 1
    }
  }
  return -1
}

func main() {
  data := readInput()
  lines := strings.Split(data, "\n")
  ids := make([]int, len(lines))
  for _, line := range lines {
    row, col := SeatLoc(line)
    rowNumber := row[0]
    colNumber := col[0]
    id := rowNumber * 8 + colNumber
    ids = append(ids, id)
  }
  sort.Ints(ids)
  fmt.Printf("Max Id is %d\n", ids[len(ids)-1])
  fmt.Printf("My seat ID is %d\n", GetMySeat(ids))
}
