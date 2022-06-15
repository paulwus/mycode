/* RZFeeser | Alta3 Research
   Reading text files with bufio package - recreating the contents on the screen
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
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        //fmt.Println("*** *** ***")   // this line would place *** *** *** between each line
    }

}
