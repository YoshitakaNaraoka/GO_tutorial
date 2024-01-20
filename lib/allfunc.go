package lib

import (
	"fmt"
)

func Allfunc() {
	// ① 引数、戻り値なし
	Func1() // Hello

	// ①-2 return　（引数、戻り値なし）
	Func1_2()

	// ② 引数あり、戻り値なし
	Func2("身長：", "cm", 175) // 身長： 175 cm

	// ③ 引数あり（可変長引数）、戻り値なし
	Func3("5教科の得点:", 70, 70, 80, 36, 90)
	// 5教科の得点: [70 70 80 36 90]
	// 国語: 70

	// ④ 引数なし、戻り値あり
	hello := Func4()
	fmt.Println(hello) // Hello

	// ⑤　複数の戻り値を返す
	fmt.Println(Func5()) // Hello 100 true

	// ⑥　戻り値に名前をつける
	name, err := Func6()
	fmt.Println(name, err == nil) // Userさん true

	// ⑦　無名関数（名前付き関数と違って、関数の中でも定義できる）
	// その場で、利用する簡略関数
	f := func() string { return "Hello" }
	fmt.Println(f()) // Hello

	// ⑦-2 コールバック関数（関数の引数に渡される関数）
	// 変数で受け取る場合
	m := func(x, y int) int { return x + y }
	Func7_2(5, 5, m) // 10

	// その場で受け取る場合
	Func7_2(5, 5, func(x, y int) int { return x + y }) // 10

	// ⑧　別の関数を内部で実行（わざわざ、関数を引数で受け取らない）
	Func8() // Hello

	// ⑨　即時関数
	Func9 := func(txt string) string { return txt }("即実行")
	fmt.Println(Func9) // 即実行

	// ⑩ defer（関数を抜ける時に実行する。複数deferがある場合、後ろから実行）
	d := func() {
		println(3)
		defer println(4)
		return
		println(5) // returnより後ろは実行されない
	}

	println(1)
	defer println(2)
	d()
	defer println(6)
	defer println(7)
	// 出力順：1 3 4 7 6 2

	// ⑩-2 ファイル操作（ファイルを開いた後は、必ずクローズさせる　クローズの書き忘れ防止）
	// Func10_2()
}
