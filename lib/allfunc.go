package lib

import (
	"fmt"
	"log"
	"os"
)



func Allfunc() {
	// 引数、戻り値なし
	Func1() // Hello

	// return　（引数、戻り値なし）
	Func1_2()

	// 引数あり、戻り値なし
	Func2("身長：", "cm", 175) // 身長： 175 cm

	// 引数あり（可変長引数）、戻り値なし
	Func3("5教科の得点:", 70, 70, 80, 36, 90)
	// 5教科の得点: [70 70 80 36 90]
	// 国語: 70

	// 引数なし、戻り値あり
	hello := Func4()
	fmt.Println(hello) // Hello

	// 複数の戻り値を返す
	fmt.Println(Func5()) // Hello 100 true

	// 戻り値に名前をつける
	name, err := Func6()
	fmt.Println(name, err == nil) // Userさん true

	// 無名関数（名前付き関数と違って、関数の中でも定義できる）
	// その場で、利用する簡略関数
	f := func() string { return "Hello" }
	fmt.Println(f()) // Hello

	// コールバック関数（関数の引数に渡される関数）
	// 変数で受け取る場合
	m := func(x, y int) int { return x + y }
	Func7_2(5, 5, m) // 10

	// その場で受け取る場合
	Func7_2(5, 5, func(x, y int) int { return x + y }) // 10

	// 別の関数を内部で実行（わざわざ、関数を引数で受け取らない）
	Func8() // Hello

	// 即時関数
	Func9 := func(txt string) string { return txt }("即実行")
	fmt.Println(Func9) // 即実行

	// defer（関数を抜ける時に実行する。複数deferがある場合、後ろから実行）
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

func Func1() {
	fmt.Println("Hello")
}

func Func1_2() {
	return // 無を返して関数から抜ける
}

func Func2(txt, cm string, num int) {
	fmt.Println(txt, num, cm)
}

//　可変長引数はスライスとして扱うことができる
func Func3(txt string, score ...int) {
	fmt.Println(txt, score)
	fmt.Println("国語:", score[0])
}

// 以下2つはエラー
// func Func3(score ...int, txt string) {}  可変長引数は必ず最後
// func Func3(txt ...string, score ...int) {}  可変長引数は1つだけ

func Func4() string {
	return "Hello" // return で返して以後出力なし
}

// (複数の戻り値の場合、戻り値をカッコでくくる必要がある)
func Func5() (string, int, bool) {
	return "Hello", 100, true
}

// 名前をつけた変数をreturnする
func Func6() (name string, err error) {
	// 戻り値変数は、以下と同義で関数内で扱うことが可能
	// var name string
	// var err error

	name = "Userさん"

	return
}

// コールバック関数を受け取る関数
func Func7_2(x, y int, f func(int, int) int) {
	fmt.Println(f(x, y))
}

func Func8() {
	Func1()
}

func Func10_2() {
	fp, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
}

