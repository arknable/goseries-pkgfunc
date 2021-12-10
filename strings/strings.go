package strings

import "github.com/arknable/goseries-pkgfunc/strings/internal"

func Concat(a, b string) string {
	return a + " " + b
}

// Function "ToLower" dapat menggunakan package internal.
func ToLower(s string) string {
	return internal.ToLower(s)
}
