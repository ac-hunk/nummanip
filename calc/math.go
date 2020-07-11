package calc

func Add(ary ...int) int {
	sum := 0
	for _, value := range ary {
		sum = sum + value
	}
	return sum
}
