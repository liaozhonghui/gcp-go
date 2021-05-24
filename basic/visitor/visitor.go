package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Shape interface {
	accept(Visitor)
}
type Visitor func(shape Shape)

type Circle struct {
	Radius int
}
type Rectangle struct {
	Width, Heigh int
}

func (c Circle) accept(v Visitor) {
	v(c)
}
func (r Rectangle) accept(v Visitor) {
	v(r)
}

func JsonVistor(shape Shape) {
	bytes, err := json.MarshalIndent(shape, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shape) {
	bytes, err := xml.MarshalIndent(shape, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	c := Circle{10}
	r := Rectangle{100, 200}
	shapes := []Shape{c, r}
	for _, s := range shapes {
		s.accept(JsonVistor)
		s.accept(XmlVisitor)
	}
}
