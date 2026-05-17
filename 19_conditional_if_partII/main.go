package main

import "fmt"

func main() {

	//x := -5 buradaki x aşağıdaki if yerine de alınabilir tek bir starı içerisine iki değer stetment yazman istersek
	//; noktalı virgül konulması gerekiyor

	/* if x := 5; x < 0 {
		fmt.Println(x, "negatif sayıdır")

	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır.")
	}
	*/
	//fmt.Println(x) buradaki fmt çalışabilmesi için x if blogu içerisinde çalışır. fmt çalışması için if blogu dışında olması lazım

	//go da iki satır yan yana kod olarak yazıdırllıacaksa ilk satırın yanına ; noktalı virgül konuluyor.
	//x := 25

	if x := 25; x < 0 {
		fmt.Println(x, "negatif sayıdır")

	} else {
		if x%2 == 0 {
			fmt.Println(x, "çift sayıdır")
		} else {
			fmt.Println(x, "tek sayıdır")
		}
	}

}
