package hamming

import (
	"strings"
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Boom!!")
	}

	aSplit := strings.Split(a, "")
	bSplit := strings.Split(b, "")
	n := 0

	for i := range aSplit {
		if aSplit[i] != bSplit[i] {
			n++
		}
	}

	return n, nil
}
