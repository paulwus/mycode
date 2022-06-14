/* RZFeeser | Alta3 Research
   switch - case and default  */

package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

    // init gover                
    var gover string = runtime.Version() 

    switch {
    case  strings.Contains(gover,"go1.18"):     
        fmt.Println("You are using the latest version of GoLang")
    case strings.Contains(gover, "go1.17"):
        fmt.Println("This version of Go is fine")
    default: 
        fmt.Println("Upgrade GoLang before you continue")
    }
}
