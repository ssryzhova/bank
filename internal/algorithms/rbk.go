package algorithms

const base = 256
const mod = 101

func RabinKarp(text, pattern string) []int {
	var result []int

	n := len(text)
	m := len(pattern)

	if m > n {
		return result
	}

	var h, p, t int = 1, 0, 0

	for i := 0; i < m-1; i++ {
		h = (h * base) % mod
	}

	for i := 0; i < m; i++ {
		p = (base*p + int(pattern[i])) % mod
		t = (base*t + int(text[i])) % mod
	}

	for i := 0; i <= n-m; i++ {

		if p == t {
			match := true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				result = append(result, i)
			}
		}

		if i < n-m {
			t = (base*(t-int(text[i])*h) + int(text[i+m])) % mod
			if t < 0 {
				t += mod
			}
		}
	}

	return result
}
