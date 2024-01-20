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