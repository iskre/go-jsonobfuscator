package gojsonobfuscator

type Config struct {
	HashKeys      bool   // enables hashing structure keys, currently not implemented
	EncryptValues bool   // enables structure value encryption, currently not implemented
	Floats        bool   // enables obfuscation of floats
	Ints          bool   // enables obfuscation of integers
	Strings       bool   // enables obfuscation of strings
	Secret        string // used for encrypting values and hashing keys
}

// sets gojsonobfuscator.config to `conf`, therefore overrides the default
// config
func SetConfig(conf Config) {
	config = conf
}
