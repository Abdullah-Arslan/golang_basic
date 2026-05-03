package main

import "fmt"

func main() {

	//DEĞİŞKEN OLUŞTURMA VE DEĞİŞKENE DEĞER ATAMA

	var name string //var (keywords-anahtar kelime) - değişken adı - değişken tipi (veri tipi)
	//            var				-    name	   -        string
	name = "Arin"

	var isim string = "Yazilim" //değişkenleri bu şekilde kısaca yazılabilir.

	var fristName, lastName string = "Go", "Lang" //Bu örnekte birden fazla değişken kullanmayı gördük.

	var age int
	age = 40

	//var age int = 20 //Değişken kısa kullanımdan örnek.

	var isMarriedd bool
	isMarriedd = true

	//Değişkenleri daha kısa da tanımlayabiliriz. Keywords olmadan vani "var" olmadan da kullanabiliriz.
	kisaName := "Kisa Kullanim"
	kisaSayi := 60
	kisaBool := false

	kisaSayi = 61 //kisaSayi değişkenine yeni değer ataması bu şekilde yapılır.

	//Kısa Kullanım fmt alanı
	fmt.Println(kisaName)
	fmt.Println(kisaSayi)
	fmt.Println(kisaBool)
	//--------------------------------------------------
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarriedd)
	fmt.Println(isim)
	fmt.Println(fristName, lastName)

	/*
				NOT:


					Kısa karşılaştırma
					:= → tanımla + ata
					= → sadece ata


					  := (kısa değişken tanımlama):
					  -Yeni bir değişken tanımlamak ve değer atamak için kullanılır.
					  -Sadece fonksiyon içinde kullanılabilir.
				      -Tipi otomatik belirlenir.

					  *Örnek*
					  x := 10   // x burada ilk kez tanımlanıyor

		 ------------------------------------------------------------------------------------------

					  = (atama operatörü):
					  -Zaten tanımlanmış bir değişkene yeni değer atamak için kullanılır.
					  -Yeni değişken oluşturmaz.

					  *Örnek*
					  var x int = 10
					  x = 20   // burada sadece değeri değiştiriyoruz

	*/

}
