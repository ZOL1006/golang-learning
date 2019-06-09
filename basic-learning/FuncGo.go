package main

import "fmt"

func main() {
	fmt.Println(max(500, 100))
	fmt.Println(min(500, 100))

	for a := 0; a < 5; a++ {
		fmt.Println(a)
	}

	numbers := [4]int{1, 2, 3, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}

	/* 定义局部变量 */
	var i, j int

	for i=2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if(i % j ==0) {
				break; // 如果发现因子，则不是素数
			}
		}
		if(j > (i / j)) {
			fmt.Printf("%d  是素数\n", i);
		}
	}
}

func max(num1 int, num2 int) int {
	var res int
	if (num1 > num2){
		res = num1
	}else {
		res = num2
	}
	return res
}

func min(num3 int, num4 int) int {
	var res int
	if ( num3 < num4 ){
		res = num3
	}else {
		res = num4
	}
	return res
}
