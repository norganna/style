package style

func runeEq(a, b []rune) bool {
	for i, r := range a {
		if r != b[i] {
			return false
		}
	}
	return true
}

func findSequence(start, end, input []rune) (sp, ep int, fn, body string) {
	sn := len(start) // Start number
	en := len(end)   // End number
	in := len(input) // Input number

	sfc := start[0] // Start first char
	efc := end[0]   // End first char

	var n int // Total number of chars to search, lessor of in - (en or sn).
	if en > sn {
		n = in - en
	} else {
		n = in - sn
	}

	sp = -1  // Start pos
	cp := -1 // Colon pos
	ep = -1  // End pos

	// check if it's the same tag on both ends, no easy way to tell starts from ends...
	same := en == sn && sfc == efc && string(start) == string(end)

	for p := 0; p <= n; p++ {
		c := input[p]

		// check end sequence first if end sequence is longer than start sequence and we already have a start sequence
		if sp > -1 && en > sn && !same && c == efc {
			es := input[p : p+en]
			if runeEq(end, es) {
				ep = p
				break
			}
		}

		// if current char is the same as the first char of the start sequence
		if c == sfc {
			ss := input[p : p+sn]
			if runeEq(start, ss) {
				if !same || sp < 0 {
					sp = p
					cp = -1
					p += sn - 1
					continue
				} else if same {
					ep = p
					break
				}
			}
		}

		// if the end sequence is different, we have a start pos and cur is same as first char of end sequence
		if sp > -1 {
			if !same && en <= sn && c == efc {
				es := input[p : p+en]
				if runeEq(end, es) {
					ep = p
					break
				}
			}

			if cp == -1 {
				if p > sp+sn && c == ':' {
					cp = p
					continue
				}

				// First chars must be < 7 letters `/[a-zA-Z]{1,7}/` long.
				if p > sp+sn+6 || c < 'A' || c > 'z' || (c > 'Z' && c < 'a') {
					sp = -1
					continue
				}
			}
		}
	}

	if sp > -1 && ep > -1 {
		if cp > -1 {
			fn = string(input[sp+sn : cp])
			body = string(input[cp+1 : ep])
		} else {
			fn = string(input[sp+sn : ep])
		}
		return
	}

	return -1, -1, "", ""
}
