// Mendefinisikan package bernama "math".
// Satu directory harus mendefiniskan package dengan nama yang sama
// di posisi root-nya.
package math

// Sum adalah exported function (public),
// berarti dapat digunakan di luar package math.
func Sum(a, b int) int {
	return doSum(a, b)
}

// doSum adalah non-exported function (private),
// berarti tidak dapat digunakan di luar package math.
func doSum(a, b int) int {
	return a + b
}

// Function bisa memiliki return value lebih dari satu.
func Mod(a, b int) (int, int, int) {
	return a % b, a, b
}
