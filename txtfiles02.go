/* RZFeeser | Alta3 Research
   Reading text files with bufio package - writing contents into a slice
   See documentation @ https://pkg.go.dev/bufio  */

package main
 
import (
    "bufio"
    "fmt"
    "log"
    "os"
)
 
func main() {
    
    // os.Open() is platform-independent function 
    file, err := os.Open("zelda.txt")  // for read access

    // error checking
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    
    // ensure the file is closed
    defer file.Close()
    
    // this snippet will reproduce the file on the screen
    scanner := bufio.NewScanner(file)

    // we could also write the data into a slice
    var txtlines []string  // create a slice of strings
    for scanner.Scan() {
            txtlines = append(txtlines, scanner.Text())  // each time through the loop build the slice
    }

    // now we have a slice that contains the contents of our file
    // we can loop across that file and reproduce it on the screen
    // via the following snippet
    for _, eachline := range txtlines {    // the _ would be the "line number" (starting at 0)
            fmt.Println(eachline)
    }

}
