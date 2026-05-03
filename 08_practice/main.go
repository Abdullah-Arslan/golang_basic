// 1-) studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini
// 3 farklı yöntem ile oluşturup, çıktısını yazdırınız.

package main

import "fmt"

//func main() { */

//1.YÖNTEM

/* var studentName string = "Jhon Doe"
var grade int = 77
var isPassed bool = true

fmt.Println(studentName)
fmt.Println(grade)
fmt.Println(isPassed) */

//2.YÖNTEM

/* var studentName = "Jhon Doe"
var grade =77
var isPassed= true

fmt.Println(studentName)
fmt.Println(grade)
fmt.Println(isPassed) */

//3. YÖNTEM "var" KEYWORDS OLMADAN

/* 	studentName := "Jhon Doe"
	grade := 77
	isPassed := true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

}
*/

// 2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.
/*
func main() {

	/* var studentName, grade, isPassed = "Jhon Doe", 77, true

	fmt.Println(studentName, grade, isPassed) */

/* 	studentName, grade, isPassed := "Jhon Doe", 77, true

	fmt.Println(studentName, grade, isPassed)
}
*/

// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)

/* func main() {

	var studentName string = "Jhon Doe"// var studentName string Bu kısım Declaration oluyor.
									   // = "Jhon Doe" Assing etme deniyor. Yani atama
	studentName="Mahmut Erdem"//Assing etme atama kısımı.

	//var studentName string = "Jhon Doe" bu kısmın tamamına Initialization deniyor bir değeri oluşturup başlatma.
	//Initial Value ilk değer anlamında studentName in ilk değeri "Jhon Doe" yu almasına deniyor.

	fmt.Println(studentName)

}
*/

// 4-) "Statically Typed" vs "Dynamically Typed" ifadelerini GO ve Python üzerinden gösteriniz.

//VİDEO DAN TAKİBİNİ YAP

// 5-) ":=" vs "=" aradaki farkı gösteriniz, double declaration

func main() {

	/* 	var studentName string = "John Doe"
	   	studentName = "Mahmut Erdem" */

	studentName := "John Doe"    //Buradaki önemli bir fark var := ifadesi  değeri "Declaration" ve "Assign" ediyoruz yani ilk değer atamasını yapıyoruz anlamına gelir.
	studentName = "Mahmut Erdem" // Bu kısımdaki = ise Declaration yapmıyoruz yeni bir değer ataması yapıyoruz anlamına gelir

	fmt.Println(studentName)

}
