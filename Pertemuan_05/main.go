package main

import "fmt"

func main(){
	var n, x, totalJarak int;
	fmt.Print("Masukkan jumlah data: ");
	fmt.Scan(&n);

	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan jarak ke-", i, ": ");
		fmt.Scan(&x);
		totalJarak = x * 5;
		fmt.Println("Jarak ke-", i, " adalah ", totalJarak);
	}


}