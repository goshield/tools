package tools

import (
	"errors"
	"reflect"
)

// Clone returns a pointer which is a copied of input type
func Clone(t reflect.Type) (interface{}, error) {
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.New(t.Elem()).Interface(), nil
	case reflect.Struct:
		return reflect.New(t).Interface(), nil
	default:
		return nil, errors.New("clone invalid type. Support types are Ptr, Struct")
	}
}

// StructOf returns type of struct
func StructOf(v interface{}) (reflect.Type, error) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Ptr:
		return t.Elem(), nil
	case reflect.Struct:
		return t, nil
	default:
		return nil, errors.New("calls StructOf with invalid type. Support: Ptr, Struct")
	}
}
