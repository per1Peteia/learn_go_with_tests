package main

import "reflect"

// a function walk(x interface{}, fn func(string)) which takes a struct x and calls fn for all strings fields found inside.
func walk(x interface{}, fn func(in string)) {
	val := getValue(x)
	numberOfFieldValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		numberOfFieldValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfFieldValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numberOfFieldValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
