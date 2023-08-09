package gojsonobfuscator

import (
	"testing"
)

func TestObfuscate(t *testing.T) {
	type A struct {
		Name string
		Age  int64
	}
	type F struct {
		Name string
		Age  int64
		A    A
	}
	f := F{
		Name: "xnacly",
		Age:  1,
		A: A{
			Name: "test",
			Age:  3,
		},
	}
	obf, err := Obfuscate(f)
	if err != nil {
		t.Error(err)
	}
	for k, v := range obf {
		if v == nil {
			t.Errorf("a key value pair in the map is null: %v", k)
		}
	}
	if f.Name == obf["Name"] || f.Age == obf["Age"] {
		t.Errorf("%+v == %+v", obf, f)
	} else {
		t.Logf("%+v", obf)
	}
}
