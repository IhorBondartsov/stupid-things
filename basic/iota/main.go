package main

import "fmt"

const(
	x = iota  // x == 0
	y = iota  // y == 1
	z = iota  // z == 2
	w  // если значение константы не указано, ей присваивается значение идущей перед ней константы, следовательно здесь также получается w = iota.
	   // Поэтому w == 3, а в случаях y и x также можно было бы опустить "= iota".
)

const v = iota // так как iota встречает ключевое слово `const`, происходит сброс на 0, поэтому v = 0.

const (
	e, f, g = iota, iota, iota // e=0,f=0,g=0, значения iota одинаковы, так как находятся на одной строке.
)

func main() {
	fmt.Println(x,y,z,w)
	fmt.Println(v)
	fmt.Println(e,f,g)
}