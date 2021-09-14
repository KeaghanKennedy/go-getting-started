package notes

import "fmt"

// 4.5 IOTA/CONSTANT DECLARATION

// Constant declarations work the same way that import blocks do.
// All constant rules regarding initialization/compiletime apply.
const (
	// As per the godocs, "Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants.".
	first  = iota
	second = iota
	third
)

const (
	fourth = iota * 3
	fifth
	sixth
)

func PrimitiveDataTypes() {
	// 4.2 PRIMITIVE TYPES

	// Super explicit. Declaration and initialization on separate lines.
	var i int
	i = 42
	fmt.Println(i)

	// More concise. For when you NEED to specify a type (it can't be inferred by the compiler).
	var f float32 = 3.14
	fmt.Println(f)

	// Maaad go energy. Type of first name will be inferred by the compiler here.
	firstName := "Arthur"
	fmt.Println(firstName)

	// Boolean
	b := true
	fmt.Println(b)

	// Complex numbers. Compulsory science...
	c := complex(3, 4)
	fmt.Println(c)

	// Multiple variables can be declared and initialized on the same line.
	r, im := real(c), imag(c)
	fmt.Println(r, im)
	fmt.Println()

	// Output:
	// 42
	// 3.14
	// Arthur
	// true
	// (3+4i)
	// 3 4

	// 4.3 POINTERS

	var lastName *string
	fmt.Println(lastName)
	fmt.Println()

	// Output:
	// <nil>

	// Initialize pointer using reference operator.
	var initializedName *string = new(string)
	// Assign value to location in memory using dereference operator.
	*initializedName = "Arthur"

	// Running this will print memory address (e.g. 0x40c128).
	fmt.Println(initializedName)
	// We wanna dereference to access string value.
	fmt.Println(*initializedName)
	fmt.Println()

	// Output:
	// 0x40c128
	// Arthur

	// Initialize pointer using address of operator.
	pntr := &firstName
	fmt.Println(pntr, *pntr)
	firstName = "Tricia"
	fmt.Println(pntr, *pntr)
	fmt.Println()

	// Output:
	// 0x40c128 Arthur
	// 0x40c128 Tricia

	// 4.4 CONSTANTS

	// Constants have to be declared in one line and known at compiletime.
	// You can't, for example, initialize a constant with the return value of a function.
	const constant int = 3

	// If we were letting the compiler infer the type of constant, this would work.
	// The compiler would reinterpret the type of constant each time its used.
	// fmt.Println(constant + 1.2)

	fmt.Println(float32(constant) + 1.2)
	fmt.Println()

	// 4.5 IOTA/CONSTANT BLOCKS

	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)

	// Output:
	// 0 1 2
	// 0 3 6
}
