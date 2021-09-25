package popcount

func PopCount(x uint64) int {
	var res int = 0
	for i := 0; i < 64; i++ {
		smaller_bit := x & 1
		if smaller_bit == 1 {
			res += 1
		}
		x = (x >> 1)
	}
	return res
}
