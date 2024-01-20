package lib

import "fmt"

func SwitchSeigyo() {
	score := 5
	switch score {
	case 1:
		fmt.Println("退学決定")
	case 2:
		fmt.Println("不可")
	case 3:
		fmt.Println("可")
	case 4:
		fmt.Println("良")
	case 5:
		fmt.Println("優良") // 出力
	default:
		fmt.Println("単位0")
	}

	txt := "running"
	switch txt {
	case "running":
		fmt.Println("実行中") // 出力
		fallthrough        // "running"の場合も、実行中　停止中　の2つを出力させる。
	case "stop":
		fmt.Println("停止中") // 出力
	default:
		fmt.Println("その他")
	}

	// カンマ区切りでcaseの条件を複数指定できる
	// result := 1　出力なし
	result := 90
	switch result {
	case 10, 20:
		fmt.Println("問題あり")
	case 50, 60:
		fmt.Println("通常通り")
	case 70, 80, 90, 100:
		fmt.Println("快適") // 出力
	}
}

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