/* RZFeeser | Alta3 Research
   HTTP Server with r.URL.Query()  */

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", handler)

    fmt.Printf("Starting server at port 8055\n")
    log.Fatal(http.ListenAndServe(":8055", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

    // declare and init a value for name to be "guest"
    name := "guest"

    // check to see if they actually provided a value
    // for the name parameter
    keys, ok := r.URL.Query()["name"]

    // if they passed a parameter called "name"
    if ok {
        name = keys[0]  // in the event they used "name" param
                        // multiple times, we just get the
                        // first instance
    }

    fmt.Fprintf(w, "Hello %s!\n", name)
}

