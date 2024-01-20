package lib

import (
	"fmt"
	"log"
	"os"
)

// ①
func Func1() {
	fmt.Println("Hello")
}

// ①-2
func Func1_2() {
	return               // 関数から抜ける
	fmt.Println("Hello") // 出力されない
}

// ②
func Func2(txt, cm string, num int) {
	fmt.Println(txt, num, cm)
}

// ③　可変長引数はスライスとして扱うことができる
func Func3(txt string, score ...int) {
	fmt.Println(txt, score)
	fmt.Println("国語:", score[0])
}

// 以下2つはエラー
// func Func3(score ...int, txt string) {}  可変長引数は必ず最後
// func Func3(txt ...string, score ...int) {}  可変長引数は1つだけ

// ④
func Func4() string {
	return "Hello"
}

// ⑤ (複数の戻り値の場合、戻り値をカッコでくくる必要がある)
func Func5() (string, int, bool) {
	return "Hello", 100, true
}

// ⑥ 名前をつけた変数をreturnする
func Func6() (name string, err error) {
	// 戻り値変数は、以下と同義で関数内で扱うことが可能
	// var name string
	// var err error

	name = "Userさん"

	return
}

// ⑦-2
// コールバック関数を受け取る関数
func Func7_2(x, y int, f func(int, int) int) {
	fmt.Println(f(x, y))
}

// ⑧
func Func8() {
	Func1()
}

// ⑩-2
func Func10_2() {
	fp, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
}
