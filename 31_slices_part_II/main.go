package main

import "fmt"

func main() {

	/* underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(underArray)

	mySlc := underArray[2:5] //buradaki kısmın açıklaması yukarıdaki underArray değişkeni içerisindeki 2 den başla 5 dahil olmayacak şekilde yazdır demektir.
	//burada 5 şi almıyor.
	fmt.Println(mySlc)

	mySlc2 := underArray[:6] //burada 0 dan balayarak 5 şe kadar alır 6 dahil değildir.
	fmt.Println(mySlc2)

	mySlc3 := underArray[3:] //burada 3 dahil son rakam olan 9 za kadar yazdırır.
	fmt.Println(mySlc3)

	mySlc4 := underArray[:] //başlangıç ve son için herhangibir bir sınırlama getirmeden hepsini yazdır demektir.
	fmt.Println(mySlc4)

	mySlc[0] = 100//buradaki atama diğer tüm slice larıda etkiliyor.
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	fmt.Println(mySlc3)
	fmt.Println(mySlc4) */

	/* myScl := []int{1, 2, 3}
	fmt.Println(myScl)

	/* myScl = append(myScl, 4, 5)//append methodu ile 4,5 rakamlarını ekledik.
	fmt.Println(myScl) */
	/*
		myScl2 := append(myScl, 4, 5)
		fmt.Println(myScl2)

		myScl[0] = 100
		fmt.Println(myScl)
		fmt.Println(myScl2) */

	/* mySlc:=[]int{1,2,3}
	mySlc2:=[]int{4,5}

	mySlc=append(mySlc, mySlc2...)//buradaki ... nokta mySlc2 de ki sayıları parçalayarak eklemiş oluyoruz.

	fmt.Println(mySlc)
	*/

	/* //Slice larda eleman silme
	myScl := []int{01, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(myScl)
	//myScl = myScl[2:]// ilk 2 elemanı silme işlemi bu şekilde yapılıyor.

	//myScl = myScl[:len(myScl)-3] //son 3 elemanı silmek için [:len(mySlc)-n]

	//baştaki ve sondaki elemanları silmek için bu şekilde kullanıyoruz.
	myScl = myScl[2:]
	myScl = myScl[:len(myScl)-3]
	fmt.Println(myScl)
	*/
	/*
		var myArr [4]int
		fmt.Println(myArr)

		var myScl []int
		myScl = make([]int, 4) //Zero Değerler Slice elemanlara ait 0 değerlerdir
		fmt.Println(myScl)

		var myScl2 []bool
		myScl2 = make([]bool, 4) //Zero Değerler Slice elemanlara ait 0 değerlerdir
		fmt.Println(myScl2) */

	var myScl3 []int
	fmt.Printf("%#v", myScl3)

	fmt.Println()

	myScl4 := make([]int, 0)
	fmt.Printf("%#v", myScl4)

}
