package iteraciones

import (
	"fmt"
)

/*
	func Iterar() {
		for i := 0; i < 10; i++ {
			fmt.Print(i, " ")
		}
	}
*/
func Iterar() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

/*
	func IterarContinue() {
		for i := 0; i < 10; i++ {
		if i==5
		continue
		}
			fmt.Print(i, " ")
	}
*/
func IterarContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
/*func IterarBreak() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}*/
func IterarBreak() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
