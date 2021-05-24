package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Rectangle struct {
	Name   string
	Unit   string
	Length float64
	Width  float64
}

func convert(rect *Rectangle) (res map[string]string, err error) {
	e := reflect.Indirect(reflect.ValueOf(rect))
	if e.Kind() != reflect.Struct {
		return nil, errors.New("v must be struct.")
	}

	et, ev := e.Type(), e
	mapStringType := reflect.TypeOf(make(map[string]string))
	mapReflect := reflect.MakeMap(mapStringType)
	for i := 0; i < et.NumField(); i++ {
		field, val := et.Field(i), ev.Field(i)

		switch field.Type.Kind() {
		case reflect.String:
			mapReflect.SetMapIndex(reflect.ValueOf(field.Name), reflect.ValueOf(val.String()))
		case reflect.Float64:
			s := strconv.FormatFloat(val.Float(), 'f', 2, 64)
			mapReflect.SetMapIndex(reflect.ValueOf(field.Name), reflect.ValueOf(s))
		}
	}

	return mapReflect.Interface().(map[string]string), nil
}

func main() {
	res, _ := convert(&Rectangle{
		Name:   "rec-1",
		Unit:   "cm",
		Length: 12.121764,
		Width:  5.989781,
	})

	fmt.Printf("res = %+v", res)
}
