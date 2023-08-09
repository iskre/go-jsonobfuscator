# go-obfuscator

A go library to obfuscate and encrypt json objects.

> _Warning_
> This lib is in active development and currently has no strong obfuscation
> methods implemented.

## Planned features

- [ ] encrypted values
- [ ] hashed keys
- [ ] configuration
  - [ ] enable and disable every feature
  - [ ] hash type and strength
  - [ ] secret for encryption

## Example

```go
package main

import (
    gobf "github.com/iskre/go-jsonobfuscator"
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

## Methodic

- gobf replaces the given go struct keys with obfuscated representations of their values
