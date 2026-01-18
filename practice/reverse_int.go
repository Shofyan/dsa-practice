package main

func revese(x int) int {
	res := 0
	INT_MAX := 1<<32 - 1
	INT_MIN := -1 << 32

	for x != 0 {
		digit := x % 10
		x /= 10

		// cek overflow
		if res > INT_MAX/10 || (res == INT_MAX/10 && res/10 > 7) {
			return 0
		}
		if res < INT_MIN/10 || (res == INT_MIN/10 && res/10 < -1) {
			return 0
		}

		res = res*10 + digit
	}

	return res
}
