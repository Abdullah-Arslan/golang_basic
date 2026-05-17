//Neden fonksiyon kullanırız?

//**OLUŞTURULAN FONKSİYONLAR MAİN PAKETİ İÇERİSİNDE ÇALIŞIR.
//**MAİN İÇERİSİNDE ÇAĞRILMADAN FONKSİYONLAR ÇALIŞMAZ

package main

import "fmt"

func main() {

	/* var x, y, sum int
	x = 5
	y = 10
	sum = x + y
	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

	x = 7
	y = 11
	sum = x + y
	fmt.Println("%d ve %d toplamı %d\n", x, y, sum)
	*/
	//Neden fonksiyon kullanırız?
	//Daha modüler progrlama için
	//Rahat kod okuma yapmak için
	//Tek bir ana fonksiyon altında bir bütün kod blogu yerine teker teker işleme bölmerek daha rahat anlaşılır kod yazımı ve okuma

	//aşağıdaki fonksiyonu çalıştırmak için burada çağırmamız gerekiyor.
	//fmt.Println(sum(5, 10))

	//func <fonksiyon ismi> <(parametreler)> <return type döndürülecek tip> <{code - yazılacak kodlar}>

	//merhaba()
	//merhaba() //bu şekilde aynı fonksiyonu tekrar tekrar çağırabiliriz.

	/* //return vs print arasındaki fark
	z := sum(5, 10)
	fmt.Println(z) //bu şekilde de çağırabiliyoruz.

	sum2(6, 11) */

	merhaba2("Arslan", 6)

	//Fonksiyon İsimlendirme
	//İlk karekter harf olacak
	//camel Case -- mysum, myBestFunction yani ilk harf küçük ikinci ismin harfi büyük olacak şekilde başlanmalı
	//paket dışında kullanılacaksa ilk harf büyük olacak
}

//sum dan sonraki () parantezlere kullanılacak parametreleri yazıyoruz.
//bu yazılan fonksiyondan () parantezlerden sonra veri dönüşü return alınacağı için onunda türü yazılıyor.
func sum(x int, y int) int {

	return x + y //veri dönüşünü yani returnu bu şekilde alıyoruz.
	//yazılan bu fonksiyonu ana fonksiyon içerisinde çağırarak çalıştırabiliriz.

}

//her fonksiyonun kapsamı kendi içerisindedir.
func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Merhaba Dünya")
}
func merhaba2(name string, age int) { //burada verilen değerler main ana fonksiyonunda çağrılırken sıralamaya dikkat etmek gerekiyor.
	fmt.Printf("Adım %s, yaşım %d", name, age)
}

/* //bu şekilde parametre almadan da fonksiyon yazıdırılabilir.
func merhaba() {
	fmt.Println("Merhaba Dünya")
} */
