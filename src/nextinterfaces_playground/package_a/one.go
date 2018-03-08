// Package string_util contains utility functions for working with strings.
package package_a

import (
	"fmt"
)

func Public_func_one(s string) int {

	// reference Public function from same package without import
	fmt.Println("--> " + Public_func_two())

	// reference private function from same package
	fmt.Println("--> " + private_function())

	return len(s)
}
