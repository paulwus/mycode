/* Alta3 Research | RZFeeser
   Using fmt.Printf - String formatting structs and PirateTreasure*/

package main

import (
    "fmt"
)


func main() {

    // create a "struct" called, 'p' of the type "PirateTreasure"
    const urlBase = "https://example.org:6001/v2/snacks?"
    var r string = "req=snickers"
    var q string = "quantity=1"
    var s string = "size=king"

    // create url 
    result := fmt.Sprintf("%s%s&%s&%s", urlBase, r, q, s)

    // print url 
    fmt.Println(result)

}
