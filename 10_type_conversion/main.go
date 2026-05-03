package main

import "fmt"

func main() {

	/* 	x := 10
	   	y := 10.0

	   	fmt.Printf("%v %T\n", x, x)// % deli ifadelerin yanındakilerin çalışabilmesi için fmt.Printf ile çağrılması gerekiyor.
	   	fmt.Printf("%v %T\n", y, y)

	   	//Type Conversion type(value) => int(y) =10.0 => 10

	   	fmt.Println(x + int(y)) //Buradaki kısımda bir int(y) ile bir çevirme yok sadece ekrana yazdırılan kısımlar toplanıyor.

	   	fmt.Printf("%v %T\n", y, y) */

	/* var x int8 =10
	var y int16=10

	fmt.Println(x+y)//Buradaki kısımda int türleri dahi bir birinin aynısı olmalıdır. farklı oldugunda hata veriyor. */

	/* var x int8  =127
	var y int16

	y = int16(x)// type(value) küçük olan veri tipini büyük olan veri tipine dönüştürmek her zaman uygun olandır.

	fmt.Println(y) */

	/* 	x := 10
	   	y := "10"

	   	fmt.Printf("%v %T\n", x, x)
	   	fmt.Printf("%v %T\n", y, y)

		   fmt.Println(x + int(y)) */ //Burada string bir veri türünü int e çevirilemez örneği vardır. Çevrilebilir ama farklı bir yöntemle

	num1 := 106
	str1 := string(106) //Ondalıklı sayıda 106 küçük j ye denk geldiği için j olarak yazdırıldı.

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1)

}
