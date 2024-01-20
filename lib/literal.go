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

	// 文字列（改行可能）
	fmt.Println(`abc
	def
	ghi`)

	// ブール
	fmt.Println(true)

	// nil（無効な参照先を意味する）
	fmt.Println(nil)

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