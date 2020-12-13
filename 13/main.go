package main

import (
	"fmt"
	"io/ioutil"
	"log"
  "strconv"
	"strings"
  "math/big"
)

func readInput() string {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal("failed to read file")
  }
  return strings.TrimSpace(string(data)) 
}

var one = big.NewInt(1)

func crt(jj [][]*big.Int) *big.Int {
  p := new(big.Int).Set(one)
  for _, a := range jj {
    p.Mul(p, a[1])
  }
  var x, q, s, z big.Int
  for i, a := range jj {
    n1 := a[1]
    q.Div(p, n1)
    z.GCD(nil, &s, n1, &q)
    if z.Cmp(one) != 0 {
      fmt.Println("Fuck")
      return one 
    }
    x.Add(&x, s.Mul(jj[i][0], s.Mul(&s, &q)))
  }
  return x.Mod(&x, p)
}

func main() {
  data := readInput()
  lines := strings.Split(data, "\n")

  // part 1 t, _ := strconv.Atoi(lines[0])
  ids := strings.Split(lines[1], ",")

  // part 1
  // ret_id := t
  // ret_offset := t

  // for _, id := range ids {
  //   if id == "x" {
  //     continue
  //   }
  //   nid, _ := strconv.Atoi(id)
  //   a := float64(nid) * math.Ceil(float64(t) / float64(nid))
  //   b := int(a) - t
  //   if b < ret_offset {
  //     ret_id = nid
  //     ret_offset = b
  //   }
  //   fmt.Println(ret_id, ret_offset, ret_id * ret_offset)
  // }

  // part 2
  jj := make([][]*big.Int, 0)
  for i, id := range ids {
    if id == "x" {
      continue
    }
    nid, _ := strconv.Atoi(id)
    jj = append(jj, []*big.Int{big.NewInt(int64(nid-i)),big.NewInt(int64(nid))}) 
  }

  a := crt(jj)
  fmt.Println(a)

  //i := 100000000000000
  //for {
  //  fmt.Printf("\r%d", i)
  //  i += jj[0][0]
  //  nr := true
  //  for _, a := range jj {
  //    id := a[0]
  //    offset := a[1]
  //    if (i + offset) % id != 0 {
  //      nr = false
  //    } 
  //  }
  //  if nr {
  //    fmt.Printf("\n\n%d\n", i)
  //    break
  //  }
  //}
  fmt.Println("End");
}
