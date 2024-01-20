package lib

import "fmt"

func ForSeigyo() {
	// for文①（基本形）  --条件式がtrueになるまで繰り返す--
	// 5回ループ
	num := 0
	for num < 5 {
		// num := 0 コメントを外すと無限ループしてしまう（同名の変数を再定義した時点からシャドーイングが発生）
		fmt.Println(num)
		num += 1
	}

	// for文②（変数宣言と条件分岐セットで利用する方法）  --if文とほぼ同じ方法--
	// 5回ループ
	for num := 0; num < 5; {
		fmt.Println(num)
		num += 1
	}

	// for文③（変数宣言と条件分岐、増減式セットで利用する方法）
	// 5回ループ
	for num := 0; num < 5; num++ {
		fmt.Println(num)
	}

	// for文④（無限ループ）  全てを省略した場合
	for {
		// break, returnを指定しない限り無限ループ
		break
	}

	// for文⑤（スライス・マップ要素の繰り返し）
	// スライスの要素を繰り返す①
	s_slice := []string{"apple", "banana", "lemon"}

	for i, v := range s_slice {
		fmt.Println(i, v) // インデックスと値を順に表示
	}
	// スライスの要素を繰り返す②（インデックスだけ取得）
	for i := range s_slice {
		fmt.Println(i)
	}
	// スライスの要素を繰り返す③（ブランク識別子を利用して、インデックスを無視）
	for _, v := range s_slice {
		fmt.Println(v) // _（ブランク識別子）を用いることで、未定義エラーにならない。
	}

	// マップの要素を繰り返す（上記スライスの操作とほぼ同じ）
	m := map[string]int{"apple": 100, "banana": 50, "lemon": 150}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// continueとbreak
	i_slice2 := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range i_slice2 {
		if v == 3 {
			continue
		} else if v == 5 {
			break
		}
		fmt.Println(v) // 1, 2, 4だけが出力される。
	}

}