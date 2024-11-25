// // package main

// // import "fmt"

// // func main() {
// // 	fmt.Println("=== Pertemuan 9 ===")

// // 	username := "agus"
// // 	password := "124"
// // 	role := "student"

// // 	if username == "agus" && password == "1234" {
// // 		fmt.Println("anda berhasil login")
// // 		if role == "student" {
// // 			fmt.Println("Welcome, student!")
// // 		} else if role == "teacher" {
// // 			fmt.Println("Welcome, teacher!")
// // 		} else {
// // 			fmt.Println("Welcome, guest!")
// // 		}
// // 	} else {
// // 		fmt.Println("Invalid username or password")
// // 	}

// // }

// package main

// import "fmt"

// func main() {
// 	var bilangan int
// 	fmt.Print("Masukkan bilangan: ")
// 	fmt.Scan(&bilangan)

// 	if bilangan < 0 {
// 		bilangan = -1 * bilangan
// 	}

// 	fmt.Println("Hasil mutlaknya adalah: ", bilangan)
// }

package main

import "fmt"

func main() {
	var kar string
	fmt.Print("Masukkan karakter: ")
	fmt.Scan(&kar)

	if kar <= "a" && kar >= "z" || kar <= "A" && kar >= "Z" {
		fmt.Println("Karakter yang anda masukkan adalah huruf")
	} else {
		fmt.Println("Karakter yang anda masukkan adalah bukan huruf")
	}
}
