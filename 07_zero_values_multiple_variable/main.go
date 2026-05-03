package main

import "fmt"

func main() {

	/* //Aşağıdaki 3 farklı değişken tanımlama örneğidir.
	var name string = "Arin" //1. Yöntem
	var name2 = "Elis"       //2. Yöntem
	name3 := "Gamze"         //3. Yöntem Short değişken tanımlama olarak geçer.



	//Bu şekilde kısa yoldan da var ile değişken tanımlaması yapılabilir.
	var (
		age       int     = 40
		isMarried bool    = true
		weight    float32 = 72.5
	) */

	//Tek Satır içinde değişken tanımlama yukarıdaki değişkenlerin tamamı tek satırda tanımlanabiliyor.
	/* var name, age, isMarried, weight, height = "Arin", 40, true, 72.5, 172 */

	//Tek satırda var olmadan değişken tanımlama fark := atama operatörü kullanıyoruz.
	/* name, age, isMarried, weight, height := "Arin", 40, true, 72.5, 172 */

	var name string
	var weight float32
	var isMarried bool

	fmt.Println(name)      // string ---> "" boş tırnaklar gelir değer verilmesede Zero Value su string oldugu için boş tırnak gelir.
	fmt.Println(weight)    //numeric ifadelerde başlangıç değeri 0 olarak alınır Zero Value su.
	fmt.Println(isMarried) // bool ifadelerde başlangıç değer Zero Value su false olarak alınır.

	//fmt.Println(name2)
	//fmt.Println(name3)

	/* fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)
	fmt.Println(age) */
}
