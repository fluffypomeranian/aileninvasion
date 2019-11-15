// date: 11/15/2019
// autor: tim siwula <@timcsiwula>
// file: aileninvasion.go
// compile: go build
// run: ./aileninvasion
package main

import (
  "fmt"
  "os"
)

func main() {
  cli()
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
