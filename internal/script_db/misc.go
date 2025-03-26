package script_db

func arraymap[T any, U any](input []T, fn func(T) U) []U {
	output := make([]U, len(input))

	for i, v := range input {
		output[i] = fn(v)
	}

	return output
}

// Returns an array with all the elements that fn returns true
func arrayfilter[T any](input []T, fn func(T) bool) []T {
	output := make([]T, 0)

	for _, item := range input {
		if fn(item) {
			output = append(output, item)
		}
	}

	return output
}

func concat[T any](arrays ...[]T) []T {
	return flatMap(arrays)
}

func flatMap[T any](arrays [][]T) []T {
	output := make([]T, 0)

	for _, array := range arrays {
		for _, elem := range array {
			output = append(output, elem)
		}
	}

	return output
}

func remove[T any](array []T, n int) []T {
	output := make([]T, 0)

	for i, item := range array {
		if i >= n {
			output = append(output, item)
		}
	}

	return output
}

func removeBefore[T comparable](array []T, elem T) []T {
	output := make([]T, 0)

	adding := false
	for _, item := range array {
		if adding {
			output = append(output, item)
		}

		if item == elem {
			adding = true
		}
	}

	return output
}
