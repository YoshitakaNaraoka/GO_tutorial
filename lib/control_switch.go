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