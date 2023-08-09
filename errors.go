package gojsonobfuscator

type obfuscationError struct {
	txt string
}

type deObfuscationError struct {
	txt string
}

func (o *obfuscationError) Error() string {
	return o.txt
}

func (d *deObfuscationError) Error() string {
	return d.txt
}

func newObfuscationError(txt string) *obfuscationError {
	return &obfuscationError{txt}
}

func newDeObfuscationError(txt string) *deObfuscationError {
	return &deObfuscationError{txt}
}
