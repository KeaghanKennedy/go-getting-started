package notes

import "fmt"

func Collections() {
	// 5.2 CREATING ARRAYS

	// These arrays are identical. Different syntax. See primitives.
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println()

	// Arrays are fixe-length, zero-indexed and bounded. Arr[4] will cause a compeller error.

	// Output:
	// [1 2 3]
	// [1 2 3]

	// 5.3 WORKING WITH SLICES

	// If we don't specify a length, the compiler will figure out that we want a slice, not an array.
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	// We can build slices off of each other.
	// Make 'subslice' from index 1 to end.
	s2 := slice[1:]
	// Make 'subslice' from index 0 to index 2.
	s3 := slice[:2]
	// Make 'subslice' from index 1 to index 2 (single element at index 1).
	s4 := slice[1:2]

	fmt.Println(s2, s3, s4)
	fmt.Println()

	// Output:
	// [1 2 3]
	// [1 2 3 4 5 6]
	// [2 3 4 5 6] [1 2] [2]

	// 5.4 USING THE MAP DATA TYPE

	// Maps are basically slices of key-value pairs.
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 32
	m["bar"] = 21
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
	fmt.Println()

	// Output:
	// map[foo:42]
	// 42
	// map[bar:21 foo:32]
	// map[bar:21]

	// 5.5 WORKING WITH STRUCTS
	type user struct {
		id        int
		firstName string
		lastName  string
	}

	// Declares new user which will be initalized with the 0 values for each field.
	var u user
	// We can then assign some values to this object
	u.id = 1
	u.firstName = "Harvey"
	u.lastName = "Dent"
	fmt.Println(u)

	// Or just do it like this....
	// If we spread the initialization over multiple lines, we have to end lines with a comma.
	// Otherwise the compiler will start inserting semicolons at the end of lines and it will crash. Sad.
	u2 := user{
		id:        1,
		firstName: "Harvey",
		lastName:  "Dent",
	}
	fmt.Println(u2)

	// Output:
	// {1 Harvey Dent}
	// {1 Harvey Dent}
}
