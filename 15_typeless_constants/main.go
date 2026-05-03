package main

import "fmt"

func main() {

	/* const x int16 = 5

	fmt.Printf("%T,%v\n", x, x) //fmt çıktısında 2 adet x,x yazılmasının nedeni %T veri tipi %v ise değeri verdigi için
	*/

	//const x int8 = 3 //x burada belirgin bir veri tipine sahip durumda typeless değildir.
	//veri tipleri farklı olan iki ifadeyi toplayamazsın diyor. Bunun temel nedeni comple time ve run time olmasından

	/* const x = 3 // burada const tın sabit bir veri tipi olmamasından printf de yani çıktıda hata vermiyor.
	var y int16 = 12

	fmt.Printf("%T,%v\n", x, x)

	fmt.Printf("%T,%v\n", y, y)

	fmt.Printf("%T,%v\n", x+y, x+y) //int16(x)+y dönüşümü burada hata vermemesinin nedeni type conversion yani tip dönüşümü yapıyor.
	//burada y int16 tipinde veri almış x ise bir veri tipi yok go bunu ihtiyacı olan veri tipine dönüştürme yapıyor yani toplama işlemini
	//yapabilmek için x si y nin veri tipine otomatik olarak atıyor çünkü şuanda x in sabit bir veri tipi yoktur.
	//kısaca const x sabiti, sabit bir veri almamıştır y gibi değildir. Bu nedenle istenen yere çağırma yapıldıgında oranın veri tipine otomatik olarak uyum sağlar.

	fmt.Printf("%T,%v\n", x, x) */

	/* const x = 5.2 + 4.8

	fmt.Printf("%T,%v\n", x, x) //burada sonuç float64,5.2 çıkar float64 çıkmasının nedeni const değerinin otomatik olarak tanımlanmasıdır.
	*/

	/* const x = int16(5.2+4.8) //burada dönüştürme işlemi yapılıyor.

	fmt.Printf("%T,%v\n", x, x) */

	/* //Aşağıdkai iki değişken x ve y typeless durumundadır yani veri tipi yoktur. Yani bizim tarafımızda atanmamıştır.
	const x = 3
	const y = 5.6

	fmt.Printf("%T,%v\n", x, x)

	fmt.Printf("%T,%v\n", y, y)

	fmt.Printf("%T,%v\n", x+y, x+y)//Yukarıdaki x ve y için veri tipi ataması yapsaydık bize hata verecekti.
	//Go kullandıgı tüm veri tipi bilgisine hakim olmak zorundadır.
	//Const - Sabit typeless şeklinde oluşturulabilir yani veri tipini kendisi atayabilir otomatik olarak
	*/

	const x = 3
	const y = 5.6
	const z = true
	const t = "test"

	fmt.Printf("%T,%v\n", x, x)

	fmt.Printf("%T,%v\n", y, y)
	fmt.Printf("%T,%v\n", z, z)

	fmt.Printf("%T,%v\n", t, t)

}
