/* Alta3 Research | rzfeeser@alta3.com
   Interfaces - Simple RPG example with interfaces */
   
package main

import (
    "fmt"
    "reflect"
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



// make an interface
type Player interface {
    Defend() string
    //Attack() string
}



func main() {

    gandalf := Wizard{}
    // fmt.Println("Wizard defends:", gandalf.Defend())
    conan := Barbarian{}
    // fmt.Println("Barbarian defends:", conan.Defend())

    // Thanks to the inteface Player
    // we have a common type-- it isn't the structure-- but what they DO
    // this allows us to create a slice of things that have the qualities of a "Player"
    players := []Player{gandalf, conan}

    // We can loop over our slice
    // knowing all the types in the slice have the same capabilities (in this case, Defend())
    for _, a := range players {
        fmt.Println(reflect.TypeOf(a).Name(), "defends:", a.Defend())
    }

    fmt.Println("Wizard makes us all forget:", gandalf.Forget())

}

