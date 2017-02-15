// Exercise 1.3
// How to run:
// $ cd src/silvanolte.com/ch01/01_echo/exercises/1.3
// go test -bench=.

package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func main() {
  for i := 1; i < 6; i++ {
    fmt.Printf("Attempt %d:\n", i)
    Echo()
    EchoRange()
    EchoJoin()
    fmt.Println("")
  }
}

func Echo() {
  start := time.Now()
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Printf(s + " (%dns elapsed)\n", time.Since(start).Nanoseconds())
}

func EchoRange() {
  start := time.Now()
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Printf(s + " (%dns elapsed)\n", time.Since(start).Nanoseconds())
}

func EchoJoin() {
  start := time.Now()
  s := strings.Join(os.Args[1:], " ")
  fmt.Printf(s + " (%dns elapsed)\n", time.Since(start).Nanoseconds())
}