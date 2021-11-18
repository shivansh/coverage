package a

import (
	"fmt"

	"example.com/b"
)

func foo() {
	fmt.Println("foo")
	b.Bar()
}
