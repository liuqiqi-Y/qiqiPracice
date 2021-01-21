package commondivisor

func commonDivisor1(p, q int) int {
	for q != 0 {
		temp := q
		q = p % q
		p = temp
	}
	return p
}
func commonDivisor2(p, q int) int {
	if q == 0 {
		return p
	}
	return commonDivisor2(q, p%q)
}
