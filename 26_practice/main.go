/* //FONKSİYONLAR İLE İLGİLİ UYGULAMALAR

// 1 -) Iki rakam arasında toplama, çıkarma ve çarpma
// işleminin yapıldığı bir fonkiyon yazınız.

package main

import "fmt"

func main() {
	x, y := 10, 4
	sum, dif, prod := calculation(x, y) //buradaki sum,dif,prod aşağıdaki ile aynı değil ve onların bir etkisi yok bunlar farklı isimlendirmede de olabilir.
	fmt.Println("Toplam:", sum)         //burada bu şekilde yazılmasının nedeni aşağıdaki kısımla aynı olması bir etki etmesinden değil, her fonksiyon kendi içerisindeki kapsam alanına etki eder.
	fmt.Println("Fark:", dif)			//önemli olan :=calculation değişkenin aşağıdaki fonksiyon içindeki ile birebir aynı olmasıdır.
	fmt.Println("Çarpım:", prod)

}

func calculation(num1, num2 int) (sum int, dif int, prod int) { //üç dönüş oldugu için kaç tane dönüş olacaksa o kadar yazılabilir.

	sum = num1 + num2 //burada dogrudan = ataması yapılmasının nedeni func calculation kısmında return (dif int gibi) kısmında yazıyorsa yapılır sadece yoksa := olarak atama gerekirdi.
	dif = num1 - num2
	prod = num1 * num2

	return sum, dif, prod

}
*/

//-----------------------------------------------------------------------------------------------------

/* // 2 -) Kullanıcı tarafından girilen nota göre geçtiniz
// veya kaldınız geri dönüşünü yazdırınız.
// Eğer not 50 den küçükse kaldı büyükse geçti şeklinde olacak kodu yazdırsın.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Lütfen Aldıgınız Notu Giriniz:")
	grade, _ := getGrade()

	var result string

	if grade >= 50 {
		result = "Geçtin"

	} else {
		result = "Kaldın"
	}

	fmt.Println(result)

}

func getGrade() (int, error) { //iki parametre aldı bunlardan birisi int diğeri de hatayı döndüren error
	reader := bufio.NewReader(os.Stdin) //Klavyeden girilen değeri okuması için os paketinden Stdin methodunu çağırıyoruz.
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	//alınan verinin başında ve sonundaki boşluklardan kurtulmak için input yazıyoruz.
	input = strings.TrimSpace(input)

	//gelen değer string şeklinde okunaca int değerine dönüşmesi lazım bunun için aşağıdaki kodu yazıyoruz.
	num, err := strconv.Atoi(input) //Atoi inputa girilen string değerini sayıya çevirecek.
	if err != nil {
		fmt.Println(err) //hata oldugunda bunu döndürecek
	}
	return num, nil //hata olmadıgında bu kısmı yazdıracak yani nil olarak kalacak

}
*/

//-----------------------------------------------------------------------------------------------------

// 3 -) 1 ile yüz arasındaki bir sayıyı tahmin etme uygulaması
// yazınız. Toplam tahmin hakkınız 10 olsun.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	target := numRand(1, 100)

	fmt.Println("1 ile 100 Arasındaki Sayıyı Bulmaya Çalışınız ")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ { //Burada sayı tahmin edilirken toplamda kalan hakkın tutulacağı yer
		fmt.Println(10-attempts, "Bu kadar hakkınız kaldı")
		fmt.Println("Lütfen tahminizi yazınız")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if num > target {
			fmt.Println("Tahmininiz daha büyük, daha küçük bir sayı giriniz.")
		} else if num < target {
			fmt.Println("Tahmininiz daha küçük, daha büyük bir sayı giriniz.")
		} else {
			fmt.Println("Doğru Tahmin, hedef sayı", target, " idi ", attempts, " seferde bulundunuz. ")
			break
		}

	}

}

func numRand(min, max int) int {

	rand.Seed(time.Now().Unix()) //Fonksiyon her çalıştırıldıgı yeni rakamın gelmesi için rand paketinden Seed methodu alınıyor.
	return rand.Intn(max-min) + min

}
