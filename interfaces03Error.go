/* Alta3 Research | rzfeeser@alta3.com

   An error type in go is defined as follows:

   type error interface {
       Error() string
   }

*/

package main

import(
    "fmt"
)

type HardwareProblem struct {
    System string   // what system appears to be having this error
    Location string // what physical location is throwing this error
}

func (hp HardwareProblem) Error() string {
    return fmt.Sprintf("hardware error! %s, is being reported @ location, %s", hp.System, hp.Location)
}

type HTTPProblem struct {
    Code int
    Message string
}

func (httpp HTTPProblem) Error() string {
    return fmt.Sprintf("bad http response! http response was, %v - %s", httpp.Code, httpp.Message)
}


// hardware or HTTP, we can can use a common wrapper
// to deal with anything that "has" the interface error
func handleMyErrors(err error) {
    fmt.Println(err.Error())
}


func main(){

    // mimic a hardware problem
    hp := HardwareProblem{
        "battery backup not found",
        "datacenter west",
    }

    // mimic an HTTP problem
    httpp := HTTPProblem{
        401,
        "No Authorization",
    }

    // we can call a common function to clear or handle both errors
    handleMyErrors(httpp)
    handleMyErrors(hp)

}

