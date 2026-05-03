// 1-) x = x - 10 vs x -=10

/* package main

import "fmt"

func main() {

	x := 50

	//x = x - 10 //assigment statement olarak tanımlanıyor. x değerine x -10 atanıyor. Burada önce çıkarma sonra atama yapılıyor

	//x -= 10 //yukarıdakinın kısa hali assignment operation deniyor. Burada ise atama ve çıkarma aynı anda yapılıyor.

	x += 10
	x *= 10
	x /= 10

	fmt.Printf("%T, %v\n", x, x)

} */

// 2-) K = F - 32 / 1.8 + 273   => -40 F kaç K derecedir?

/* package main

import "fmt"

func main() {

	F := -40
	K := float64(F-32)/1.8 + 273 //float64 ile dönüşüm yapıyoruz.

	fmt.Printf("%v,%T\n", K, K)

} */

// 3-)

/*

age := 40 const için değişken ataması yapıldıgında kod çalışmayacaktır. Bu nedenle atama yapmadan dogrudan const olarak vermek sorunu düzeltir.

const myAge = age

fmt.Printf("%v, %T\n", myAge, myAge)

*/

/* package main

import (
	"fmt"
)

func main() {

	const myAge = 40

	fmt.Printf("%v,%T\n", myAge, myAge)

} */

// 4-) Sabitlerde Shadowing Kavramı çalışır mı?

/* package main

import (
	"fmt"
)

//var x = 14 //Shadowing kavramı değişkenlerde func main dışında olursa ilk önce main içerisindeki çalıştırılır.

const x = 14

func main() {

	//var x = 24 main içerisindeki değişken

	const x = 24 //constlarda da ilk önce burası çalışır.

	fmt.Printf("%v,%T\n", x, x)

} */

// 5-) const x = 4, const y = 5.4,  x + y?

/* package main

import "fmt"

const x = 14

func main() {

	const x = 4 //typeless durumunda yani veri tipi belirlenmemiş

	const y = 5.4 //typeless bu da aynı şekilde

	fmt.Printf("%T,%v\n", x+y, x+y) //x in veri tipi float64

	fmt.Printf("%T,%v\n", x, x) //x in veri tipi int burada default yani varsılayan veri tipini veriyor go kendi

	fmt.Printf("%T,%v\n", y, y) // y nin veri tipi folat64

} */

// 6-) const x float64 = 6.4 , y := 4 + x, y? burada y nin veri tipini soruyor o da float64 olur x float64 oldugu için

package main

import "fmt"

func main() {

	const x float64 = 6.4

	y := 4 + x

	fmt.Printf("%T,%v\n", y, y)

}
