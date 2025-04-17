package main
// this mean file belongs to main package. In Go, the main page i special
// It is the entry point for an executable program

import (
    "fmt"
)
//This imports the fmt package, which provides formatting functions for input and output. 
// It's commonly used for printing text to the console.

// Person struct defines the "class" data
type Person struct {
    Name string
    Age  int
}

// NewPerson is a constructor function
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

// SayHello is a method on the Person type
func (p *Person) SayHello() string {
    return fmt.Sprintf("Hello, my name is %s and I am %d years old", p.Name, p.Age)
}

// Birthday increases the person's age by 1
func (p *Person) Birthday() {
    p.Age++
}

func main() {
    // Create a new person using the constructor
    person := NewPerson("John", 30)
    
    // Call methods on the person
    fmt.Println(person.SayHello())
    
    // Update the person's age
    person.Birthday()
    
    // Access struct fields directly
    fmt.Printf("%s is now %d years old\n", person.Name, person.Age)
}