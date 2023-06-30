package pis

func PrintNbrInOrder(n int) {
	var t []int
	var r int
	if n == 0 {
		t = append(t, 0)
	}
	for n > 0 {
		r = n % 10
		t = append(t, r)
		n = n / 10
	}
	i := 1
	for i < len(t) {
		if t[i-1] > t[i] {
			tb := t[i]
			t[i] = t[i-1]
			t[i-1] = tb
			i = 1
		} else {
			i++
		}
	}
	for i := range t {
		z01.PrintRune(rune('0' + t[i]))
	}
}
