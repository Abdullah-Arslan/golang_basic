/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//fmt bir pakettir.  os,bufio,fmt bunlar pakettir. Bunların altında olan kısımlar Print gibi methodlar hazır olarak alınır.

	fmt.Print("Lütfen Bir Sayı Giriniz:")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') //_ blank identifier
	fmt.Println(value)

}
*/

/* package main

import "fmt"

func main(){

	//Neden paketlerin kullanımına ihtiyaç duyuyoruz? Paketler toplu olarak sunulsa daha iyi olmazmıydı?
	//Böyle bir paketleme yani genel bir paket durumu söz konusu değildir. Bunun ana nedeni ihtiyaç duyulmayan başka paketleride yazılan programa çekmiş oluruz ki buda gereksiz kasmalara ve gerek olmayan paket çekimine neden olur.
	//Kodlar bir çok paket ile sunulmasının avantajı da bakım kolaylılıgdır. Kodlarda sorun çıkması durumunda dogrudan sorun çıkan pakete giderek sorun çok kolay çözülebilir. Sorun daha hızlı bir şekilde çözülür.
	//Paketleme sistemi ile bir kod tekrar tekrar bir kez tanımladıktan sonra kullanılabilir. Program daha modüler olarak kullanılabilir.
	//Bir paket tanımlamasında her bir paket kendi içinde değerlendirilmelidir.
	//Golang sitesinde paketler mevcut buradan paketlerin tamamını görebiliriz. Burada kendimize uygun paketler seçilerek kullanılabilir.
	//Bir programlama dilinde mevcut paketler ile daha az kod yazarak istenen program daha hızlı olarak yazılabilir bu durum bir programlama dilinin en güçlü yönünü oluşturur.
	//Hazır paketler sıfırdan herşeyi sıfırdan yazamak yerine hazır paketleri kullanıyoruz.

	fmt.Println("Arin Yazılım")
} */
/*
// Buradaki kısım golang.org sitesinden yani Go nun kendi sitesi üzerinden alınmıştır aşağıdaki paket örneği
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))

	fmt.Println(strings.Count("animatrix", "a"))
}
*/

// Buradaki kısım golang.org sitesinden yani Go nun kendi sitesi üzerinden alınmıştır aşağıdaki paket örneği
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Gopher"))
}
