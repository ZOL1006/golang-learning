package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {

	/* 数组 - 5 行 2 列*/
	var a = [5][2]int { {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	var i, j int

	/* 输出数组元素 */
	for  i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
		}
	}
	var aa int = 10

	fmt.Printf("变量的地址: %x\n", &aa)

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "com"})
}