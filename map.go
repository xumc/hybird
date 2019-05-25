package hybrid

import (
	"reflect"
)

func MapInt64(s interface{}, key string) []int64 {
	values := imap(s, key)
	int64slice := make([]int64, len(values))
	for i, v := range values {
		int64slice[i] = v.Int()
	}
	return int64slice
}

func MapInt(s interface{}, key string) []int {
	values := imap(s, key)
	intslice := make([]int, len(values))
	for i, v := range values {
		intslice[i] = int(v.Int())
	}
	return intslice
}

func MapString(s interface{}, key string) []string {
	values := imap(s, key)
	stringslice := make([]string, len(values))
	for i, v := range values {
		stringslice[i] = v.String()
	}
	return stringslice
}

func imap(s interface{}, key string) []reflect.Value {
	// validate the first arg s
	if s == nil {
		panic("the first arg shouldn't be nil")
	}

	kind := reflect.TypeOf(s).Kind().String()

	if kind != "array" && kind != "slice" {
		panic("only array and slice are supported for mapping")
	}

	length := reflect.ValueOf(s).Len()
	values := make([]reflect.Value, length)

	// validate key
	keyExist := false
	elem := reflect.TypeOf(s).Elem()
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	for i := 0; i < elem.NumField(); i++ {
		if key == elem.Field(i).Name {
			keyExist = true
			break
		}
	}
	if !keyExist {
		panic("key doesn't exist in the struct")
	}

	// extract values of the key
	for i := 0; i < length; i++ {
		structVal := reflect.ValueOf(s).Index(i)
		if structVal.Kind() == reflect.Ptr {
			structVal = structVal.Elem()
		}
		values[i] = structVal.FieldByName(key)
	}

	return values
}
