package main

import "fmt"

func main() {

	// if < boolean expression(boolen açıklama)> {code}

	/* x := 27

	//x%2==0 bu ifade boolen bir ifadedir. ve bu ifade falsedur bu neden 1.if kısmı çalışmaz.
	if x%2 == 0 {
		fmt.Println(x, "çift saydır.")

	}
	if x%2 != 0 {
		fmt.Println("sonu alınmadı")
	} */

	/* //!false not false anlamında yani ! ünlem not anlamındadır.
	//if den sonra bir boolen değer bize dönmelilik if yapısı çalışsın
	//!true olsa sonuç bize vermeyecek bunun nedeni false dönmesidir.
	if !false {
		fmt.Println("Mesaj Gösteriilecek")
	} */

	/* //buradaki boolen değer de true döndüğü için sonuçu görüyoruz. Yani istenen değer boolen matıklı ve istenen ile uyumlu ise 1. kısım çalışır.
	if 5 > 3 {
		fmt.Println("MEsaj Gösterilecek mi")
	} */

	//Koşul sağlanıyorsa if kısımı çalışacak
	//Koşul sağlanmıyorsa else kısmı çalışacak
	//Aşağıdaki kısma branch denir dallandırma else oldugu için bu böyle olur.
	//Koşul nerde sağlanırsa döngü biter ve bir sonraki dala aşamaya geçmez
	/* x := 27
	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")
	} else { //yanlış durumu belirlemek için else yapısı kullanılır.
		fmt.Println(x, "tek sayıdır.")
	} */

	/* 	//if <boolen expression> {code} else {code}
	   	if false {
	   		fmt.Println("Mesaj Gösterilecek")

	   	} */

	//if yapılarında birden fazla dallandırma yapılabilir.
	x := -5

	if x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else if x%2 == 0 { //x sin 2 ile bölümünden kalan 0 ise tam bölünüyor demektir. else if ile bunu tekrar kontrol edebiliyoruz.
		fmt.Println(x, "çift sayıdır.")
	} else {
		fmt.Println(x, "tek sayıdır.")
	}

	//if <boolen expression> {code} else if <boolen expression> else {code}
	//burada önemli olan else if istenilen kadar kullanılabilir en son else ile bitireibiliriz
	//yani 10 adet else if yada daha fazla kullanılabilir bizim kullanmamız kaç tanesi o kadar kullanılabilir.

}
