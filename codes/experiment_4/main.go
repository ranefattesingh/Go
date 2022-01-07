package main

import (
	"fmt"
	"sort"
)

type IGeneral interface {
	GetLegCount() int
	GetHandCount() int
	GetName() string
}

type Person struct {
	Id    int
	Hands int
	Legs  int
	Name  string
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetLegCount() int {
	return p.Legs
}

func (p *Person) GetHandCount() int {
	return p.Hands
}

type Alien struct {
	Id    int
	Hands int
	Legs  int
	Name  string
}

func (p *Alien) GetName() string {
	return p.Name
}

func (p *Alien) GetLegCount() int {
	return p.Legs
}

func (p *Alien) GetHandCount() int {
	return p.Hands
}

type NonHuman struct {
	Id    int
	Hands int
	Legs  int
	Name  string
}

func (p *NonHuman) GetName() string {
	return p.Name
}

func (p *NonHuman) GetLegCount() int {
	return p.Legs
}

func (p *NonHuman) GetHandCount() int {
	return p.Hands
}

func Printer(species IGeneral) {
	fmt.Printf("%s is having %d hands and %d legs.\n", species.GetName(), species.GetHandCount(), species.GetLegCount())
}

// SECTION 2

type PeopleSlice []Person

func (x PeopleSlice) Len() int           { return len(x) }
func (x PeopleSlice) Less(i, j int) bool { return x[i].Id < x[j].Id }
func (x PeopleSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x PeopleSlice) SortPeopleById() PeopleSlice {
	sort.Sort(x)
	return x
}

func main() {
	p := Person{
		Id:    1,
		Hands: 2,
		Legs:  2,
		Name:  "Sahil",
	}

	a := Alien{
		Id:    1,
		Hands: 4,
		Legs:  2,
		Name:  "Four Arms",
	}

	nH := NonHuman{
		Id:    1,
		Hands: 0,
		Legs:  4,
		Name:  "Dog",
	}

	Printer(&p)
	Printer(&a)
	Printer(&nH)

	// SECTION 2

	p1 := Person{
		Id:    5,
		Hands: 2,
		Legs:  2,
		Name:  "Sahil",
	}

	p2 := Person{
		Id:    3,
		Hands: 2,
		Legs:  2,
		Name:  "Rohan",
	}

	p3 := Person{
		Id:    1,
		Hands: 2,
		Legs:  2,
		Name:  "Laximan",
	}

	p4 := Person{
		Id:    2,
		Hands: 2,
		Legs:  2,
		Name:  "Sanil",
	}

	p5 := Person{
		Id:    4,
		Hands: 2,
		Legs:  2,
		Name:  "Harish",
	}

	var people = PeopleSlice{p1, p2, p3, p4, p5}

	var x sort.Interface = people

	fmt.Println("Number of people are:", x.Len())
	sortedPeople := people.SortPeopleById()
	fmt.Println("Id\tName")
	for i := 0; i < sortedPeople.Len(); i++ {
		fmt.Printf("%d\t%s\n", sortedPeople[i].Id, sortedPeople[i].Name)
	}

}
