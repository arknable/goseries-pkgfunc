// Mendefinisikan package bernama "main".
package main

// Gunakan "import" untuk mengakses package lain.
// "fmt" adalah contoh package dari builtin.
// Sedangkan yang lainnya adalah package third-party (bukan builtin).
import (
	"fmt"

	"github.com/arknable/goseries-pkgfunc/math"
	"github.com/arknable/goseries-pkgfunc/math/alg"

	// Gunakan alias apabila package yang di-import
	// bentrok dengan package yang lain atau namanya terlalu panjang.
	mystrings "github.com/arknable/goseries-pkgfunc/strings"
	mystrings_util "github.com/arknable/goseries-pkgfunc/strings/util"
)

// Function bisa dijadikan tipe tersendiri.
// MathFunc adalah function yang menerima dua argument bertipe integer,
// dengan return value bertipe integer.
type MathFunc func(int, int) int

func main() {
	fmt.Println("1 + 1 = ", math.Sum(1, 1))
	fmt.Println("2 x 4 = ", alg.Multiply(2, 4))
	fmt.Println("'Hello' + 'World!! = ", mystrings.Concat("Hello", "World!!"))
	fmt.Println(`mystrings.ToLower("Halo-halo Bandung") = `, mystrings.ToLower("Halo-halo Bandung"))
	fmt.Println(`mystrings_util.Join("Lorem", "iPsum", "dolOr", "siT", "aMet")`, mystrings_util.Join([]string{
		"Lorem", "iPsum", "dolOr", "siT", "aMet",
	}))

	// Function bisa memiliki return value lebih dari satu.
	mod, a, b := math.Mod(5, 2)
	fmt.Printf("%d mod %d = %d\n", a, b, mod)

	// Function bisa juga dideklarasikan ditempat,
	// bisa disebut anonymous function.
	var subFunc MathFunc = func(a, b int) int {
		return a - b
	}
	fmt.Println("subFunc(5, 3): ", subFunc(5, 3))

	// Function math.Sum() bisa dipadankan dengan tipe MathFunc
	// karena strukturnya identik.
	var sumFunc MathFunc = math.Sum
	fmt.Println("sumFunc(2, 2): ", sumFunc(2, 2))
}
