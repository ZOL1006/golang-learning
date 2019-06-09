package main

import "fmt"

func main() {
	var nums = [10]int {1, 2, 3}
	for index, num :=range nums {
		fmt.Println(index,num)
	}
	var kvs = map[string]string {"a": "aa"}
	for key, value :=range kvs {
		fmt.Println(key,value)
	}
	var sli = make(map[string]string)
	sli["b"] = "bb"
	sli["c"] = "bb"
	fmt.Printf("len=%d slice=%v\n", len(sli), sli)
	/*查看元素在集合中是否存在 */
	value, flag := sli [ "c" ] /*如果确定是真实的,则存在,否则不存在 */
	if (flag) {
		fmt.Println("ok 的首都是", flag)
		fmt.Println("b 的首都是", value)
	} else {
		fmt.Println("American 的首都不存在")
	}
}
