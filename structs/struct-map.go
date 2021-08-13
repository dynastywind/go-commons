package structs

import (
	"reflect"
)

// Map maps a concrete struct to a map[string]interface{}
// This is an anti-process of structs.Struct, which converts a map[string]interface{} into a concrete struct
// Field will be converted to its exact name or specified in field tag
//
// @param data The original struct
// @return A map[string]interface{} containing data defined in original struct
func Map(data interface{}) map[string]interface{} {
	if data == nil {
		return make(map[string]interface{})
	}
	kind := retrieveElement(reflect.TypeOf(data)).Kind()
	if kind != reflect.Struct {
		return make(map[string]interface{})
	}
	return handleStruct(reflect.ValueOf(data))
}

func retrieveElement(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return retrieveElement(t.Elem())
	}
	return t
}

func retrieveElementValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return retrieveElementValue(v.Elem())
	}
	return v
}

// Omit kinds like Chan, Func, Interface and UnsafePointer
func handleValue(v reflect.Value) interface{} {
	kind := v.Kind()
	switch kind {
	case reflect.Int:
		return int(v.Int())
	case reflect.Int8:
		return int8(v.Int())
	case reflect.Int16:
		return int16(v.Int())
	case reflect.Int32:
		return int32(v.Int())
	case reflect.Int64:
		return v.Int()
	case reflect.Uint:
		return uint(v.Uint())
	case reflect.Uint8:
		return uint8(v.Uint())
	case reflect.Uint16:
		return uint16(v.Uint())
	case reflect.Uint32:
		return uint32(v.Uint())
	case reflect.Uint64:
		return v.Uint()
	case reflect.Float32:
		return float32(v.Float())
	case reflect.Float64:
		return v.Float()
	case reflect.Complex64:
		return complex64(v.Complex())
	case reflect.Complex128:
		return v.Complex()
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		return v.String()
	case reflect.Array, reflect.Slice:
		return handleArray(v)
	case reflect.Map:
		return handleMap(v)
	case reflect.Ptr:
		return v.Elem()
	case reflect.Uintptr:
		return v.Interface()
	case reflect.Struct:
		return handleStruct(v)
	default:
		return nil
	}
}

func handleArray(v reflect.Value) interface{} {
	result := reflect.MakeSlice(v.Type(), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		result = reflect.Append(result, (v.Index(i)))
	}
	return result.Interface()
}

func handleMap(v reflect.Value) interface{} {
	result := reflect.MakeMap(v.Type())
	for _, k := range v.MapKeys() {
		result.SetMapIndex(k, v.MapIndex(k))
	}
	return result.Interface()
}

func handleStruct(v reflect.Value) map[string]interface{} {
	self := make(map[string]interface{})
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		self[GenerateFieldName(field)] = handleValue(retrieveElementValue(v.FieldByName(field.Name)))
	}
	return self
}
