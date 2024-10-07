// // // package main

// // // import "fmt"

// // // func main() {
// // // 	// var a int
// // // 	// a = -128
// // // 	// fmt.Println("Nilai a adalah", a)
// // // 	var a uint8
// // // 	a = 71
// // // 	var b uint8
// // // 	b = 21
// // // 	var c uint8
// // // 	c = a * b
// // // 	fmt.Println("C =", c)

// // // 	fmt.Println("a + b =", a+b)
// // // 	fmt.Println("a - b =", a-b)
// // // 	fmt.Println("a * b =", a*b)
// // // 	fmt.Println("a / b =", a/b)
// // // 	fmt.Println("a % b =", a%b)

// // // 	var hasilGeserKekiri uint16
// // // 	hasilGeserKekiri = 71 << 2
// // // 	var hasilGeserKekanan uint16
// // // 	hasilGeserKekanan = 71 >> 10

// // // 	fmt.Println("71 << 2 =", hasilGeserKekiri)
// // // 	fmt.Println("71 >> 2 =", hasilGeserKekanan)

// // // 	fmt.Println("71 & 51 =", 71&51)
// // // 	fmt.Println("71 | 51 =", 71|51)
// // // }

// // package main

// // func main() {
// // 	// var d, d0, t, v int
// // 	// // d0 ini adalah posisi awal meter
// // 	// // t ini adalah waktu
// // 	// // v ini adalah kecepatan
// // 	// fmt.Scan(&d0, &v, &t)
// // 	// d = d0 + (v * t)
// // 	// fmt.Println(d)

// // 	var a float32
// // 	var b float64

// // }

// package main

// import "fmt"

// func main() {
// 	var reamur, celcius, fahrenheit, kelvin float32
// 	fmt.Scan(&celcius)
// 	reamur = celcius * 4 / 5
// 	fahrenheit = celcius*9/5 + 32
// 	kelvin = celcius + 273.15

// 	fmt.Println("Reamur =", reamur)
// 	fmt.Println("Fahrenheit =", fahrenheit)
// 	fmt.Println("Kelvin =", kelvin)
// }

package main

import "fmt"

func main() {
	var p, l, K, L float32
	fmt.Scan(&p, &l)
	K = 2 * (p + l)
	L = p * l
	fmt.Println("Keliling =", K)
	fmt.Println("Luas =", L)
}
