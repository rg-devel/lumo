package lumo

import (
	"time"
	"fmt"
)

// Foo is my foo function
func Foo() error {
	fmt.Println("in Foo")
	time.Sleep(2 * time.Second)
	return nil
}
