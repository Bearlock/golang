package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Allomancer struct {
	Name        *Person
	MistingType string
	Metals      string
}

func NewAllomancer(name, mistingType, metals string) Allomancer {
	return Allomancer{
		Name:        &Person{name},
		MistingType: mistingType,
		Metals:      metals,
	}
}

func (p Person) sayName() string {
	return fmt.Sprintf("%s", p.Name)
}

func (mancer Allomancer) stringify() string {
	return fmt.Sprintf("%s is a %s who burns %s", mancer.Name.sayName(), mancer.MistingType, mancer.Metals)
}

func (mancer *Allomancer) updateTypeAndMetals(mistingType, metals string) {
	mancer.MistingType = mistingType
	mancer.Metals = metals
}

func main() {
	vin := NewAllomancer("Vin", "Mistborn", "All")
	spook := NewAllomancer("Spook", "Tineye", "Tin")
	ham := NewAllomancer("Hammond", "Pewterarm", "Pewter")
	elend := NewAllomancer("Elend", "Nobleman", "Nothing")

	fmt.Println(vin.stringify())
	fmt.Println(spook.stringify())
	fmt.Println(ham.stringify())
	fmt.Println(elend.stringify())

	elend.updateTypeAndMetals("Mistborn", "All")
	fmt.Println(elend.stringify())
}
