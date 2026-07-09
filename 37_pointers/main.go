package main

//BİR POİNTER BAŞKA BİR DEĞİŞKENİN ADRESİNİ TUTAN DEĞİŞKENDİR

import "fmt"

func main() {

	//name := "arin"

	/*
		fmt.Println(name)
		fmt.Println(&name) //& ---> address operator ü denir */

	//BİR POİNTER BAŞKA BİR DEĞİŞKENİN ADRESİNİ TUTAN DEĞİŞKENDİR

	//x := 22 //Burada int olan verileri sadece pointer yapabilir.
	/* fmt.Println(x)
	fmt.Println(&x)

	//FORMATLANMIŞ PRİNTLERİN GÖSTERİMİ
	fmt.Println()
	fmt.Printf("%T,%v\n", x, x)
	fmt.Printf("%T,%v\n", &x, &x) */

	/* 	y := &x
	   	fmt.Printf("%T,%v\n", y, y)

	   	z := &name
	   	fmt.Printf("%T,%v\n", z, z) */

	/* 	x := 22
	   	fmt.Println(x)        //22 gösterir
	   	fmt.Println(&x)       //22 nin adresini gösterir ---> veri tipi *int
	   	fmt.Println(*(&x))    //derefencing biz tekrar * yıldız operatörü ile var olan değeri döner.
	   	fmt.Println(&(*(&x))) // * ---->ilgili adresteki değeri gösterir
	   	fmt.Println(*(&(*(&x))))

	   	fmt.Println(3 * 5) */

	/* x1 := 10
	x2 := x1
	fmt.Println(x1, x2)
	x1 = 5
	fmt.Println(x1, x2) */

	/* x1 := 10
	x2 := &x1
	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)

	*x2 = 3 //x2 burada pointer dır.
	fmt.Println(x1, *x2)

	x3 := &x1
	*x3 = 5
	fmt.Println(x1, *x2, *x3) */

	//x1 := [4]int{1, 10, 100, 1000}//array  pass by value

	x1 := [4]int{1, 10, 100, 1000}

	x2 := x1
	fmt.Println(x1, x2)

	x2[0] = 3
	fmt.Println(x2)
	fmt.Println(x1)

}
