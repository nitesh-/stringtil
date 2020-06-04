package main

import (
	"fmt"
	"github.com/nitesh-/stringtil"
)

func main() {
	//var str string = stringtil.Substr("hello world", 0)

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==")
	fmt.Println("Reverse", stringtil.Reverse("Hello"))
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==")
	fmt.Println("Repeat", stringtil.Repeat("Hello", 2))
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==")
	fmt.Println("Ucfirst", stringtil.Ucfirst("hello"))
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==")
	fmt.Println("Ucwords", stringtil.Ucwords("hello world"))
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==")
	fmt.Println("Substr", stringtil.Substr("hello world", 1, 2))
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==")
	fmt.Println("ParseUrl", stringtil.ParseUrl("hello=world&world=hello"))
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==")
	fmt.Println("Strtr", stringtil.Strtr("Hilla Warld","ia","eo"))
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==")

}
