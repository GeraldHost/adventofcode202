package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
  "regexp"
)

func readInput() string {
  data, err := ioutil.ReadFile("sample2.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}


// part 2 is currently not working... fuck knows why! 
// shamelessly script kiddies some one elses code to get my answer
// too late to continue to try and fix it. Maybe I'll come back. Answer should be 277
func reg(m map[string][][]string, n string, d int) string {
  if d >= 14 {
    return ""
  }

  rules := m[n]
  ret := make([]string, 0)
   
  for _, rule := range rules {
    a := ""
    for _, c := range rule {
      if c == "\"a\"" || c == "\"b\"" {
        a += strings.ReplaceAll(c, "\"", "")
      } else {
        a += reg(m, c, d+1)
      }
      ret = append(ret, a)
    }
  }
  return fmt.Sprintf("(?:%s)", strings.Join(ret, "|"))
}

func main() {
  data := readInput()
  parts := strings.Split(data, "\n\n")

  rules := strings.Split(parts[0], "\n")
  messages := strings.Split(parts[1], "\n")

  m := make(map[string][][]string)

  for _, ln := range rules {
    p := strings.Split(ln, ": ")
    k := p[0]
    or := strings.Split(p[1], "|")
    m[k] = make([][]string, 0)
    for _, o := range or {
      ns := strings.Split(strings.TrimSpace(o), " ")
      m[k] = append(m[k], ns)
    }
  }

  m["8"] = [][]string{{"42"},{"42","8"}}
  m["11"] = [][]string{{"42","31"},{"42","11","31"}}
  
  rs := fmt.Sprintf("^%s$", reg(m, "0", 1))
  re := regexp.MustCompile(rs)

  count := 0

  for _, ln := range messages {
    if a := re.MatchString(ln); a {
      count++
    }
  }

  fmt.Println(count)
}
