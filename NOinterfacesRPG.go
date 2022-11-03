/* Alta3 Research | rzfeeser@alta3.com
   Interfaces - Simple RPG example with NO interfaces */
   
package main

import (
    "fmt"
)

// Create Wizard struct
type Wizard struct{}

// Wizard Receiver Function (method) - This is how we want the Wizard to Defend()
func (w Wizard) Defend() string { 
    return "Expelliarmus" 
}

// Wizard Receiver Function (method) - A Wizard has the unique ability to cast a Forget() spell
func (w Wizard) Forget() string {
    return "Obliviate"
}

// Create Barbarian struct
type Barbarian struct{}

// Barbarian Receiver Function (method) - This is how we want the Barbarian to Defend()
func (b Barbarian) Defend() string { 
    return "Dodge" 
}

func main() {

    gandalf := Wizard{}
    fmt.Println("Wizard defends:", gandalf.Defend())
    
    // repeat the same pattern for the Barbarian
    conan := Barbarian{}
    fmt.Println("Barbarian defends:", conan.Defend())

    // continue the pattern for any other players


    fmt.Println("Wizard makes us all forget:", gandalf.Forget())

}

