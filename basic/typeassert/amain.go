package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 233
	p := reflect.ValueOf(&i)
	v := reflect.ValueOf(i)

	cp := p.Interface().(*int)
	cv := v.Interface().(int)

	fmt.Println(cp, cv)
}
