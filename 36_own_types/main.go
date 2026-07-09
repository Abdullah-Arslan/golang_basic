/* package main

import (
	"fmt"
	"strings"
)

/* //Defined Type kendi oluşturdugumuz veri tiplerine verilen adlardır.
type employee struct {//struct---> underlygin type employee ---> Defined Type, Name Type (bunlar tanımlı veri tipi denemktir.)
	name      string
	age       int
	isMarried bool
}
*/

//type mile float64 //Kendimiz bir veri tipi oluşturmak istiyorsak önce typle ile başlayacağız.

//type kilometre float64

//type mystring string

//func main() { */

/* 	e1 := employee{
   		name:      "Gurcan",
   		age:       40,
   		isMarried: true,
   	}

   	fmt.Println(e1)
*/

/* var m1 mile
m1 = 3.2 */
//fmt.Println(m1)
//fmt.Printf("%T, %v", m1, m1)

//fmt.Println()

//m2 := mile(4.6)

/* k1 := kilometre(7.8)
fmt.Printf("%T, %v", k1, 11)

fmt.Println(m1+mile(k1)) */

/* fmt.Println(m1 + m2)
fmt.Printf("%.2f", (m1 * m2))
fmt.Println()
fmt.Println(m1 + m2)

fmt.Println(m1 + m2 + 2.1)

//fmt.Println(m1+m2+"arin")//bunun yazdırılabilmesi için farklı bir method kullanılması gerekiyor. */

/* mystring := "arin"

fmt.Println(strings.ToUpper(mystring)) //sting bir değeri float64 ile yan yana toplayabilmek için bu yöntem kullanılıyor. */

/* f1 := float64(4.4)
fmt.Println()
fmt.Printf("%T, %v", (m1 + mile(f1)), (m1 + mile(f1))) //burada mile(f1) yaparak dönüştürme işlemi yapıyoruz ki işlem gerçekleşsin)
fmt.Println()
fmt.Printf("%T",%v", (float64(m1) + f1), (float64(m1) + f1)) */

//}

package main

import "fmt"

type mile float64
type kilometer float64

func main() {

	//m1=10, k=?
	m1 := mile(10)

	k1 := toKilometer(m1)

	fmt.Println(k1)

	//k2 =10, m2=?
	k2 := kilometer(10)

	m2 := toMile(k2)
	fmt.Println(m2)
}

func toKilometer(m mile) kilometer { //metreyi kilometreye çeviriyor.
	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile { //Kilometreyi mile çeviriyor.

	return mile(k * 0.62)

}
