//constants bir programın çalışma süresi boyunca değişmeyen veri değerlerine constants denir yani sabitler
//örnek olarak pi sayısı, ışık hızı gib gibi değişmeyen sabitlerdir. Golang da sabitler diğer dillere göre daha da çoktur.

package main

import (
	"fmt"
)

func main() {

	/* bunlarda sabitler constants dır.
	5
	3.14
	"passed"
	true */

	/* 	r := 3.0

	   	//Sabit constants tanımlama aşağıdaki gibidir.
	   	const pi float64 = 3.14

	   	areaOfCirle := 3.14 * (math.Pow(r, 2))

	   	fmt.Println(areaOfCirle) */

	//constants yani sabitleri isimlendirme değişkenleri isimlendirme ile aynıdır.
	/* const x int = 2
	const y float64 = 3.4
	const z string = "test"
	const t bool = false

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t) */

	//constants program boyunca değişmeyecek yani bunlara değer ataması program boyunca değişmeyecek.
	/* const x=2

	//x=5 const için yani sabit için yeni değer atamada bu şekilde yapılırsa hata verecektir. sabitlerde const da yani aşağıdaki x deki gibi değer ataması yapılamaz
	// x++
	//x=x+1

	fmt.Printf("%T,%v\n", x,x)
	*/

	//------------------------------------------------------------------------------------------
	//constants sabitlerin en önenmli özelliği derleme zamanına aittir.
	//değişkenler ise runtime ma yani çalışma zamamına aittirler.

	//const ----> compile time  anlamı yazılan kodun makine diline çevrilmesi
	//var -----> run time anlamı makine diline çevrilen kodun çalıştırılması
	//sabitleri compile time ne oldugunu biliyor. Ama değişkenler ise çalışma zamanında hesaplayarak ne oldugunu bilgisayar bilebiliyor.

	//var x= math.Pow(3,4)
	//const x= math.Pow(3,4) - burada hata verilir nedeni ise Pow run time zamınında oluşturulur const ise derleme compile time da oluşturulur bu nedenle hata verecektir.

	/* const x = 5 //const ta mutlaka değeri verilmelidir. Değeri verilmeyen const çalışmaz.
	fmt.Printf("%T, %v", x, x) */

	/* y := 3

	const x = y bu şekilde olursa y değişken oldugu ve çalışama zamanı farklı run time gibi olmasından dolayı hata verecektir.

	fmt.Printf("%T,%v\n", y, y)
	fmt.Printf("%T,%v\n", x, x) */

	/* const x = 1
	var y = 3

	fmt.Printf("%T,%v\n", x, x)
	fmt.Printf("%T,%v\n", y, y)
	fmt.Printf("%T,%v\n", x+y, x+y) //aynı veri tipinde olmasından dolayı ekrana yazıdma veriliyor. */

	const x, y = 3, 5

	fmt.Printf("%T,%v\n,", x, x)
	fmt.Printf("%T,%v\n", y, y)

}
