package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func readInput() string {
  data, err := ioutil.ReadFile("test2.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

type Rule struct {
  min int
  max int
}

func isInvalid(a [][]Rule, n int) (bool, []int) {
  invalid := true
  is := make([]int, 0)
  for _, j := range a {
    for i, r := range j {
      is = append(is, i)
      if n <= r.max && n >= r.min {
        invalid = false
      }
    }
  }
  return invalid, is
}

//func getRules(rulesData string) []Rule {
//  rules := make([]Rule, 0)
//  for _, ln := range strings.Split(rulesData, "\n") {
//    reg := regexp.MustCompile("(\\d+)-(\\d+)")
//    res := reg.FindAllStringSubmatch(ln, -1)
//
//    min1, _ := strconv.Atoi(res[0][1])
//    max1, _ := strconv.Atoi(res[0][2])
//    
//    min2, _ := strconv.Atoi(res[1][1])
//    max2, _ := strconv.Atoi(res[1][2])
//
//    rules = append(rules, Rule { min1, max1 })
//    rules = append(rules, Rule { min2, max2 })
//  }
//  return rules
//}

func getRulePairs(rulesData string) [][]Rule {
  pairs := make([][]Rule, 0)
  for _, ln := range strings.Split(rulesData, "\n") {
    reg := regexp.MustCompile("(\\d+)-(\\d+)")
    res := reg.FindAllStringSubmatch(ln, -1)
    
    rules := make([]Rule, 0)
    min1, _ := strconv.Atoi(res[0][1])
    max1, _ := strconv.Atoi(res[0][2])
    min2, _ := strconv.Atoi(res[1][1])
    max2, _ := strconv.Atoi(res[1][2])
    rules = append(rules, Rule { min1, max1 })
    rules = append(rules, Rule { min2, max2 })

    pairs = append(pairs, rules)
  }
  return pairs
}

//func partOne(nearbyData string, rules []Rule) {
//  sum := 0
//  for _, ln := range strings.Split(nearbyData, "\n")[1:] {
//    for _, ns := range strings.Split(ln, ",") {
//      n, _ := strconv.Atoi(ns)
//      if isInvalid(rules, n) {
//        sum += n
//      }
//    }
//  }
//
//  fmt.Println(sum)
//}

func partTwo(nearbyData string, rules [][]Rule) {
  valid := make([]string, 0)
  numbers := make(map[int][]int, 0)

  for _, ln := range strings.Split(nearbyData, "\n")[1:] {
    invalid := false
    numbers1 := make(map[int]int, 0)
    for i, ns := range strings.Split(ln, ",") {
      n, _ := strconv.Atoi(ns)
      numbers1[i] = n
      if isInvalid(rules, n) {
        invalid = true
      }
    } 
    if !invalid {
      valid = append(valid, ln)
    }
  }

  fmt.Println(numbers)
}

func main() {
  data := readInput()
  parts := strings.Split(data, "\n\n")
  
  rulesData := parts[0]
  //yourData := parts[1]
  nearbyData := parts[2]
  
  rules := getRulePairs(rulesData)

  partTwo(nearbyData, rules)
}
