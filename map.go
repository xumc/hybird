package hybrid

import (
	"reflect"
)

func MapInt64Int(s interface{}, key string, value string) map[int64]int {
	m := imap(s, key, value)
	ret := make(map[int64]int)

	keys := m.MapKeys()
	for i := 0; i < len(keys); i++ {
		ret[keys[i].Int()] = int(m.MapIndex(keys[i]).Int())
	}

	return ret
}

func imap(s interface{}, key string, value string) reflect.Value {
	// validate the first arg s
	if s == nil {
		panic("the first arg shouldn't be nil")
	}

	kind := reflect.TypeOf(s).Kind()

	if kind != reflect.Array && kind != reflect.Slice {
		panic("only array and slice are supported for mapping")
	}

	// find elem type
	elem := reflect.TypeOf(s).Elem()
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	// validate key and value
	keyExist := false
	valueExist := false
	var keyType reflect.Type
	var valueType reflect.Type

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if key == field.Name {
			keyExist = true
			keyType = field.Type
		}
		if value == field.Name {
			valueExist = true
			valueType = field.Type
		}
	}
	if !keyExist {
		panic("key doesn't exist in the struct")
	}
	if !valueExist {
		panic("value doesn't exist in the struct")
	}

	length := reflect.ValueOf(s).Len()
	mapType := reflect.MapOf(keyType, valueType)
	m := reflect.MakeMap(mapType)

	for i := 0; i < length; i++ {
		structVal := reflect.ValueOf(s).Index(i)
		if structVal.Kind() == reflect.Ptr {
			structVal = structVal.Elem()
		}

		keyVal := structVal.FieldByName(key)
		valueVal := structVal.FieldByName(value)

		m.SetMapIndex(keyVal, valueVal)
	}

	return m
}
