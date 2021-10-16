package structs

import (
	"reflect"
)

// Struct converts a map[string]interface{} into a concrete struct
// This is an anti-process of structs.Map, which converts a struct into a map[string]interface{}
// Thus, a struct typed field in the target struct should be matched with a map[string]interface{} typed value in original map
// If a field name claimed in the map is not found in given struct type (by matching the exact name or customized name defined in field tag), it will be discarded
//
// @param data	Data to be converted
// @param t		Struct type
// @return The target struct
func Struct(data map[string]interface{}, t reflect.Type) interface{} {
	ptr := reflect.New(t)
	value := ptr.Elem()
	for k, v := range data {
		if fieldValue := value.FieldByName(k); fieldValue.IsValid() {
			field, _ := t.FieldByName(k)
			setValue(fieldValue, field, v)
		} else {
			if fieldValue, field := findByTag(value, t, k); fieldValue.IsValid() {
				setValue(fieldValue, field, v)
			}
		}
	}
	return ptr.Elem().Interface()
}

func findByTag(v reflect.Value, t reflect.Type, name string) (reflect.Value, reflect.StructField) {
	for i := 0; i < t.NumField(); i++ {
		if GenerateFieldName(t.Field(i)) == name {
			return v.Field(i), t.Field(i)
		}
	}
	return reflect.Value{}, reflect.StructField{}
}

func setValue(fieldValue reflect.Value, field reflect.StructField, value interface{}) {
	switch fieldValue.Kind() {
	case reflect.Int:
		fieldValue.SetInt(int64(value.(int)))
	case reflect.Int8:
		fieldValue.SetInt(int64(value.(int8)))
	case reflect.Int16:
		fieldValue.SetInt(int64(value.(int16)))
	case reflect.Int32:
		fieldValue.SetInt(int64(value.(int32)))
	case reflect.Int64:
		fieldValue.SetInt(value.(int64))
	case reflect.Uint:
		fieldValue.SetUint(uint64(value.(uint)))
	case reflect.Uint8:
		fieldValue.SetUint(uint64(value.(uint8)))
	case reflect.Uint16:
		fieldValue.SetUint(uint64(value.(uint16)))
	case reflect.Uint32:
		fieldValue.SetUint(uint64(value.(uint32)))
	case reflect.Uint64:
		fieldValue.SetUint(value.(uint64))
	case reflect.Float32:
		fieldValue.SetFloat(float64(value.(float32)))
	case reflect.Float64:
		fieldValue.SetFloat(value.(float64))
	case reflect.Complex64:
		fieldValue.SetComplex(complex128(value.(complex64)))
	case reflect.Complex128:
		fieldValue.SetComplex(value.(complex128))
	case reflect.Bool:
		fieldValue.SetBool(value.(bool))
	case reflect.String:
		fieldValue.SetString(value.(string))
	case reflect.Array, reflect.Slice:
		setArray(fieldValue, value)
	case reflect.Map:
		setMap(fieldValue, value)
	case reflect.Ptr:
		handlePtr(fieldValue, field, value)
	case reflect.Uintptr:
		fieldValue.Set(reflect.ValueOf(value))
	case reflect.Struct:
		fieldValue.Set(reflect.ValueOf(Struct(value.(map[string]interface{}), field.Type)))
	default:
		return
	}
}

func setArray(fieldValue reflect.Value, value interface{}) {
	t := fieldValue.Type()
	v := reflect.ValueOf(value)
	slice := reflect.MakeSlice(t, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		slice = reflect.Append(slice, v.Index(i))
	}
	fieldValue.Set(slice)
}

func setMap(fieldValue reflect.Value, value interface{}) {
	t := fieldValue.Type()
	v := reflect.ValueOf(value)
	m := reflect.MakeMap(t)
	for _, k := range v.MapKeys() {
		m.SetMapIndex(k, v.MapIndex(k))
	}
	fieldValue.Set(m)
}

func handlePtr(fieldValue reflect.Value, field reflect.StructField, value interface{}) {
	v := reflect.New(fieldValue.Type().Elem())
	setValue(v.Elem(), field, value)
	fieldValue.Set(v)
}
