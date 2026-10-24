package util_inner

// ClampNumber 钳位函数
func ClampNumber(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
