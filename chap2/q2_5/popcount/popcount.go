package popcount

func PopCount(x uint64) int {
	var res int = 0
	for i := 0; i < 64; i++ {
		cleared_x := x & (x - 1)
		if x == 0 {
			break
		}
		res += 1
		x = cleared_x
	}
	return res
}
