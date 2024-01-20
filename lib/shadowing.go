package lib

import "fmt"

func Shadow() {
	// ①スコープ
	//  if文
	age := 10
	fmt.Println(age) // 10

	if true {
		fmt.Println(age) // 10
		age += 20        // ブロック外のageを参照し、代入
		fmt.Println(age) // 30
	}

	fmt.Println(age) // 30（ブロック外のageが更新されている）

	// for文
	num := 0
	fmt.Println(num) // 0
	i_slice := []int{1, 2, 3}
	for _, v := range i_slice { //_, vはシャドーイング
		num += v
	}
	fmt.Println(num) // 6

	// ②シャドーイング
	// if文
	age2 := 10
	fmt.Println(age2) // 10

	if true {
		age2 := 20        // ブロック外のageと同名の変数を定義（シャドーイング）
		fmt.Println(age2) // 20

		age2 += 30
		fmt.Println(age2) // 50
	}
	fmt.Println(age2) // 10（ブロック外のageは更新されない）

	// for文
	num2 := 0
	fmt.Println(num2) // 0
	i_slice2 := []int{1, 2, 3}
	for _, v := range i_slice2 { //_, vはシャドーイング
		num2 := 0 // シャドーイング
		num2 += v
	}
	fmt.Println(num2) // 0
}

/*
出力：
10
10
30
30
0
6
10
20
50
10
0
0
*/