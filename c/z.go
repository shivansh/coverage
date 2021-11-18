package c

import (
	"fmt"

	"example.com/b"
)

func qux() {
	b.Bar()
	fmt.Println("qux")
}
