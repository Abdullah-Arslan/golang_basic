/*
	Boyut (Size)

•	Array: Sabit bir boyuta sahiptir. Dizinin boyutu tanımlandığı anda belirlenir ve daha sonra değiştirilemez. Boyut, dizinin türünün bir parçasıdır.
•	Örnek: [3]int ve [4]int farklı türlerdir.
•	Slice: Dinamik bir boyuta sahiptir. İhtiyaca göre büyüyüp küçülebilirler. Bir dilimin boyutu türünün parçası değildir.
2. Bellek Yönetimi ve Geçiş (Passing)
•	Array: Bir değer tipidir (value type). Bir fonksiyon diziyi parametre olarak aldığında, dizinin tüm içeriği kopyalanır. Yani orijinal dizi üzerinde yapılan değişiklikler kopyayı etkilemez.
•	Slice: Bir referans tipidir (aslında bir "header" yapısıdır). Bir dilimi fonksiyona geçirdiğinizde, dizinin kendisi kopyalanmaz; sadece arka plandaki diziye (underlying array) işaret eden bir işaretçi (pointer), uzunluk ve kapasite bilgisi kopyalanır. Bu nedenle, bir fonksiyon içinde dilimin elemanlarını değiştirirseniz, orijinal dilim de değişir.
*/
package main

import "fmt"

//Slices lar Arrays ile benzerdirler, fakat Slices lar daha gelişmişlerdir.

func main() {

	/* //Array oluşturuyoruz.
	myArr := [3]int{1, 2, 3}       //1. Array gösterim şekli
	myArr2 := [...]int{1, 2, 3, 4} //2. Array gösterim şekli
	fmt.Println(myArr)
	fmt.Println(myArr2)
	fmt.Println(len(myArr))
	fmt.Println(len(myArr2))

	var myArr3 [5]int //Buradaki Array ye bir değer girilmediği için değer olarak beş tane sıfır alır "zero value" döner. Slicelarda ise değer girilmeden baş olarak value dönüşü olmaz.
	fmt.Println(myArr3) */

	/* 	//Slices lar ile Array ler arasındaki en temel fark Arrayler oluşturulurken eleman sayısını bilirken Sliceslar eleman sayısını bilmiyorlar.
	   	//Slices oluşturma yapısı aşağıdaki gibidir.
	   	mySlc := []int{1, 2, 3}//Slices larda eleman sayısı belirtilmiyor. [] boş olarak giriliyor.
	   	fmt.Println(mySlc)
	   	fmt.Println(len(mySlc))//Elaman sayısını verir.
	*/

	/* 	//Burada Array oluşturma
	   	var myArr [4]int //Burada herhangibir deger ataması yapmadan Array oluşturulur.
	   	fmt.Println(myArr)
	   	myArr[0] = 5
	   	fmt.Println(myArr)

	   	//Burada Slices oluşturması yapıyoruz. Slices lar Array lerin fonksiyonel olarak daha da geliştirilmiş halidir.
	   	var mySlc []int
	   	mySlc = make([]int, 4)
	   	fmt.Println(mySlc)
	   	mySlc[0] = 10
	   	fmt.Println(mySlc) */

	/* myArr := [3]int{1, 2, 3}
	fmt.Println(myArr)
	myArr2 := myArr
	fmt.Println(myArr2)

	myArr2[0] = 100
	fmt.Println(myArr2)
	fmt.Println(myArr2) //Arrayler değerleri paylaşır.
	*/

	mySlc := []int{1, 2, 3}
	fmt.Println(mySlc)
	mySlc2 := mySlc
	fmt.Println(mySlc2)
	mySlc2[0] = 33
	fmt.Println(mySlc2)
	fmt.Println(mySlc) //Slice lar referansı paylaşır.

}
