package notes

func ProgramFlow() {
	// 7.3 CREATING LOOP THAT TERMINATES BASED ON CONDITION

	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			// Break out over entire loop.
			// break

			// Break out of current iteration of loop ("continuing..." will not be logged this time).
			continue
		}
		println("continuing...")
	}

	// Output:
	// 0
	// continuing...
	// 1
	// continuing...
	// 2
	// 3
	// continuing...
	// 4
	// continuing...

	// 7.4 USING CONDITIONAL LOOPS WITH POST CLAUSES

	// If using implicit initialization syntax, variable i will be scoped to the for loop.
	for j := 0; j < 5; j++ {
		println(j)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4

	// 7.5 CREATING INFINITE LOOPS

	var k int
	for {
		if k == 5 {
			break
		}
		println(k)
		k++
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4

	// 7.6 LOOPING OVER COLLECTIONS

	// We could iterate over a map here with k, v (key, value) instead of i, v (index, value).
	// Can omit second returned val. The compiler has a special case for returning multiple variables here.
	// e.g. for v := ...
	// If we want to omit the first, we have to use a write only variable.
	//e.g for _, v := ...
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	// 7.7 WORKING WITH THE PANIC FUNCTION

	// She's properly cooked. We can't recover from here
	//panic("Something went really wrong")

	// 7.8 CREATING IF STATEMENTS

	// Standard if/else logic
	n := "odd"

	if n == "odd" {
		println("She's odd")
	} else if n == "even" {
		println("She's even")
	} else {
		println("That would have to be zero right....?")
	}

	// Output:
	// She's odd

	// 7.10 WRITING SWITCH STATEMENTS
	r := "GET"

	// Switches in Go have implicit breaking. We don't need a 'break' at the end of each case.
	switch r {
	case "GET":
		println("Get request")
		// Fallthrough can be triggered with fallthrough statement.
		// fallthrough
	case "POST":
		println("Post request")
	case "DELETE":
		println("DELETE request")
	// Can provide a default case.
	default:
		println("Unhandled method")
	}

	// Output:
	// Get request

}
