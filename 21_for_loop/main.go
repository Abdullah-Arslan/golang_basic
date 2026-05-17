package main

import "fmt"

func main() {

	//i++ post ifadesi yani döngüde göndeirlmek istenen yada yapılmak istenen kısım ona yazılıyor.
	//tek satırda birden fazla işlem yapılırsa ; noktalı virgül ile ayırma yapıyoruz
	// for <init statement - başlangıç değeri>; <condition - şartımızı yazıyoruz>; <post statement - for döngüsü devam edecek ama nasıl devam edeceğini belirtiyoruz>
	/* for i := 1; i < 10; i++ {
		fmt.Println(i)
	} */

	//i:= 0 for dşında verildiğinde for döngüsü dışında da fmt ile çağrılabilir.
	//for ; i < 10; i++bu şekilde çalıştırmak için alttaki kısmı i ye kadar silemeliyiz
	/* for i := 1; i < 10; i++ {
		fmt.Println(i)
	} */

	/* for { //Infınıtı Loop sonsuz döngüye girer for
		fmt.Println("Benim Adım Arin")
	} */

	// buradaki koşul durumu true olmasından dolayı sonsuza kadar çalışacak.
	/* for i :=0; true; i =+5{
		fmt.Println(i)
	} */

	//burada döngü çalışmaz çünkü koşul ifadesi yani şart false olumsuz olmasından dolayı
	/* 	for i := 0; false; i = +5 {
		fmt.Println(i)
	} */

	//for yanında sadece şartı belirttiğimiz tanımlama kısmıda vardır.
	/* i := 10
	for i >= 0 {
		fmt.Println(i)
		i-- //i azalarak yazdırılır.
	} */

	//continue ve break kısımları
	/* 	for i := 0; i <= 10; i++ { // continue döngünün başına gider. koşul sağlandıgında
		if i%3 == 0 {
			continue //koşul sağlanırsa sağlandıgı yerden bırak ve for döngüsünün başına geri dön demektir.
		}
		fmt.Println(i)
	} */

	//break --> döngüden çıkar
	for i := 0; i <= 10; i++ {
		if i == 3 {
			break //break şartın sağlandıgı yerde döngüyü keser ve bitirir.
		}
		fmt.Println(i)
	}

}
