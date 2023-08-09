package gojsonobfuscator

type ObfuscationError struct {
	txt string
}

type DeObfuscationError struct {
	txt string
}

func (o *ObfuscationError) Error() string {
	return o.txt
}

func (d *DeObfuscationError) Error() string {
	return d.txt
}

func NewObfuscationError(txt string) *ObfuscationError {
	return &ObfuscationError{txt}
}

func NewDeObfuscationError(txt string) *DeObfuscationError {
	return &DeObfuscationError{txt}
}
