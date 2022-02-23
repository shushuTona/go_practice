package package_interface

import (
	"fmt"
)

type Creature interface {
	ToStringName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) ToStringName() string  {
	return person.Name
}

func (animal Animal) ToStringName() string  {
	return animal.Name
}

func PrintOut(creature Creature) {
	fmt.Println(creature.ToStringName())
}
