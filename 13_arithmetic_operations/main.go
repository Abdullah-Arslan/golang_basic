// addition + => 15 + 10 => 15, 10 operand(girilen değer demek), + operator(yapılacak işlem demek)
// substruction -
// product *
// division /
// remainder %

package main

import "fmt"

func main() {

	//x, y := 15, 10

	//fmt.Printf("%T,%v\n", x, x)
	//fmt.Printf("%T, %v\n", y, y)
	//fmt.Printf("%T,%v\n", (x + y), (x + y))
	//fmt.Printf("%T,%v\n", (x - y), (x - y))
	//fmt.Printf("%T,%v\n", (x * y), (x * y))
	//fmt.Printf("%T,%v\n", (x / y), (x / y))

	//fmt.Printf("%T,%v\n", (x % y), (x % y))

	//z := 5.00 / 2 //bu şekilde yazılırsa sonuç 2.5 çıkar. Veri tipide float64 olarak alınır. Yani 5.00 . ve 0 eklemek otamatik olarak float olarak alınır.
	//fmt.Printf("%T,%v\n", z, z)

	//Increment ++, Decrement -- POSTFIX

	x := 10

	fmt.Println(x)

	x = x + 1 //Sağdaki değeri soldakine atama yaparız.

	fmt.Println(x)

	x++

	fmt.Println(x)

}
