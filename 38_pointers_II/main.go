/* package main

import "fmt"

func main() { //GO pass by value yani bir fonksiyonu çağırırken ona gelen parametre argünanın bir kopyasıdır.

	x := 5
	fmt.Println(x)
	double(x)//burada yer alan kısımda aşağıda yazılan fonksiyon kısmında çarpım yapılıyor ve burada çağrılıyor.
	fmt.Println(x)

}

func double(num int) {//çarpma işlemini yukarıdaki x değerine göre yapılacak

	num *= 2
	fmt.Println(num)
}
*/

/* package main

import "fmt"

func main() { //Slice oluşturuyoruz.

	mySlc := []int{1, 10, 100}

	fmt.Println(mySlc)

	double(mySlc) //Burada double fonksiyonunu yazdırıyoruz fakat argüman olarak alınacak ve işlem yapılacak sayıları func main içerisindeki mySlc içerinden alıyor.

	fmt.Println(mySlc)
}

func double(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc)

} */

/* package main

import "fmt"

func main() { //Yukarıdaki Slice işleminin aynısını Array ile yapıyoruz.

	myArr := []int{1, 10, 100}

	fmt.Println(myArr)

	double([3]int(myArr))

	fmt.Println(myArr)

}

func double(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	fmt.Println(arr)

} */

package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x)
	double(&x) //Değeri pointera çevirmek için & işareti kullanılıyor.
	fmt.Println(x)

}

func double(num *int) { //Pointer bir veriyi oldugu yerde değiştirmek için kullanılır.
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)

}
