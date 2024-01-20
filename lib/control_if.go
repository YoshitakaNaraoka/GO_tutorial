package lib

import "fmt"

func IfSeigyo() {
	// if文①
	age := 18
	if age < 0 {
		fmt.Println("生まれていません。")
	} else if age < 18 {
		fmt.Println("子どもです。")
	} else {
		fmt.Println("大人です。") // 出力
	}

	// if文②-1（かつ）　左の条件がfalseの場合、右の条件は無視される。
	kokugo := 100
	sugaku := 100
	if kokugo > 80 && sugaku > 80 {
		fmt.Println("合格") // 出力
	} else {
		fmt.Println("不合格")
	}
	// if文②-2（または）　左の条件がtrueの場合、右の条件は無視される。
	rika := 100
	syakai := 50
	if rika == 100 || syakai > 80 {
		fmt.Println("合格") // 出力
	} else {
		fmt.Println("不合格")
	}
	// if文②-3（ノットイコール）
	zukou := 1
	if zukou != 0 {
		fmt.Println("合格") // 出力
	} else {
		fmt.Println("不合格")
	}

	// if文③（変数宣言と条件分岐セットで利用する方法）
	m := map[string]int{"apple": 100}

	if value, ok := m["apple"]; ok {
		fmt.Println(value) // 100
	}
	fmt.Println(m) // map[apple:100]
	// fmt.Println(value, ok) スコープエラー
}