package lib

import (
	"fmt"
)

// 変数宣言
func Var() {
	// 整数型の変数
	var num1 int = 123

	// 右辺で型が決まるなら型名は省略可能
	var num2 = 456

	// 省略記法（関数内でのみ利用可能）
	num3 := 789

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
}
/*
出力：
123
456
789
*/