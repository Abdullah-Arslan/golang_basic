/* package main

import (
	"errors" //bu kısım tanımlanan paketlerin yazıldıgı yer.
	"fmt"
)

func main() {

	//Go da hatalar err, Eror olarak tanımlanır. Go hatalara value olarak bakar.
	result, err := eventNum(7)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Girdiğiniz Sayı:", result)
	}

}

// Basit bir fonksiyon yazılıyor tek sayıları yada çift sayıları gösterir olacak
// (int) yazan kısım dönüş yapılacaksa onu yazıyoruz ve hata kısmıda dönecekse onuda int yanına atıyoruz.
// (int,eror) kısmında iki değer bize dönecek
func eventNum(num int) (int, error) {//error başlangıçta bir değeri yoksa ona nil verilir. Buradaki error nil olarak döner. Biz bir başlangıç değeri vermedigimiz için

	if num%2 != 0 {
		return 0, errors.New("HATA: Çift Sayı Girmediniz")
	}

	//bu kısımda herhangibir hata yoksa dönecek kısım onuda return ile dönüyoruz.
	return num, nil //buradaki nil ifadesi genel olarak başlangıç değerleri olmayan ifadelerdir.
	//herhangibir ifadenin özel bir değeri olmadıgı zaman ona nil olarak atama yapılabilir.
}
*/

/* // Bir sayının karekökünün alındıgı bir fonksiyon yazılacak.
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	//burada iki değer tanımlaması yapıyoruz bunun nedeni fonksiyonda iki değer dönmesinden
	result, err := sRoot(-5)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result)
	}

}

func sRoot(num float64) (float64, error) {
	//hata karaköklü ifadelerde hata çıkan sonuç sıfır dan küçükse döner
	if num < 0 {
		return 0, errors.New("Negatif Sayıların Karakökü Alınamadı")
	}

	//hata olmadıgı zaman bu işlemi yap
	return math.Sqrt(num), nil //nil pointer tapanlı veri tipleri için kullanılır. Eror da interface tabanlıdır.
}
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	//Bir dosyayı açmaya çalışacağız

	file, err := os.Open("test.txt") //os.Open, Open methodunun üzerine geldiğinde (*os.File, error) kısmında iki değer döndüğünü görüyoruz.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız", file)
	}

}
