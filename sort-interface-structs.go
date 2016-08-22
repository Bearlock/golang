package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
}

type Allomancer struct {
	Name        *Person
	MistingType string
	Metals      string
}

type ByName []*Allomancer

func (allomancerSlice ByName) Len() int {
	return len(allomancerSlice)
}

func (allomancerSlice ByName) Swap(i, j int) {
	allomancerSlice[i], allomancerSlice[j] = allomancerSlice[j], allomancerSlice[i]
}

func (allomancerSlice ByName) Less(i, j int) bool {
	return allomancerSlice[i].Name.sayName() < allomancerSlice[j].Name.sayName()
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

	allomancers := []*Allomancer{&vin, &spook, &ham, &elend}

	fmt.Println("\nOriginal array:")
	for _, allomancer := range allomancers {
		fmt.Println(allomancer.stringify())
	}

	elend.updateTypeAndMetals("Mistborn", "All")
	fmt.Println("\n" + elend.stringify())

	fmt.Println("\nArray after updated metals and type for Elend:")
	for _, allomancer := range allomancers {
		fmt.Println(allomancer.stringify())
	}

	sort.Sort(ByName(allomancers))

	fmt.Println("\nArray after sort by name:")
	for _, allomancer := range allomancers {
		fmt.Println(allomancer.stringify())
	}

}
