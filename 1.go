package main 

import (
  "fmt"
  "log"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {
  body, err := ioutil.ReadFile("1.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }

  numbers := strings.Split(string(body), "\n")
  
  var ret1 int 
  var ret2 int
  var ret3 int

  for _, n1 := range numbers {
    nn1, _ := strconv.Atoi(n1)
    for _, n2 := range numbers {
      nn2, _ := strconv.Atoi(n2)
      for _, n3 := range numbers {
        nn3, _ := strconv.Atoi(n3)
        if nn1 + nn2 + nn3 == 2020 {
          ret1, ret2, ret3 = nn1, nn2, nn3
          break
        }
      }
    }
  }

  fmt.Println(ret1 * ret2 * ret3)
}
