// Package bisa berada di dalam suatu package yang lain,
// bisa juga dianggap sebagai sub-package.
// Definisi package "alg" tidak akan konflik dengan "math"
// karena definisi "package" hanya berlaku untuk directory di posisi root-nya.
package alg

func Multiply(a, b int) int {
	return a * b
}
