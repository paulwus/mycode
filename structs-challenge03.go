/* RZFeeser | Alta3 Research
   CHALLEGE 01 - Create a struct named Astro */

package main

import "fmt"

type astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

type nasaMission struct {
    people []astro
    number int
    message string
}


func main() {

    ast1 := astro{"Megan McArthur", 35, "ISS", false}
    ast2 := astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := astro{"Raja Chari", 39, "SpaceX Crew-3", true}

    p := []astro{ast1, ast2, ast3}

    s := nasaMission{p, 3, "success"}

    // print without field names
    fmt.Println(s)

    // print blank line
    fmt.Println("")

    // print with field names
    fmt.Printf("%+v", s)
}
