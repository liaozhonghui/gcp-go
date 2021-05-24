package my_interface

import (
	"fmt"
	"math"
)

type Person struct {
	Name   string
	Sexual string
	Age    int
}

func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n", p.Name, p.Sexual, p.Age)
}
func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n", p.Name, p.Sexual, p.Age)
}

type shape interface {
	area() float64
	perimeter() float64
}

type Basic struct {
	style string
}
type rect struct {
	Basic
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}
func (r *rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Country struct {
	Name string
}

type City struct {
	Name string
}
type Stringable interface {
	ToString() string
}

func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

func (c *Country) ToString() string {
	return "Country= " + c.Name
}
func (c *City) ToString() string {
	return "City= " + c.Name
}

func main() {
	p := Person{
		Name:   "Liao xin",
		Sexual: "Male",
		Age:    25,
	}

	PrintPerson(&p)
	p.Print()

	r := rect{Basic: Basic{style: "background-color:red;"}, width: 2.9, height: 4.8}
	c := circle{radius: 4.3}
	fmt.Println("r.basic.style:", r.style)
	s := []shape{&r, &c}
	for _, sh := range s {
		fmt.Println(sh)
		fmt.Println(sh.area())
		fmt.Println(sh.perimeter())
	}
	d1 := Country{Name: "USA"}
	d2 := City{Name: "Los Angeles"}

	PrintStr(&d1)
	PrintStr(&d2)
}
