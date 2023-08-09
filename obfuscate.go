package gojsonobfuscator

import (
	"reflect"
)

var depth = 1

const maxdepth = 256

func Obfuscate(obj interface{}) (map[string]interface{}, error) {
	r := make(map[string]interface{}, 0)
	ref := reflect.ValueOf(obj)
	if ref.Type().Kind() != reflect.Struct {
		return nil, NewObfuscationError("obj is not a structure")
	}
	for _, field := range reflect.VisibleFields(ref.Type()) {
		var val any
		reflectedVal := ref.FieldByName(field.Name)
		fieldName := field.Name
		switch reflectedVal.Kind() {
		case reflect.Invalid:
			return nil, NewObfuscationError("hit invalid value")
		case reflect.String:
			val = reflectedVal.String()
		case reflect.Int:
			if reflectedVal.CanInt() {
				val = reflectedVal.Int()
			} else {
				NewObfuscationError("failed to get int from int field")
			}
		case reflect.Int8:
			if reflectedVal.CanInt() {
				val = int8(reflectedVal.Int())
			} else {
				NewObfuscationError("failed to get int8 from int8 field")
			}
		case reflect.Int16:
			if reflectedVal.CanInt() {
				val = int16(reflectedVal.Int())
			} else {
				NewObfuscationError("failed to get int16 from int16 field")
			}
		case reflect.Int32:
			if reflectedVal.CanInt() {
				val = int32(reflectedVal.Int())
			} else {
				NewObfuscationError("failed to get int32 from int32 field")
			}
		case reflect.Int64:
			if reflectedVal.CanInt() {
				val = int64(reflectedVal.Int())
			} else {
				NewObfuscationError("failed to get int64 from int64 field")
			}
		case reflect.Float32:
			if reflectedVal.CanFloat() {
				val = float32(reflectedVal.Float())
			} else {
				NewObfuscationError("failed to get float32 from float32 field")
			}
		case reflect.Float64:
			if reflectedVal.CanFloat() {
				val = reflectedVal.Float()
			} else {
				NewObfuscationError("failed to get float64 from float64 field")
			}
		case reflect.Struct:
			if depth >= maxdepth {
				return nil, NewObfuscationError("recursive iteration depth of 256 hit")
			}
			depth++
			t, err := Obfuscate(reflectedVal.Interface())
			if err != nil {
				return nil, err
			}
			val = t
		}
		r[fieldName] = val
	}
	return r, nil
}
