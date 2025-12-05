package utils

import "slices"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveEmptyStrings(l []string) []string {
	return slices.Collect(func(yield func(string) bool) {
		for _, v := range l {
			if len(v) != 0 {
				if !yield(v) {
					return
				}
			}
		}
	})
}
