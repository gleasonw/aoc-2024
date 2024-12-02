package utils

type Set map[string]struct{}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
