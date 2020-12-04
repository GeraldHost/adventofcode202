package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strconv"
  "strings"
)

var requiredKeys = []string{
  "byr",
  "iyr",
  "eyr",
  "hgt",
  "hcl",
  "ecl",
  "pid"}

var validEcl = map[string]bool{
  "amb": true,
  "blu": true,
  "brn": true,
  "gry": true,
  "grn": true,
  "hzl": true,
  "oth": true}

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return string(data) 
}

func validateDate(v string, min, max int) bool {
  n, e := strconv.Atoi(v)
  if e != nil {
    return false
  }
  return len(v) == 4 && n >= min && n <= max
}

func validateHeight(v string) bool {
  unit := v[len(v)-2:] 
  if !(unit == "cm" || unit == "in") {
    return false
  }
  n, e := strconv.Atoi(v[:len(v)-2])
  if e != nil {
    return false
  }
  if unit == "cm" {
    return n >= 150 && n <= 193
  } else {
    return n >= 59 && n <= 76
  }
}

func validateHairColor(v string) bool {
  if v[0] != '#' {
    return false
  }
  validSet := map[rune]bool{'0': true,'1': true,'2': true,'3': true,'4': true,'5': true,'6': true,'7': true,'8': true,'9': true,'a': true,'b': true,'c': true,'d': true,'e': true,'f': true}
  for _, c := range v[1:] {
    if _, ok := validSet[c]; !ok {
      return false
    }
  }
  return true
}

func validateEcl(v string) bool {
  _, ok := validEcl[v]
  return ok
}

func validatePid(v string) bool {
  if len(v) != 9 {
    return false
  }
  for _, c := range v {
    if _, e := strconv.Atoi(string(c)); e != nil {
      return false
    }
  }
  return true
}

func validateValue(k, v string) bool {
    switch k {
      case "byr":
        return validateDate(v, 1920, 2002)
      case "iyr":
        return validateDate(v, 2010, 2020)
      case "eyr":
        return validateDate(v, 2020, 2030)
      case "hgt":
        return validateHeight(v)
      case "hcl":
        return validateHairColor(v)
      case "ecl":
        return validateEcl(v)
      case "pid":
        return validatePid(v)
      default:
        return true
    }
}

func isValid(m map[string]string) bool {
  for _, key := range requiredKeys {
    v, ok := m[key]
    if !ok {
      return false
    }
    if a := validateValue(key, v); !a {
      return false
    }

  }
  return true 
}

func main() {
  data := readInput()
  items := strings.Split(data, "\n\n")
  valid := 0
  count := 0

  for _, item := range items {
    s := strings.ReplaceAll(item, "\n", " ")
    p := strings.Split(s, " ")
    m := make(map[string]string)
    for _, a := range p {
      x := strings.Split(a, ":")
      k := x[0]
      v := x[1]
      m[k] = v
    } 
    if isValid(m) {
      valid += 1
    }
    count += 1
  }
  fmt.Println(valid, count, len(items))
}

