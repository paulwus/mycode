/* Alta3 Research | RZFeeser
   Template - Using a Map to build a template  */

package main


import (
    "html/template"
    "os"
    "log"
    "io"
    "strings" // we need to this trim the extension off our template
)

func parse(path string) {
    
    // template.ParseFiles() will READ a template file
    // t is now our template (yes, we could of just passed our template string to parse(), but then
    // we couldn't have seen how template.ParseFiles() is used)
    t, err := template.ParseFiles(path)
    
    // check to see if our template was read correctly
    if err != nil {
        log.Print(err)
        return
    }

    // get rid of ".template" on our path name
    path = strings.TrimRight(path, ".template")

    // create a file "just" named example.css (not example.css.template)
    f, err := os.Create(path)
    if err != nil {
        log.Println("create file: ", err)
        return
    }

    // A sample config
    config := map[string]string{
        "textColor":      "#abcdef",
        "linkColorHover": "#ffaacc",
    }

    // create our file using the fill in the blank info we have within our map
    err = t.Execute(f, config)
    if err != nil {
        log.Print("execute: ", err)
        return
    }
    f.Close()
}

func main() {

    // Create a file in which we can store our template string
    f, _ := os.Create("example.css.template")
    // write this data into our template
    f.Write([]byte("The text color is {{.textColor}} and the link color is {{.linkColorHover}}"))
    // close our file
    f.Close()
    
    // call our function, "parse()" and tell it to use the template we just made
    parse("example.css.template")
    

    // try opening the finished file we just created (we also created example.css.template)
    f, _ = os.Open("example.css")
    io.Copy(os.Stdout, f)
    f.Close()
}
