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
  return strings.TrimSpace(string(data)) 
}

func convertToInt(arr []string) []int {
  ret := make([]int, len(arr))
  for i, l := range arr {
    n, _ := strconv.Atoi(l)
    ret[i] = n
  }

  sort.Slice(ret, func(i, j int) bool {
    return ret[i] < ret[j]
  })
  return ret
}

func nextJoltage(numbers []int) (int, int) {
  target := numbers[0]
  for i, n2 := range numbers[1:] {
    if n2 - target <= 3 {
      return n2, i
    }
  } 
  return 0, 0
}

func differences(numbers []int) {
  i1 := 0
  r1 := 0
  r3 := 0

  for {
    n, i2 := nextJoltage(numbers[i1:])
    if n == 0 { 
      break
    }
    if n - numbers[i1] == 3 { r3++ } 
    if n - numbers[i1] == 1 { r1++ } 
    i1 += i2 + 1
  } 
  fmt.Println(r1, r3)
}

func countArrangements(numbers []int) {
  mem := map[int]int{}
  var ways func(int) int
  ways = func (i1 int) int {
    if i1 >= len(numbers)-1 {
			return 1
		}
		if mem[i1] != 0 {
			return mem[i1]
		}
	  
    c := 0
    // start the loop from the passed in index
		for i2 := i1 + 1; i2 < len(numbers); i2++ {
      // for each next number that is within 3 increment the count and
      // then run this ways function again
      if numbers[i2]-numbers[i1] > 3 {
        break
			}
      n := ways(i2)
			c += n
      mem[i2] = n
    }
    return c 
	}
  fmt.Println(ways(0))
}


func main() {
  data := readInput()
  numbers := convertToInt(strings.Split(data, "\n"))
  numbers = append(numbers, numbers[len(numbers)-1] + 3)
  start := 0
  numbers = append([]int{start}, numbers...)
  // part 1 differences(numbers)
  countArrangements(numbers)
}
