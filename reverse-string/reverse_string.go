package reverse

func Reverse(str string) string {
	rns := []rune(str)

	for inc, dec := 0, len(rns)-1; inc < dec; inc, dec = inc+1, dec-1 {
		rns[inc], rns[dec] = rns[dec], rns[inc]
	}

	return string(rns)
}
