package main

import (
	"fmt"
)

//Animal is interface of hewan
type Animal interface {
	Talk() string
	Walking() string
}

//Dog struct is object of Dog
type Dog struct {
	legs int
}

//Talk func is how animal talk
func (d Dog) Talk() string {
	return "Bark"
}

//Walking func is how much animal legs
func (d Dog) Walking() string {
	d.legs = 4
	return fmt.Sprintf("Walking using %d legs", d.legs)
}

//Chicken struct is object of Dog
type Chicken struct {
	legs int
}

//Talk func is how animal talk
func (d Chicken) Talk() string {
	return "Petok"
}

//Walking func is how much animal legs
func (d Chicken) Walking() string {
	d.legs = 2
	return fmt.Sprintf("Walking using %d legs", d.legs)
}

//Walking func is how much animal legs
func (d Chicken) Eating() string {
	d.legs = 2
	return fmt.Sprintf("Walking using %d legs", d.legs)
}

func main() {
	var dog Animal
	dog = Dog{legs: 4}
	fmt.Println(dog.Talk(), "and", dog.Walking())
	fmt.Printf("%T %v", dog, dog)
	fmt.Println()

	var chicken Animal
	chicken = Chicken{legs: 4}
	fmt.Println(chicken.Talk(), "and", chicken.Walking())
	fmt.Printf("%T %v", chicken, chicken)
}
