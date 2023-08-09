package gojsonobfuscator

import "testing"

func TestDeobfuscate(t *testing.T) {
	type F struct {
		Name string
		Age  int64
	}
	f := F{
		Name: "xnacly",
		Age:  1,
	}
	_, err := Obfuscate(f)
	if err != nil {
		t.Error(err)
	}
	// deobf := Deobfuscate(obf)
	// if f.Name != deobf.Name || f.Age != deobf.Age {
	// 	t.Errorf("%+v != %+v", deobf, f)
	// }
}
