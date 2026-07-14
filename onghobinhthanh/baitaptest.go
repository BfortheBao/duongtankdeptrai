git@github.com:BfortheBao/duongtankdeptrai.giipackage bai_6

import "fmt"

func Bai6() {
	for n := 2; n <= 9; n++ {
		fmt.Printf("Bảng cửu chương %d\n", n)
		for i := 1; i <= 10; i++ {
			if n*i%2 == 0 {
				fmt.Printf("%d x %d = %d\n", n, i, n*i)
			} else {
				continue
			}
		}
	}
}
t
