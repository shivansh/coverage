package c

import (
	"fmt"

	"example.com/b"
)

func qux() {
	b.Bar()
	b.Baz()
	fmt.Println("qux")
}
