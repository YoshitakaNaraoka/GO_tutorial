package lib

import (
	"fmt"
)

func Array() {
	// 配列作成①
	var f_arr [3]float64 = [3]float64{.1, .2, 1.3}
	fmt.Println(f_arr) // [0.1 0.2 1.3]
	fmt.Println("配列の長さ：", len(f_arr))

	// 配列作成②（省略記法）
	s_arr := [3]string{"abc", "def", "ghi"}
	fmt.Println(s_arr[0], s_arr[1], s_arr[2]) // abc def ghi
	// fmt.Println(s_arr[3])　配列の要素数を超えたindexはエラー

	// 配列の要素変換
	s_arr[0] = "jkl"
	fmt.Println(s_arr) // [jkl def ghi]

	// スライス作成①
	i_slice := [...]int{1, 2, 3, 4, 5}
	fmt.Println(i_slice) // [1 2 3 4 5]
	fmt.Println("スライスの長さ：", len(i_slice))

	// スライス作成②（省略記法）
	i_slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(i_slice2) // [1 2 3 4 5]

	// スライス作成③（配列から作成）
	i_slice3 := f_arr[0:2] // indexが0から2の手前（1）までのアクセスでスライス作成
	fmt.Println(i_slice3)  // [0.1 0.2]

	// スライス作成④（スライスから作成）
	i_slice4 := i_slice[0:5] // indexが0から5の手前（4）までのアクセスでスライス作成
	fmt.Println(i_slice4)    // [1 2 3 4 5]

	// スライスの要素変換①
	i_slice[0] = 9
	fmt.Println(i_slice) // [9 2 3 4 5]

	// スライスの要素変換②（スライスから作成したスライスの要素変換）
	i_slice4[0] = 9
	fmt.Println(i_slice4) // [9 2 3 4 5]
	fmt.Println(i_slice)  // [9 2 3 4 5] 取得元のスライスも変換される

	// スライスの要素変換③（配列から作成したスライスの要素変換）
	i_slice5 := f_arr[0:3]
	fmt.Println(i_slice5) // [0.1 0.2 1.3]
	i_slice5[0] = 6
	fmt.Println(i_slice5) // [6 0.2 1.3]
	fmt.Println(f_arr)    // [6 0.2 1.3] 取得元の配列も変換される

	// 要素の追加（配列はできない）
	i_slice6 := []int{1, 2, 3}
	fmt.Println("配列の長さ：", len(i_slice6), i_slice6) //配列の長さ： 3 [1 2 3]
	i_slice6 = append(i_slice6, 4)
	fmt.Println("配列の長さ：", len(i_slice6), i_slice6) //配列の長さ： 4 [1 2 3 4]
	// append(i_slice6, 4)はエラー　（変数に代入すること）

}
/*
出力：
[0.1 0.2 1.3]
配列の長さ： 3
abc def ghi
[jkl def ghi]
[1 2 3 4 5]
スライスの長さ： 5
[1 2 3 4 5]
[0.1 0.2]
[1 2 3 4 5]
[9 2 3 4 5]
[9 2 3 4 5]
[9 2 3 4 5]
[0.1 0.2 1.3]
[6 0.2 1.3]
[6 0.2 1.3]
配列の長さ： 3 [1 2 3]
配列の長さ： 4 [1 2 3 4]
*/