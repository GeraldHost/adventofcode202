package main

import (
	"fmt"
	"io/ioutil"
	"log"
  "strconv"
	"strings"
  "regexp"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

func bitMask(str string) (int64, int64) {
  bit_mask := ""
  bit_replace := ""
  for _, c := range str {
    if c == 'X' {
      bit_mask += "1"
      bit_replace += "0"
    } else {
      bit_mask += "0"
      bit_replace += string(c)
    }
  }

  bit_mask_b, _ := strconv.ParseInt(bit_mask, 2, 64)
  bit_replace_b, _ := strconv.ParseInt(bit_replace, 2, 64)
 
  return bit_mask_b, bit_replace_b
}

func combo(mask string, chnl chan string) {
  c1 := make(chan string)
  defer close(chnl)

  if len(mask) <= 0 {
    chnl <- ""
    return
  }

  go combo(mask[1:], c1)
  for m := range c1 {
    if mask[0] == '0' {
      chnl <- "X" + m
    } else if mask[0] == '1' {
      chnl <- "1" + m
    } else if mask[0] == 'X' {
      chnl <- "0" + m
      chnl <- "1" + m
    }
  }
  return
}

func floatingBitMask(str string) []string {
  chnl := make(chan string)
  go combo(str, chnl)

  ret := make([]string, 0)
  for a := range chnl {
    ret = append(ret, a)
  }
  return ret
}

func main() {
  data := readInput()
  lines := strings.Split(data, "\n")
  
  mem := make(map[int64]int64)
  var mask, replace int64
  var masks []string

  for _, m := range lines {

    if m[0] == 'm' && m[1] == 'a' {
      x := strings.ReplaceAll(m, "mask = ", "")
      masks = floatingBitMask(x)
      fmt.Println("masks",masks)
      //mask, replace = bitMask(m)
      continue
    }
    
    r1, _ := regexp.Compile("mem\\[([0-9]+)\\]")
    r2, _ := regexp.Compile("mem\\[[0-9]+\\] = ([0-9]+)")
    m1 := r1.FindStringSubmatch(m)[1]
    m2 := r2.FindStringSubmatch(m)[1]
  
    v1, _ := strconv.Atoi(m1)
    v2, _ := strconv.Atoi(m2)
    
    for _, str := range masks {
      mask, replace = bitMask(str)
      ma := int64(v1)
      ma = ma & mask
      ma = ma | replace
      mem[ma] = int64(v2)
    }
  
  }
  
  sum := int64(0)
  for _, v := range mem {
    sum += v
  }

  fmt.Println(sum)
}
