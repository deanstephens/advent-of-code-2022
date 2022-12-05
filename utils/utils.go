package utils

func Reverse[S ~[]E, E any](s S) []E {
	r := s[:]
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}

func Pop[S ~[]E, E any](s S) E {
	popped := s[len(s)-1]
	s = s[:len(s)-1]

	return popped
}

func IndexOf[S ~[]E, E comparable](s S, e E) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}

	return -1
}
