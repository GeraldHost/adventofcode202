package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"regexp"
	"strings"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

func process(str string) int {
  tokens := strings.Split(str, " ")
  sum := 0
  op := ""
  for _, token := range tokens {
    if token == "+" || token == "*" {
      op = token
    } else {
      n, _ := strconv.Atoi(token)
      switch op {
        case "+":
          sum += n
        case "*":
          sum *= n
        default:
          sum = n
      }
    }
  }
  return sum
}

func eval(str string) int {
  re := regexp.MustCompile(`\([^\(\)]+\)`)
  for re.MatchString(str) {
    str = re.ReplaceAllStringFunc(str, func(s string) string {
      n := eval(s[1:len(s)-1])
      return strconv.Itoa(n)
    })
  }

  // part 2
  re = regexp.MustCompile(`\d+ \+ \d+`)
	for re.MatchString(str) {
		str = re.ReplaceAllStringFunc(str, func(s string) string {
      p := process(s)
      return strconv.Itoa(p)
		})
	}

  return process(str)
}

func main() {
  data := readInput()
  lines := strings.Split(data, "\n")
  sum := 0
  for _, ln := range lines {
    v := eval(ln)
    sum+=v
  }
  fmt.Println(sum)
}
