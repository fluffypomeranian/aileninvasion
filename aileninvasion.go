//
// date: 11/15/2019
// autor: tim siwula <@timxor>
// file: aileninvasion.go
//
// compile: go build
//
// run: ./aileninvasion N=5 cities.txt
//
// one liner: go build && ./aileninvasion cities.txt

package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
)

func main() {
  cli()
  read()
  //simulate()
}

// prints its command-line arguments.
func cli() {
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
}

// reads contents of a file
// line max length is +-65,000 characters
func read() {
    // file, err := os.Open("/path/to/file.txt")
    file, err := os.Open(os.Args[1])
  if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
