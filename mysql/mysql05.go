/* Alta3 Research | RZFeeser
   mySQL - Return the "RowsAffected" by an update, insert, or delete */   

package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {

    db, err := sql.Open("mysql", "student:playstation5@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    // here we delete this data from the table
    sql := "DELETE FROM cities WHERE id IN (2, 4, 6)"
    res, err := db.Exec(sql)

    if err != nil {
        panic(err.Error())
    }

    // return the rows effected, and the changes
    affectedRows, err := res.RowsAffected()

    if err != nil {
        log.Fatal(err)
    }

    // formatter to display the changes
    fmt.Printf("The statement affected %d rows\n", affectedRows)
}
