package main

import "fmt"

func main() {

	//Senaryo
	//İlk okulda 1 5 arası alınan notlarda 5 pekiyi 1 kötü not
	//olarak verilecek.

	//Switch anahtar kelimesi öncelikle kullanılır.
	//Sonra karşılatırılacak durum yazılır grade:=3 bu duruma örnek, biz grade di karşılaştırıyoruz.
	//grade:=3 olarak verilen karşılatırmayı aşağıdaki kısımlarda karşılatırmaları gerçekleştirilir.

	//grade := 3 //buradaki grade değişkenini switch içerisine ; ile alabiliriz if else deki gibi

	//if grade == 5 {fmt.Println("Pekiyi")} bu ifade ile aşağıda yazılan switch ifade birebir aynıdır hiçbir fark yok sadece
	//go da switch if, else if yapısı karmaşık görünmesin if, else if yapısı daha basit görünsün diye kullanılan bir yapıdır. if ve else if den hiçbir farkı yoktur.
	//çok fazla yapıda switch case yapısı if else göre daha okunaklı oluyor tek farkı bu
	/* switch grade := 3; grade {
	case 5:
		fmt.Println("Pek iyi") */

	//değişken de kullanılan veri tipi ile case lerde kullanılan veri tipi aynı olmalıdır.
	grade := 4
	//default istenen yere yazılabilir.
	/* default:
	fmt.Println("Geçersiz Not") */

	switch /*grade := 3;*/ grade {
	case 5:
		fmt.Println("Pek iyi")

	case 4:
		fmt.Println("İyi")
		y := 100
		fmt.Println(y) //buradaki y değişkeni ve fmt çıktısı sadece case 4 içerisinde yazdırılır ve görülür. Diğer case lerden çağırılamaz.

	case 3:
		fmt.Println("Orta")

	case 2:
		fmt.Println("Geçer")

	case 1:
		fmt.Println("Başarısız")

	//Switch yapısında farklı bir durum girilmesinde else if deki else komutunun benzeri default tur.
	default:
		fmt.Println("Geçersiz Not")

	}

	switch {
	case false: //false olması nedeniyle konsolda terminalde gözükmez
		fmt.Println("Bu yazdığımız konsolda görünmez.")

	case true: //true olması deniyle kosolda gözükür.
		fmt.Println("Bu yazdığımız konsolda görünecek.")

	}

	//Yukarıdaki switch yapısı aşağıda if, else if yapısı ile yazılmıştır.
	/* 	if grade == 5 {
	   		fmt.Println("Pekiyi")
	   	} else if grade == 4 {
	   		fmt.Println("İyi")
	   	} else if grade == 3 {
	   		fmt.Println("Orta")
	   	} else if grade == 2 {
	   		fmt.Println("Geçer")
	   	} else if grade == 1 {
	   		fmt.Println("Başarısız")
	   	} else {
	   		fmt.Println("Geçersiz Not")
	   	} */

}
