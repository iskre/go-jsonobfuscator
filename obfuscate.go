package gojsonobfuscator

import (
	"fmt"
	"reflect"
)

var config = Config{
	HashKeys:      false,
	EncryptValues: false,
	Ints:          true,
	Floats:        true,
	Strings:       true,
	Secret:        "",
}

var depth = 1

// defines how deep the obfuscator will move in a struct
const maxdepth = 256

// returns the obfuscated given atomic value `a`
func obfuscateAtomic(a any) any {
	var r any
	r = a
	switch a.(type) {
	case float32, float64:
		if config.Floats {
			r = fmt.Sprintf("%0x", a)
		}
	case int, int8, int16, int32, int64:
		if config.Ints {
			r = fmt.Sprintf("%0x", a)
		}
	case string:
		if config.Strings {
			r = []byte(a.(string))
		}
	default:
		return a
	}
	return r
}

func createMapFromStruct(obj interface{}, obfuscateFunc func(any) any) (map[string]interface{}, error) {
	r := make(map[string]interface{}, 0)
	ref := reflect.ValueOf(obj)
	if ref.Type().Kind() != reflect.Struct {
		return nil, newObfuscationError("obj is not a structure")
	}
	for _, field := range reflect.VisibleFields(ref.Type()) {
		var val any
		reflectedVal := ref.FieldByName(field.Name)
		fieldName := field.Name
		switch reflectedVal.Kind() {
		case reflect.Invalid:
			return nil, newObfuscationError("hit invalid value")
		case reflect.String:
			val = reflectedVal.String()
		case reflect.Int:
			if reflectedVal.CanInt() {
				val = reflectedVal.Int()
			} else {
				newObfuscationError("failed to get int from int field")
			}
		case reflect.Int8:
			if reflectedVal.CanInt() {
				val = int8(reflectedVal.Int())
			} else {
				newObfuscationError("failed to get int8 from int8 field")
			}
		case reflect.Int16:
			if reflectedVal.CanInt() {
				val = int16(reflectedVal.Int())
			} else {
				newObfuscationError("failed to get int16 from int16 field")
			}
		case reflect.Int32:
			if reflectedVal.CanInt() {
				val = int32(reflectedVal.Int())
			} else {
				newObfuscationError("failed to get int32 from int32 field")
			}
		case reflect.Int64:
			if reflectedVal.CanInt() {
				val = int64(reflectedVal.Int())
			} else {
				newObfuscationError("failed to get int64 from int64 field")
			}
		case reflect.Float32:
			if reflectedVal.CanFloat() {
				val = float32(reflectedVal.Float())
			} else {
				newObfuscationError("failed to get float32 from float32 field")
			}
		case reflect.Float64:
			if reflectedVal.CanFloat() {
				val = reflectedVal.Float()
			} else {
				newObfuscationError("failed to get float64 from float64 field")
			}
		case reflect.Struct:
			if depth >= maxdepth {
				return nil, newObfuscationError("recursive iteration depth of 256 hit")
			}
			depth++
			t, err := Obfuscate(reflectedVal.Interface())
			if err != nil {
				return nil, err
			}
			val = t
		}
		r[fieldName] = obfuscateFunc(val)
	}
	return r, nil
}

// Obfuscates the given structure, returns a map of key value pairs included in
// the original structure but with obfuscated values.
//
// Throws an error if:
//   - `obj` is not a structure
//   - `obj` contains invalid field values
//   - extracting values with types from structure fails
//   - the recursive depth counter exceeds 256
func Obfuscate(obj interface{}) (map[string]interface{}, error) {
	return createMapFromStruct(obj, obfuscateAtomic)
}
