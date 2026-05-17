/* package main

import "fmt"

func main() {
	merhaba("Arin", 6)Argüman ise fonksiyonu çalıştırırken kullandıgımız değişkenlere denir. Bu kısım argüman kısmı
}

func merhaba(name string, age int) { //Parametre fonksiyon yazma işlemine denir. Bu kısım parametre kısmı

	fmt.Printf()

}
*/

/* package main

import "fmt"

func main() {

	fmt.Println(result(45))

}

//Notunuzu giriyorsunuz 50 den büyükse geçiyorsunuz. Küçükse kalıyorsunuz.

func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
	} else {
		return "kaldınız"
	}
} */

/* package main

import "fmt"

func main() { //Bu ana main kısmı go çalışmaya başladığı yerdir.
	//buranın dışındaki yazılan butun fonksiyonlar ana main içerisinde çağırılarak çalıştırılabilir.
	merhaba("Arslan", 6)

	name := "Elis"
	age := 4
	fmt.Printf("Adım %s, yaşım %d\n", name, age)
}

func merhaba(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d\n", name, age)

}
*/

/* package main

import (
	"fmt"
	"time"
)

// Go da ve diğer progralama dillerinde bi işi iki şekilde yapabilirsiniz.
// Dilin çekirdeğin kullanılacak fonksiyonlar yazılmıştır bu hazır olanı kullanabilirsiniz
// Yada farklı paketlerdeki fonksiyonları kullanmaya method denir yani bu veri tipindeki fonksiyonlara method denir.
// Methodlar aynı zamanda bir fonksiyondur.
func main() { //Composite Veri tipi oldugu yani hazır.
	//Bizim hazırlayacağımız veri tipide olabilir.

	var x int = 10

	fmt.Println(x)

	//time.Time farklı veri tiplerinin bir arada kullanıldıgı fonksiyonlardır.
	//burada now fonksiyonu bizim yazdıgım bir fonksiyon olmadıgı için method oluyor. Yazılan kodu üzerine gelindiginde turu yazmakta, fonksiyon mu, type mı, yada struct mı
	var zaman time.Time = time.Now() //Şimdiki zamanı öğrenmek için yazılan bir kod
	//Now bir method dur. Üzerine geldigimizde fonksiyon oldugu görünür. Now, time paketine ait bir fonksiyon dur.
	//Methodlar aynı zamanda bir fonksiyondur.
	fmt.Println(zaman)

}
*/

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

//Ekrandan kullanıcının bir veri girilmesi isteniyor
//bu veri alınacak ekrana yazdırılacak

func main() {

	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz:") //bu girileni okumak için reader değişkeni yazılıyor
	reader := bufio.NewReader(os.Stdin)           //reader adında bir değişken oluşturduk. bufio ve os birer paketlerdir import altına eklenir.

	//Ekrana girilen değeri bastırmak için value değişkeni yazılıyor. Okuyup ekrana yazdırıyor.
	value, _ := reader.ReadString('\n') // _ blank indetifier - boş hata değişkeni atama go da bir değişken tanımladıktan sonra o değişken kesinlikle kullanılmalıdır. err hata değişkeni kullanmak istemiyorsak _ alt tire koyarak geçiyoruz.
	fmt.Println(value)//fmt pakettir, Println method dur. Biz yamadıgımız için methoddur.

} */

package main

import "fmt"

//Multiple Return - Birden fazla değeri dönüştürme

func main() {

	bolum, kalan := bolme(104, 5) //buradaki bolum ve kalan değişken isimleri farklı olabilir anlaşılır olması için alttaki func bolme fonksiyonu ile aynı olmalıdır.
	fmt.Println(bolum, kalan)

}

//104/5 -------> 20 kalan 4 20 ve 4 dönecek
//bu fonksiyon iki tane değer dönecek kalan ve bölen diye
func bolme(bolunen, bolen int) (bolum, kalan int) {

	bolum = bolunen / bolen
	kalan = bolunen % bolen

	return bolum, kalan

}
