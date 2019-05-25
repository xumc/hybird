package hybrid

import (
	"reflect"
)

func ExtractInt64(s interface{}, key string) []int64 {
	values := iextract(s, key)
	int64slice := make([]int64, len(values))
	for i, v := range values {
		int64slice[i] = v.Int()
	}
	return int64slice
}

func ExtractInt(s interface{}, key string) []int {
	values := iextract(s, key)
	intslice := make([]int, len(values))
	for i, v := range values {
		intslice[i] = int(v.Int())
	}
	return intslice
}

func ExtractString(s interface{}, key string) []string {
	values := iextract(s, key)
	stringslice := make([]string, len(values))
	for i, v := range values {
		stringslice[i] = v.String()
	}
	return stringslice
}

func iextract(s interface{}, key string) []reflect.Value {
	// validate the first arg s
	if s == nil {
		panic("the first arg shouldn't be nil")
	}

	kind := reflect.TypeOf(s).Kind()

	if kind != reflect.Array && kind != reflect.Slice {
		panic("only array and slice are supported for mapping")
	}

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

	length := reflect.ValueOf(s).Len()
	values := make([]reflect.Value, length)

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
