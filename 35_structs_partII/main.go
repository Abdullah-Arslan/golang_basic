package main

import (
	"fmt"
)

type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct { //Manager employee den farklı bir struct türüdür. Yani ikiside bir birinden farklıdır. yani ikiside aynı struct değildir. farklı veri tiplerini barındırırlar.
	employee  //bu kısımda yukarıda oluşturulan struct giriliyor yani onun içindeki verileri girmektense structını girerek işlem yapılıyor.
	hasDegree bool
}

// 15 A Relation -->Klasik OOP

func main() {

	e1 := employee{

		name:      "Gurcan",
		age:       40,
		isMarried: true,
	}
	fmt.Println(e1)

	/* e2 := e1
	fmt.Println(e2)
	e2.name="Arin"
	fmt.Println(e2)//e1 ve e2 struct lar bir biri arasındaki değerleri paylaşıyorlar.
	fmt.Println(e1) */

	/* 	m1 := manager{
		employee: employee{
			name:      "Ayşe",
			age:       28,
			isMarried: false,
		},
		hasDegree: true,
	} */

	m1 := manager{}
	m1.name = "Ayşe"
	m1.age = 28
	m1.isMarried = false
	m1.hasDegree = true

	fmt.Println(m1)

	// Anonim Struct

	theBoss := struct { //Bu strut yapısı sadece tek kullanımlık oldugu için func main dışında yazılmaz tek seferlik kullanım için bu şekilde yazılabilir.
		name  string
		money bool
	}{name: "THE BOOS", money: true}

	fmt.Println(theBoss)

}
