package main

import "fmt"

var packVar = "Package Scope"

func main() {

	if true { //Köşeli parantezler scop sayılır açılan her yer için bu neden fmt kendi scopu içinde tanımlanmalı

		var blokVar = "Blok Scope"
		fmt.Println(blokVar)

	}

	var funcVar = "Fcunc Scope"

	fmt.Println(funcVar)
}
