package main

import (
	"fmt"
	"reflect"
)

type Container struct {
	s reflect.Value
}

func NewContainer(t reflect.Type, size int) *Container {
	if size <= 0 {
		size = 200
	}
	return &Container{
		s: reflect.MakeSlice(reflect.SliceOf(t), 0, size),
	}
}

func (c *Container) Put(val interface{}) error {
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		return fmt.Errorf("Put: cannot put a %T into a slice of %s", val, c.s.Type().Elem())
	}

	c.s = reflect.Append(c.s, reflect.ValueOf(val))
	return nil
}
