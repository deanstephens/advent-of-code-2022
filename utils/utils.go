package utils

import (
	"math"
	"regexp"
)

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

func GetMagnitudeAndDirection2d(p1 int, p2 int) (magnitude int, direction int) {
	if p1 == p2 {
		return 0, 0
	}

	diff := p2 - p1
	magnitude = int(math.Abs(float64(diff)))
	direction = diff / magnitude

	return magnitude, direction
}

func FindNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}
	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}
	return results
}
