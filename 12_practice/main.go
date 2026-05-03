// 1 : int x, float64 y type conversion sample
/*
package main

import "fmt"

func main() {

	x := 75
	var y float64
	y = float64(x)// type conversion yani tü dönüştürme bu şekilde yapılıyor
	//dönüştürmek istenen değişken parantez içerisinde yazılıyor.

	fmt.Println(y)

}
*/

// 2 : multiple assing sample x, y = y, x

/* package main

import "fmt"

func main() {

	x := 5
	y := 10

	fmt.Println("X:", x, "Y:", y)

	x, y = y, x//burada yukarıdaki verilen değerlerin sadece yerleri değiştiriliyor.

	fmt.Println("X:", x, "Y", y)
} */

// 3 : non English variable names - İngilizce olmayan değişken adları

/* package main

import "fmt"

func main() {

	//go da farklı alfabede de değişkenler oluşturulabilir.

	名稱 := "Arin"
	年齡 := 40

	fmt.Println("Name:", 名稱, "Age:", 年齡)

} */

// 4 : shadowing kavramı? gölgeleme

/* package main

import "fmt"

func main() {

	x := 5

	if true {
		x = 10 //buradaki x değişkeni main altındaki x değişkenini gölgeler.
		//burada := şekliden atama yaparsak x yeni bir değer alır ama sadece = eşittir olarak atama yaparsak
		//main altındaki x in değerini değiştirir ve ona atama yapmış olur.
		x++
		fmt.Println(x)

	}

	fmt.Println(x)

}
*/

// 5 : 40 as a string - 40 rakamını string olarak yazıdırınız.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 65

	s := string(x)

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", s, s)

	//65 de sting bir ifade ile görmek istenirse
	y := strconv.Itoa(x)
	fmt.Printf("%v,%T\n", y, y)

}
