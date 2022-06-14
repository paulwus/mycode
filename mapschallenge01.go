/* Alta3 Research | RZFeeser
   Map - hosts to IP:port resolution  */

package main

import (  
    "fmt"
)

func main() {  
    
    // create a map type, "langExtensions"
    langExtensions := map[string]string{
        "Python": ".py",
        "Java": ".java",
        "C++": ".cpp",
    }
    
    // name to lookup
    fmt.Println(langExtensions)
    
   // delete C++    
    delete(langExtensions, "C++")

    fmt.Println(langExtensions)
    
   // add julia
   langExtensions["Julia"] = ".jl"
   
   fmt.Println(langExtensions)

}
