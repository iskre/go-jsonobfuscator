# go-obfuscator

A go library to obfuscate and encrypt json objects.

> _Warning_
> This lib is in active development and currently has no strong obfuscation
> methods implemented.

## Example

```go
package main

import (
    gobf "github.com/iskre/go-jsonobfuscator"
	"encoding/json"
	"fmt"
)

type A struct {
    Name  string
    Age   int64
    Money float32
}
type F struct {
    B A
    A A
}
func main() {
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
    obf, _ := gobf.Obfuscate(f)
    jsonOutput, _ := json.MarshalIndent(obf, "", "\t")
    fmt.Println(string(jsonOutput))
}
```

Results in:

```json
{
  "A": {
    "Age": "15c20",
    "Money": "0x1.9p+03",
    "Name": "dGVzdA=="
  },
  "B": {
    "Age": "f",
    "Money": "0x1.88p+04",
    "Name": "dGVzdA=="
  }
}
```

## Documentation

## Configuration

`gobf`'s configuration enables fine grained control over the resulting
structure.

The default configuration can be overridden via the `gobf.SetConfig()`-Method,
that accepts a `gobf.Config`-Structure.

### Default configuration

```go
Config{
	HashKeys:      false,
	EncryptValues: false,
	Ints:          true,
	Floats:        true,
	Strings:       true,
	Secret:        "",
}
```

## Features

### Key hashing

Enabling `Config.HashKeys` hashes structure keys with `Config.Secret`

### Value encryption

Enabling `Config.EncryptValues` encrypts structure values with `Config.Secret`

### Value Obfuscation

Value obfuscation is used to describe the process of replacing the values of
the given go struct with obfuscated representations of the original values.

#### Strings

Enabling `Config.Strings` obfuscates structure values of type `string`

#### Integers

Enabling `Config.Ints` obfuscates structure values of type `int`, `int8`, `int16`, `int32`, `int64`

#### Floats

Enabling `Config.Ints` obfuscates structure values of type `float32`, `float64`
