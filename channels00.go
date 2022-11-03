/* Alta3 Research | RZFeeser
   Channels - Create a basic channel     */

package main
  
import "fmt"
  
func main() {
  
    // Creating a channel
    // Using var keyword - initializes with 'nil' - cannot transport data with us
    var mychannel chan int
    fmt.Println("Value of the channel: ", mychannel)       // the value will be 'nil', which is not usable
    fmt.Printf("Type of the channel: %T ", mychannel)
  
    // Creating a channel using make() function
    mychannel1 := make(chan int)
    fmt.Println("\nValue of the channel1: ", mychannel1)   // returns a valid memory address
    fmt.Printf("Type of the channel1: %T ", mychannel1)
}

