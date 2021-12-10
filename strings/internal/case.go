package internal

// Package "internal" hanya bisa digunakan dari dalam
// package yang menaunginya, yaitu package "strings".
// Ini berarti package "util" atau siblings-nya
// juga bisa menggunakan package "internal" ini.
import "strings"

func ToLower(s string) string {
	return strings.ToLower(s)
}
