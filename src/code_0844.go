package src

func backspaceCompare(s string, t string) bool {
	ps, pt := len(s)-1, len(t)-1
	ss, st := 0, 0
	for ps > -1 || pt > -1 {
		if ps > -1 && s[ps] == '#' {
			ss++
			ps--
			continue
		}
		if pt > -1 && t[pt] == '#' {
			st++
			pt--
			continue
		}
		if st > 0 {
			pt--
			st--
			continue
		}
		if ss > 0 {
			ps--
			ss--
			continue
		}

		if ps > -1 && pt > -1 && s[ps] != t[pt] {
			return false
		} else if ps > -1 && pt < 0 && s[ps] != '#' {
			return false
		} else if pt > -1 && ps < 0 && t[pt] != '#' {
			return false
		}

		ps--
		pt--
	}
	return true
}
