package structs

import (
	"reflect"
	"strings"
)

func GenerateFieldName(field reflect.StructField) string {
	structsTag := field.Tag.Get("structs")
	tagValues := strings.Split(structsTag, ",")
	var fieldName string
	if len(tagValues[0]) > 1 {
		fieldName = tagValues[0]
	} else {
		fieldName = field.Name
	}
	return fieldName
}

func FindTypeMap(field reflect.StructField, typeMap map[string]reflect.Type) (reflect.Type, bool) {
	structsTag := field.Tag.Get("structs")
	tagValues := strings.Split(structsTag, ",")
	if len(tagValues) > 1 && len(tagValues[1]) > 0 {
		if t, ok := typeMap[tagValues[1]]; ok {
			return t, true
		}
	}
	return nil, false
}
