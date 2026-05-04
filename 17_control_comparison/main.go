// Comparison operators Karşılaştırma Operatörleri - bir birleri ile karşılaştırılabilen değişkenler kullanılmalıdır.
// == Equal != Not equal
// < Less than > Greater than
// <= Less than or equal >= Greater than or equal

// Logical Operators Mantıksal İfadeler. - Mantıksal operatörler sadece boolen ifadeleri işleme alabilriz. ture yada false şeklinde yada bu ikisini karşılatırma şeklinde kullanılabilir.
// &&    conditional AND    p && q  is  "if p then q else false"
// ||    conditional OR     p || q  is  "if p then true else q"
// !     NOT                !p      is  "not p"

package main

import "fmt"

func main() {

	//x, y := 3, 7 // karşılaştırma kısmı

	//x, y := "a", "b"//buradaki string ifadeler desimal değer olarak 97 98 olarak alınır. Bunlarda bir birine eşit değildir.

	//x, y := 3, 5.0 //Aynı türdeki veri tipindeki değerler kıyaslanabilir yani karşılaştırılabilir.

	/* fmt.Printf("%T,%v\n", x == y, x == y)
	fmt.Printf("%T, %v\n", x == y, x == y)
	fmt.Printf("%T, %v\n", x != y, x != y)
	fmt.Printf("%T, %v\n", x < y, x < y)
	fmt.Printf("%T, %v\n", x > y, x > y)
	fmt.Printf("%T, %v\n", x >= y, x >= y)
	fmt.Printf("%T, %v\n", x <= y, x <= y) */

	/* x, y := 15, 20

	set1 := (x == y) //false
	set2 := (x < y)//true

	fmt.Printf("%T,%v\n", set1, set2)
	fmt.Printf("%T,%v\n", set1, set2)

		fmt.Printf("%T,%v\n"(set1 && set2), (set1 && set2))//false && true ---> false yani ve değerinde karşılatırmalardan birisi false olması durumunda sonuçta false olur
		//her iki koşulda true olması durumunda sonuçta true olur.
	*/

	//x, y := 15, 20

	//set1 := (x != y) //false
	//set1 := (x == y) //false

	//set2 := (x > y) //false

	set3 := false

	//fmt.Printf("%T,%v\n"(set1 && set2)  AND ----> sadece 2 durum true ---> true
	//fmt.Printf("%v\n", (set1 || set2)) //OR ---> sadece 2 durum false ---> false

	fmt.Printf("%v\n", (!set3)) //! not
}
