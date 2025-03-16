package main

import "fmt"

// Excel的最后一列编号是XFD，请问Excel总共多少列？
// 这是26进制
func main() {
	fmt.Printf("A = %d, Z = %d\n", 'A', 'Z')
	var base int = 'Z' - 'A' + 1 // 进制
	fmt.Println(base, "进制")
	// 总和
	var total int 
	total += 'D' - 'A' + 1
	total += base * ('F' - 'A' + 1)
	total += base * base * ('X' - 'A' + 1)
	fmt.Println("total", total)
}