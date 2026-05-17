package game

func maxi(a, b int) int {
	if a < b { return b }
	return a
}

func mini(a, b int) int {
	if a < b { return a }
	return b
}
