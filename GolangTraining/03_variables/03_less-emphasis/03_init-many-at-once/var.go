package main

import "fmt"

func main() {

	var message = "Hello World!"
	var a, b, c int = 1, 2, 3

	var d []string
	var f []string = []string{"33", "dsf", "dddd", "fff"}
	d = append(d, f...)
	f = append(f[0:2], f[3:]...)
	fmt.Println(message, a, b, c, d, f, f[0:0])
	d = append([]string{}, []string{}...)
	fmt.Println(message, a, b, c, d)

}
