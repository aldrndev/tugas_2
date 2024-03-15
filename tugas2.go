package main

import "fmt"

func main(){
	angka := 20

	if angka % 2 == 0{
		fmt.Println(angka, "adalah bilangan genap")
	}else{
		fmt.Println(angka, "adalah bilangan ganjil")
	}
}