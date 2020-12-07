package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "regexp"
  "strings"
  "strconv"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return string(data) 
}

func Count(containedIn map[string][]string, target string) map[string]bool {
  set := make(map[string]bool)
  for _, name := range containedIn[target] {
    set[name] = true
    for k, v := range Count(containedIn, name) {
      set[k] = v
    }
  }
  return set
}

func Cost(contains map[string][]map[string]int, target string) int {
  count := 0
  for _, m := range contains[target] {
    for name, cost := range m {
      count += cost
      count += (Cost(contains, name) * cost)
    }
  }
  return count
}

func main() {
  r1 := regexp.MustCompile("(.+) bags contain")
  r2 := regexp.MustCompile("(?P<N>[0-9]+) (?P<C>.+?) bags?[,.]")

  data := readInput()
  data = strings.TrimSpace(data)
  
  lines := strings.Split(data, "\n")

              // parent, child
  contains := make(map[string][]map[string]int)
                 // child, parent
  containedIn := make(map[string][]string)

  for _, line := range lines {
    parent := r1.FindStringSubmatch(line)[1]
    childMatches := r2.FindAllStringSubmatch(line, 10)
    for _, matches := range childMatches {
      color := matches[2]
      number, _ := strconv.Atoi(matches[1])
      fmt.Println("n", number)
      contains[parent] = append(contains[parent], map[string]int{ color: number })
      containedIn[color] = append(containedIn[color], parent)
    }
  } 

  fmt.Println(len(Count(containedIn, "shiny gold")))
  fmt.Println(Cost(contains, "shiny gold"))
}
