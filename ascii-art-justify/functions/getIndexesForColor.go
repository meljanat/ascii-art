package functions

func GetIndexsForColor(text string, letters string) []int {
	slIndexs := []int{}
	for i := 0; i <= len(text)-len(letters); i++ {
		j := i
		to := i + len(letters)
		if text[i:to] == letters {
			for j < to {
				slIndexs = append(slIndexs, j)
				j++
			}
		}
	}
	return slIndexs
}
