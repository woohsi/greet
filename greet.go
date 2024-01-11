package greet

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hi [%s]", name)
}
