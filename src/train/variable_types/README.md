[< Back Awards](../../../README.md)

## Topics

### What is the difference `int**` and `uint**`?

- `int**` type can allowed by negative and positive numbers.
- `uint**` type can allowed by positive numbers.

### Wath is the `rune` in golang?

- `Rune` serves as a alias for the `int32` tpye.
- It is specifically designed for representing Unicode points.
- Go uses UTF-8 encoding, a variable-length encoding scheme for Unicode charactrers.

- [Integer Rune (example)](./integer_rune.go)

```go
package main

import "fmt"

func integer_rune() {
	rune := 'G'
	fmt.Println("[integer_rune]", rune)
}
```