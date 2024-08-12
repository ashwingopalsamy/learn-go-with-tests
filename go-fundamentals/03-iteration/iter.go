package iteration

func iter(s string) string {
	i := 0
	res := ""
	for i < 5 {
		res = res + s
		i++
	}
	return res
}
