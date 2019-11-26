//
// date: 11/15/2019
// autor: tim siwula <@timxor>
// file: aileninvasion.go
//
// compile: go build
//
// run: ./aileninvasion N=5 cities.txt
//
// one liner: go build && ./aileninvasion ./cities.txt

package main

import (
  "fmt"
  "strings"
  "os"
  // "bufio"
  "log"
    "io/ioutil"
)

// array
//


// line = Foo​ ​north=Bar​ ​west=Baz​ ​south=Qu-ux
// key = line[0], directions = line[1]
// game.addcity(key, directions)
// game.updatecity()

// type Digraph struct {
//     v int
//     e int
// }
//
// func (k int, w int) addEdge() {
//     return map[]
// }
//
//
//
//
// // update - slice of aliens
// func update(aliens []Alien) Digraph {
//     return
// }
//
//
//
// // v(), e(), addEdge()
//
// type City struct {
//     vertex string
//     edge string
//     Residents string[]
//     Roads string[]
// }
//
// type Alien struct {
//     Name string
//     Start string
//     Life int // 10k
//     stocks := map[string]float64
// }
//
// type Game struct {
//     aliens Ailens
//     cities City
//     state int // 0, 1 or 2
//     aliens := map[Alien]boolean
//     cities := map[City]string
//
// }
//
// type TwoDArray struct {
//     a = [2][2]int
// }


// NewAlien()
// NewGame()
// NewCity()


// aliens["001"] = "location, currentMoves, ..."
// city["Foo"] = "north=Bar, couth=qu-ux, ..."



func main() {

  cli()
  readFileFromCli()
    //fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 3))
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
// func read() {
//
//     // file, err := os.Open("./cities.txt")
//     file, err := os.Open(os.Args[1])
//   if err != nil {
//         log.Fatal(err)
//     }
//     defer file.Close()
//
//     // var lines []string
//     scanner := bufio.NewScanner(file)
//
// 	// fmt.Println("\nread.forLoop:\n", result)
//
// 	fmt.Println("\nread.forLoop:\n")
//
// var count = 0
//
//     for scanner.Scan() {
//
//         // for each line
//         line := string(scanner.Text())
//         //fmt.Println(count, ": ", line, "\n")
//         fmt.Println(line)
//
//         count += 1
//
//         // split it
//         //Foo​ ​north=Bar​ ​west=Baz​ ​south=Qu-ux
//         fmt.Printf("%q\n", strings.SplitAfterN(line, "=", 3), "\n")
//
//
//
//
//     }
//
//     if err := scanner.Err(); err != nil {
//         log.Fatal(err)
//     }
// }

func readFileFromCli() {

    // assume the file is small enough to fit into memory
	fileName, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    // convert by wraping in string function
	data := string(fileName)
    lines := strings.Split(string(data), "\n")
	fmt.Println("Lines:\n\n", lines, "\n\nLine:\n")

    for start := 0; start < len(lines); start++ {

        sentence := lines[start]
        words := strings.Split(sentence, " ")
        if( sentence != ""){
            fmt.Println("sentence: ", sentence)
            fmt.Println("words: ", words)
            //fmt.Println("words[0]: ", words[0])
            cityName := words[0]
            fmt.Println("city: ", cityName)



            if strings.Contains(sentence, "north=") {
                // northCity := words[1]
                key := `north=`
                keyStart := strings.Index(sentence, key) + len(key)
                keyStop := keyStart + strings.Index(sentence[start:], ` `)
                fmt.Println("edge: ", "north")

                // fmt.Println("northCity: ", northCity)
                // fmt.Println("keyStart=",keyStart)
                // fmt.Println("keyStop=",keyStop)
                // northCity := words[1]

                //fmt.Println("len(northCity)=",len(northCity))



                vertex := sentence[keyStart:keyStop]

                fmt.Println("vertex: ", vertex)


                //fmt.Printf("%q\nnorthCity: ", strings.SplitAfterN(northCity, "=", 1))

                //fmt.Printf("%q\nnorth: ", strings.SplitAfterN(sentence, "=", 3), "\n")

            }

            if strings.Contains(sentence, "east") {

            }

            if strings.Contains(sentence, "south") {

            }

            if strings.Contains(sentence, "west") {

            }



 fmt.Println()
 fmt.Println()

        }

    }

    // i := strings.Index(lines, " ")
	// fmt.Println("Index: ", i)
    // cityName := lines[:i]
    // fmt.Println(cityName)

    //fmt.Printf("%q", strings.SplitAfterN(data, "=", 3))


}
