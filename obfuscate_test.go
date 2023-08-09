package gojsonobfuscator

import "testing"

func TestObfuscate(t *testing.T) {
	type F struct {
		Name string
		Age  int64
	}
	f := F{
		Name: "xnacly",
		Age:  1,
	}
	obf := Obfuscate(f).(F)
	if f.Name == obf.Name || f.Age == obf.Age {
		t.Errorf("%+v == %+v", obf, f)
	}
}
