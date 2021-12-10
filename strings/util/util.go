package util

import (
	"strings"

	"github.com/arknable/goseries-pkgfunc/strings/internal"
)

// Function "Join" dapat menggunakan package internal.
func Join(s []string) string {
	for i := range s {
		s[i] = internal.ToLower(s[i])
	}
	return strings.Join(s, " ")
}
