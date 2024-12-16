package main

import "fmt"

func main(){
	var val int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&val)

	if val >= 0{
		fmt.Println("Nilai yang anda masukkan adalah positif")
		if val % 2 == 0 {
			fmt.Println("Nilai yang anda masukkan adalah genap")
		}else{
			fmt.Println("Nilai yang anda masukkan adalah ganjil")
		}
	} else {
		fmt.Println("Nilai yang anda masukkan adalah negatif")
	}
}