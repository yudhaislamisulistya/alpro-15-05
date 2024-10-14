// // // package main

// // // import "fmt"

// // // func main() {
// // // 	var bilangan int8

// // // 	var status bool

// // // 	bilangan = 0

// // // 	status = bilangan == 0

// // // 	fmt.Println("Status bilangan", bilangan, "adalah", status)
// // // }

// // package main

// // import "fmt"

// // func main() {
// // 	var bilangan int32
// // 	var status bool

// // 	bilangan = 76

// // 	status = bilangan%2 == 1

// // 	if status {
// // 		fmt.Println("Bilangan", bilangan, "adalah bilangan ganjil")
// // 	} else {
// // 		fmt.Println("Bilangan", bilangan, "adalah bilangan bukan merupakan bilangan ganjil (genap)")
// // 	}
// // }

// package main

// import "fmt"

// type bilangan int
// type desimal float64

// type Mahasiswa struct {
// 	nama  string
// 	nim   string
// 	kelas string
// }

// func main() {
// 	// var a bilangan = 10
// 	// var b desimal = 3.14

// 	// fmt.Println("a =", a)
// 	// fmt.Println("b =", b)
// 	var mhs Mahasiswa
// 	mhs.nama = "Budi"
// 	mhs.nim = "123456"
// 	mhs.kelas = "IF-1"

// 	fmt.Println("Nama =", mhs.nama)
// 	fmt.Println("NIM =", mhs.nim)
// 	fmt.Println("Kelas =", mhs.kelas)

// }

// package main

// import "fmt"

// const PI float64 = 3.14

// func main(){
// 	var a int = 10
// 	var luas float64

// 	PI = 1000

// }

package main

import "fmt"

const PI float64 = 3.14

// Tipe data bentukan ini berfungsi untuk mendeklarasikan
// tipe data baru yang terdiri dari beberapa tipe data yang sudah ada.
type Tabung struct {
	tinggi      float64
	jarijari    float64
	luas        float64
	volume      float64
	luasAlas    float64
	luasDinding float64
}

type Barang struct {
	nama   string
	jumlah int
	harga  float64
}

func main() {
	var tabung Tabung
	fmt.Scan(&tabung.jarijari, &tabung.tinggi)
	// Menghitung luas alas tabung dengan rumus
	// PI * r * r
	tabung.luasAlas = PI * (tabung.jarijari * tabung.jarijari)
	tabung.luasDinding = 2 * PI * tabung.jarijari * tabung.tinggi
	tabung.luas = 2*tabung.luasAlas + tabung.luasDinding
	tabung.volume = tabung.luasAlas * tabung.tinggi

	fmt.Println("Luas Alas =", tabung.luasAlas)
	fmt.Println("Luas Dinding =", tabung.luasDinding)
	fmt.Println("Luas Tabung =", tabung.luas)
	fmt.Println("Volume Tabung =", tabung.volume)
}
