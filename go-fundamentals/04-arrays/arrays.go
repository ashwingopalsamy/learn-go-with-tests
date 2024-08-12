package arrays

func ArraysSum(a [5]int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}
