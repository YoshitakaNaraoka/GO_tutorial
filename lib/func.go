package lib

import (
	"fmt"
	"log"
	"os"
)

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
