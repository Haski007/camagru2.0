package main

import (
	"html/template"
	"reflect"
)

var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

func RangeStructer(args ...interface{}) []interface{} {
    if len(args) == 0 {
        return nil
    }

    v := reflect.ValueOf(args[0])
    if v.Kind() != reflect.Struct {
        return nil
    }

    out := make([]interface{}, v.NumField())
    for i := 0; i < v.NumField(); i++ {
        out[i] = v.Field(i).Interface()
    }

    return out
}
