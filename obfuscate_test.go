package gojsonobfuscator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// returns true if the map or any key value pair in the map is null
func nullCheck(a map[string]any) (string, bool) {
	if a == nil {
		return "whole map", true
	}
	for k, v := range a {
		vReflect := reflect.ValueOf(v)
		if vReflect.Kind() == reflect.Map {
			v, ok := v.(map[string]any)
			if ok {
				if key, null := nullCheck(v); null {
					return key, true
				}
			} else {
				return "type assertion failed", true
			}
		}
		if v == nil {
			return k, true
		}
	}
	return "", false
}

func TestObfuscate(t *testing.T) {
	type A struct {
		Name  string
		Age   int64
		Money float32
	}
	type F struct {
		B A
		A A
	}
	f := F{
		B: A{
			Name:  "test",
			Age:   15,
			Money: 24.5,
		},
		A: A{
			Name:  "test",
			Age:   89_120,
			Money: 12.5,
		},
	}
	obf, err := Obfuscate(f)
	if err != nil {
		t.Error(err)
	}
	if key, null := nullCheck(obf); null {
		t.Errorf("%q was null", key)
	}
	out, _ := json.MarshalIndent(obf, "", "\t")
	fmt.Println(string(out))
}
