package main

import "fmt"

var x = 10 //direk atama olarak değer verilir fonksiyon dşında paket içerisinde := tanımlama - atama yapılmaz.

func main() {

	/* fmt.Println("Merhaba") //Println den sonra yeni bir satıra geçiş yapılır.
	fmt.Print("Merhaba")//Yazdırıldıgında satır başlangıcı yapmaz
	fmt.Println("") //Println yeni bir satır başlangıcına geçer yani Print ve Printf gibi yan yana yazdırmaz.
	fmt.Printf("Merhaba")//Yazdırıldıgında satır başlangıcı yapmaz.

	*/

	/* name := "Arin" */

	/* fmt.Printf(name)
	fmt.Print(name)
	fmt.Println(name) */

	/* fmt.Print("Benim Adım", name)
	fmt.Println("") //Yan yana yazıdmaması için alt alta yazdırması için println yazıyoruz.
	fmt.Println("Benim Adım", name)
	fmt.Printf("Benim Adım %v %T ", name, name) */ //Yadırılan değişkenin type nı yazdırır.
	//Printf deki %T değişkenin tipini yazdırır. %v değişkenin değerini yazdırır. golang.org sitesinde hepsi var.

	/* x := 100
	y := 20
	z := 30

	fmt.Printf("%b %d %o", x, y, z) */

	//name, age := "Arin", 5

	//fmt.Print("Benim Adım ", name, ", ve ben", age, " yaşındayım")

	//fmt.Println("Benim Adım", name, "ve ben", age, "yaşındayım")//println kendi boşluklaırnı otomatik birakır.

	//fmt.Printf("Benim Adım %v, ve ben %v yaşındayım", name, age)
	//print ve println ham şekilde çıktıyı alır. printf de formatlanmış şekilde alıyor yani %v gibi özel işaretleri alıyor.

	// VISIBILITY (GÖRÜNÜRLÜLÜK)
	//Yazılan kodların paket içerisinde olması ve kendi fonksiyonu içerisinde bulunması gerekliliğine denir.

	fmt.Println(x)

	myFunc() //func myFunc değişkenin çalışması için önce func main içerisinde çağrılması gerekiyor.

}

func myFunc() {

	fmt.Println(x)

}
