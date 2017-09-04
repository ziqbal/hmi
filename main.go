
package main

import (

    "os"
    "log"
    "fmt"
    "time"
    "path/filepath"
    "github.com/boltdb/bolt"

)

func main( ) {

    fmt.Println( "HMI r0.1.YYYYMMDDHHSS (C)2017 Zafar Iqbal <ultrasine@gmail.com> " ) //_REV_

    //////////////////////////////////////////////////////////

    ex, err := os.Executable( )

    if( err != nil ) {
    
        panic( err )

    }

    exPath := filepath.Dir( ex )

    //////////////////////////////////////////////////////////

    db, err := bolt.Open( exPath + "/hmi.db" , 0600 , &bolt.Options{ Timeout : 1 * time.Second } )

    if( err != nil ) {

        log.Fatal( "Database Opening Error! Exiting..." )
        log.Fatal( err )

    }

    defer db.Close( )

    //////////////////////////////////////////////////////////

    time.Sleep( 100 * time.Millisecond )

    //////////////////////////////////////////////////////////

    fmt.Println( "FINI" )

}