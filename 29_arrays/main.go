package main

import "fmt"

//Arrays karmaşık veri tipleri anlamındaki verileri göreceğiz. Go da aynı veri tipine ait elemandan birden fazlada varsa
//Bir sonuçtan bir tanede değilde 100 tane olması, bir mail adresinden yüzlercesi varsa yani aynı veri tipine ait birden fazla bilgi veri varsa
//Aynı veri tipine ait olan bu elemanları bir liste yapısı içerisinde kuruyoruz buna Array deniyor.

func main() {

	/* city1 := "istanbul"
	city2 := "roma"
	city3 := "tahran"
	city4 := "belgrad"

	fmt.Println(city1, city2, city3, city4) */

	//Yukarıdaki şehir isimlerini bir program için ayrı ayrı değişkenler ile yazamak zor ve kullanım açısından çok verilim değildir.
	//Bunun yerine yukarıdaki şehirleri bir arrays içine alıp yazmak daha kullanışlı ve okunaklı olur.
	//Array oluşturma, array oluşturulurken bütün arrayler aynı veri tiplerinde olmalıdır.
	//Veri tipi  arraylerde aşağıdaki gibi oluşturulur.
	//cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

	/* cities := [...]string{"istanbul", "roma", "tahran", "belgrad"} //buradaki ... nokta sayısı bilinmeyen arrayleri çağırmak için kullanılır.

	fmt.Println(cities)
	fmt.Println(cities[0]) //go arrayleri sıfırdan başlayan indexleri kullanırlar.
	fmt.Println(cities[3]) //yandaki kısımda dogrudan array içerisinde veri çağırmalarda yapabiliriz.

	fmt.Println(len(cities))//burada array in kaç elemandan oluştugunu görmek için len komutunu kullanıyoruz.
	*/
	/*
		//Başlangıç değeri kendinden olan array örneği aşağıdaki gibidir.
		var myArr [5]int
		fmt.Println(myArr) //Bunun çıktısı 0 0 0 0 0 beş adet sıfır olarak verilir. Bunun nedeni kendiliğinden değer verilmesidir. Kaç eleman istendi ise o kadar değerde default olarak 0 gelir.

		//myArr[0]="istanbul"//buradaki atama çalışmaz nedeni ise array lerde verilen değerlerin veri tipi aynı olmalıdır. Buradaki değerin veri tipi string yukarıdakiler ise int olarak verilmiştir.

		myArr[0] = 100
		fmt.Println(myArr)        //Buradaki atama çalıştı bunun nedeni verilen değer int tipinde olmasından dolayıdır. Yani yukarıdaki myArr kısmı ile aynı veri tipi ile uyumludur.
		myArr[len(myArr)-1] = 200 //Buradaki -1 array deki son elemana değer vermek için kullanılır bu şekilde sondan değer vererek atama - ile yapılabilir.
		fmt.Println(myArr)
	*/

	/* //Beş adet int veri tipine sahip array yazmak
	//Aşağıdaki myArr ve myArr2  biri 5 elemana sahip diğeri 4 elemana sahip bunlardaki alacakları değer sayıları farklı olmasından dolayı veri tipleride aynı sayılmıyor.
	var myArr [5]int
	var myArr2 [4]int

	fmt.Println(myArr)
	fmt.Println(myArr2) */

	/* //Biz arraylerde ne kadar uzunluk yada eleman sayısı oldugunu bilmediginiz zaman for döngüsünü kullanabiliriz.
	cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}
	cities[0] = "ANKARA"

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	} */

	/* //Sıfırdan 9 za kadar oluşan rakamların karesini alacak array oluşturma
	myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//Oluşan array i mySquare ye paremetre olarak alacağımız kısım
	myArr = mySquare(myArr) //First Class Functions burada fonksiyonu bir değişkene atıyoruz. myyArr değişken mySquare fonksiyondur.

	fmt.Println(myArr) */

	//FOR --- RANGE YAPISI
	//for index, value := range myArr - for ve range anahtar kelimelerdir.

	cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

	/* for index, city := range cities {//range buradaki tüm elemanları yazdırmamızı sağlıyor.
		fmt.Println(index, city)
	} */

	for _, city := range cities { //range buradaki tüm elemanları yazdırmamızı sağlıyor.
		fmt.Println(city)
	}

}

/* //Yukarıdaki arrayin tamamını işlemlerden geçirerek karesini bulacağız.

func mySquare(arr [10]int) [10]int { //Burada parametre olarak array giriliyor aynı şekilde return olarak da array dönüyor.
	for i := 0; i < len(arr); i++ { //Buradaki for döngüsü ile her bir array elemanını alıyoruz.
		arr[i] = arr[i] * arr[i] //Burada her bir array elemanına kendisinin karesini atıyoruz.
	}
	return arr
}
*/
