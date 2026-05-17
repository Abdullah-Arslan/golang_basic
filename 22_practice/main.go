// 1-) 1 ile 10 arasındaki sayıları if yapısıyla tek - çift olarak yazdırınız.

/* package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Çifttir")
		} else {
			fmt.Println(i, "Tektir")
		}
	}

} */

// 2-) for yapısını kullanarak Go'da olmayan while döngüsüne örnek veriniz.

/* package main

import "fmt"

func main() {
	//go da while döngüsü yok bunu for ile yapabiliriz.

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++ //bu yazılmassa yani post olmassa sonsuza kadar yazdırır.
	}

} */

// 3-) Switch fallthrough ifadesini açıklayınız.
/* package main

import "fmt"

func main() {

	switch x := 25; {
	case x < 20:
		fmt.Printf("%d küçüktür 20\n", x)
		fallthrough //koşul sağlandıgında case biter ama fallthrouhg ile diğer koşullara da bakar bu işlemi yapar sadece.

	case x < 50:
		fmt.Printf("%d küçüktür 50\n", x)
		fallthrough

	case x < 100:
		fmt.Printf("%d küçüktür 100\n", x)
		fallthrough

	case x < 200:
		fmt.Printf("%d küçüktür 200\n", x)
	}

} */

// 4-) Aşağıdaki if döngüsünü daha idiomatic hale getiriniz.
//idiomatic kod yazımı go da vardır. Anlamı parçaların bir birinden daha bagımsız, daha sadecek kod yazımı
//ve daha anlaşılır kod yazımı demektir.
//idiomatic if else yapısını daha kısa ve az kullanmak da demektir.
/* package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Çifttir")
		} else {
			fmt.Println(i, "Tektir")
		}
	}

} */

//BURADAKİ KISIM DAHA İDİOMATİC OLARAK YAZILMIŞTIR.
/* package main

import "fmt"

func main() {

	x := 20
	if x%2 == 0 {
		fmt.Println(x, "çifttir")
		return
	}
	fmt.Println(x, "tektir")

}
*/

// 5-) 1 ile 50 arasındaki asal sayıları gösteren bir program yazınız.
package main

import "fmt"

func main() {

	var x, y int
	//asal sayılar 2 den başlar. Asal sayılar kendisine ve 1 re bölünür.
	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ { //parantez içi x in y ye bölümünden y küçük olmalıdır.
			if x%y == 0 {
				break
			}

		}

		if y > (x / y) {
			fmt.Printf("%d bir asal sayıdır.\n", x)
		}
	}

}
