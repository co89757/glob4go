package glob4go

import (
	"unicode"
)

func tolower(b byte) byte {
	return byte(unicode.ToLower(rune(b)))
}

//Glob matches a string against a glob wildcard pattern, nocase flag is for case-insensitive match
func Glob(pattern []byte, s []byte, nocase bool) bool {
	patternLen := len(pattern)
	stringLen := len(s)
	pattern = append(pattern, 0)
	s = append(s, 0)
	ip := 0
	is := 0
	for patternLen > 0 {
		switch pattern[ip] {
		case '*':
			for pattern[ip+1] == '*' {
				ip++
				patternLen--
			}
			if patternLen == 1 {
				return true
			}
			for stringLen > 0 {
				newpattern := make([]byte, len(pattern[ip+1:]))
				copy(newpattern, pattern[ip+1:])
				news := make([]byte, len(s[is:]))
				copy(news, s[is:])
				if Glob(newpattern, news, nocase) {
					return true
				}
				is++
				stringLen--
			}
			return false
		case '?':
			if stringLen == 0 {
				return false
			}
			is++
			stringLen--
		case '[':
			var not, match bool
			ip++
			patternLen--
			not = (pattern[ip] == '^')
			if not {
				ip++
				patternLen--
			}
			match = false
			for {
				if pattern[ip] == '\\' {
					ip++
					patternLen--
					if pattern[ip] == s[is] {
						match = true
					}
				} else if pattern[ip] == ']' {
					break
				} else if patternLen == 0 {
					ip--
					patternLen++
					break
				} else if pattern[ip+1] == '-' && patternLen >= 3 {
					start := pattern[ip]
					end := pattern[ip+2]
					c := s[is]
					if start > end {
						start, end = end, start
					}
					if nocase {
						start = tolower(start)
						end = tolower(end)
						c = tolower(c)
					}
					ip += 2
					patternLen -= 2
					if c >= start && c <= end {
						match = true
					}
				} else {
					if !nocase {
						if pattern[ip] == s[is] {
							match = true
						}
					} else {
						if tolower(pattern[ip]) == tolower(s[is]) {
							match = true
						}
					}
				}
				ip++
				patternLen--

			} // end of while(1)
			if not {
				match = !match
			}
			if !match {
				return false //no match
			}
			is++
			stringLen--
		case '\\':
			if patternLen >= 2 {
				ip++
				patternLen--
			}
			fallthrough
		default:
			if !nocase {
				if pattern[ip] != s[is] {
					return false
				}
			} else {
				if tolower(pattern[ip]) != tolower(s[is]) {
					return false
				}
			}
			is++
			stringLen--
		} // end of switch
		ip++
		patternLen--
		if stringLen == 0 {
			for pattern[ip] == '*' {
				ip++
				patternLen--
			}
			break
		}
	} //end of big while
	if patternLen == 0 && stringLen == 0 {
		return true
	}
	return false
}
