package functions

func CountSpaces(input string) int {
	Spaces := 0

	for _, char := range input {
		if char == ' ' {
			Spaces++
		}
	}

	return Spaces
}
