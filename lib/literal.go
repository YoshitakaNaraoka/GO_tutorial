package lib

import (
	"fmt"
)

func Lit() {
	// 数字
	fmt.Println(123)

	// 浮動小数点数
	fmt.Println(1.23)

	// 文字列
	fmt.Println("abc")

	// 文字列は改行可能
	fmt.Println(`abc
	def
	ghi`)

	
	fmt.Println(true) // bool

	
	fmt.Println(nil) // nil（無効な参照先を意味する）

}
/*
出力：
123
1.23
abc
abc
        def
        ghi
true
<nil>
*/

// 変数宣言
func Var() {
	
	var num1 int = 123 // 整数型の変数

	
	var num2 = 456 // 右辺で型が決まるなら型名は省略可能

	
	num3 := 789 // 省略記法（関数内でのみ利用可能）

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