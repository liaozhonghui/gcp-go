package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// My first example: struct property default values
	var a int = 1
	val := reflect.Indirect(reflect.ValueOf(&a))
	fmt.Println("val: ", val)
	fmt.Println("val kind:", val.Kind())
	val.Set(reflect.ValueOf(2))
	fmt.Println("after val set, a=", a)
	NewDS()

	// second: struct property default values
	rect := &shape{}
	myStruct(rect)

	// third: field.Tag.Lookup
	type S struct {
		F0 string `val:"123456"`
		F1 string `val:""`
		F2 string
	}
	s := S{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		if value, ok := field.Tag.Lookup("val"); ok {
			if value == "" {
				fmt.Println("(Empty.)")
			} else {
				fmt.Println(value)
			}
		} else {
			fmt.Println("(Non specific)")
		}
	}
}

type shape struct {
	Name  string `default:"shape"`
	Color string `default:"red"`
}

type DS struct {
	FieldOne string `default:"something"`
}

func NewDS() *DS {
	ds := &DS{}
	if err := initStruct(ds); err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("FieldONe= %s\n", ds.FieldOne)
	return ds
}

func initStruct(v interface{}) error {
	e := reflect.Indirect(reflect.ValueOf(v))
	fmt.Println("e kind is:", e.Kind())
	if e.Kind() != reflect.Struct {
		return errors.New("v must be struct")
	}
	fmt.Println("e type is:", e.Type())
	et, ev := e.Type(), e
	for i := 0; i < et.NumField(); i++ {
		field, val := et.Field(i), ev.Field(i)
		defaultValue, ok := field.Tag.Lookup("default")
		if !ok {
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			val.SetString(defaultValue)
		case reflect.Int:
			if x, err := strconv.ParseInt(defaultValue, 10, 64); err != nil {
				val.SetInt(x)
			}
		}
	}
	return nil
}

func myStruct(v interface{}) {
	e := reflect.Indirect(reflect.ValueOf(v))
	fmt.Println(v)
	fmt.Println("NumField:", e.Type().NumField())
	for i := 0; i < e.NumField(); i++ {
		field, val := e.Type().Field(i), e.Field(i)

		tagValue, ok := field.Tag.Lookup("default")
		if !ok {
			continue
		}
		val.SetString(tagValue)
		fmt.Printf("num: %d; field: %s\n", i, reflect.ValueOf(field))
		fmt.Printf("num: %d, tag-value: %s\n", i, tagValue)
		fmt.Printf("num: %d; val: %s\n", i, val)
	}
}
