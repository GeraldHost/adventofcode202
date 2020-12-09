package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
  "sort"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return string(data) 
}

func validate(arr []string, v int) bool {
  found := false
  for _, s1 := range arr {
    n1, _ := strconv.Atoi(s1)
    for _, s2 := range arr {
      n2, _ := strconv.Atoi(s2)
      if n1 != n2 && n1 + n2 == v {
        found = true
      }
    }
  }
  return found
}

func main() {
  data := readInput()
  lines := strings.Split(strings.TrimSpace(data), "\n")

  lookback := 25

  for i := lookback; i < len(lines); i++ {
    tmp := make([]string, lookback)
    copy(tmp, lines[i-lookback:])
    n, _ := strconv.Atoi(lines[i])
    found := validate(tmp, n)
    if !found {
      // part 1: fmt.Printf("Invalid: %s, Index: %d", lines[i], i)
      tmp2 := make([]string, len(lines[:i]))
      copy(tmp2, lines[:i])
      
      for i := 0; i < len(tmp2); i++ {
        sum, _ := strconv.Atoi(tmp2[i])
        for j := i + 1; j < len(tmp2); j++ {
          n2, _ := strconv.Atoi(tmp2[j])
          sum += n2
          if sum == n {
            tmp3 := make([]string, len(tmp2[i:j]))
            copy(tmp3, tmp2[i:j])
            sort.Slice(tmp3, func(i, j int) bool {
              n1, _ := strconv.Atoi(tmp3[i])
              n2, _ := strconv.Atoi(tmp3[j])
              return n1 < n2
            })
            // part 2
            fmt.Println(tmp3[0], tmp3[len(tmp3)-1])
            break;
          }
        }
      }
    }
  }
}
