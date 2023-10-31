package f_test

import (
	"fmt"

	"atomicgo.dev/f"
)

type Person struct {
	Name string
	Age  int
}

func Example_demo() {
	// Format a string with a struct
	john := Person{Name: "Bob", Age: 22}
	fmt.Println(f.Format("${Name} is ${Age} years old", john))

	// Format a string with a map
	alice := map[string]any{"Name": "Alice", "Age": 27}
	fmt.Println(f.Format("${Name} is ${Age} years old", alice))

	// Evaluate an expression
	fmt.Println(f.Format("John is 22 years old: ${Age == 22}", john))

	// Ternary operator
	fmt.Println(f.Format("John is 22 years old: ${Age == 22 ? 'yes' : 'no'}", john))

	// Output:
	// Bob is 22 years old
	// Alice is 27 years old
	// John is 22 years old: true
	// John is 22 years old: yes
}
