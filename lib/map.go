package lib

import "fmt"

func Map() {
	// マップの宣言
	var m map[string]int = map[string]int{}
	fmt.Println("マップの要素数:", len(m)) // マップの要素数: 0
	fmt.Println(m)                  // map[]

	// マップの宣言（省略版）
	m2 := map[string]int{}
	fmt.Println(m2) // map[]

	// マップの宣言（make()を利用）
	m3 := make(map[string]int)
	fmt.Println(m3) // map[]

	// マップの初期化（要素の末尾に「,」が無いとエラー）
	m4 := map[string]int{
		"apple":  100,
		"lemon":  150,
		"banana": 50,
	}
	fmt.Println("マップの要素数:", len(m4)) // マップの要素数: 3
	fmt.Println(m4)                  // map[apple:100 banana:50 lemon:150]

	// 要素の取得
	fmt.Println(m4["apple"])  // 100
	fmt.Println(m4["lemon"])  // 150
	fmt.Println(m4["banana"]) // 50
	// 存在しないキーを指定した場合、ゼロ値を返す（値がstringの場合空白文字）
	fmt.Println(m4["app"]) // 0

	// 要素の変換（代入）
	m4["apple"] = 500
	fmt.Println(m4) // map[apple:500 banana:50 lemon:150]

	// 要素の追加
	m4["orange"] = 200
	fmt.Println(m4) // map[apple:500 banana:50 lemon:150 orange:200]

	// 要素の削除
	delete(m4, "lemon")
	fmt.Println(m4) // map[apple:500 banana:50 orange:200]

	// 第２戻り値を受け取ることで、要素の存在チェックが可能
	// 要素の存在チェック（第１戻り値に値。第２戻り値に判定結果。）
	v, ok := m4["orange"]
	fmt.Println(v)  // 200
	fmt.Println(ok) // true

	// 要素の存在チェック（第１戻り値にゼロ値。第２戻り値に判定結果。）
	v2, ok2 := m4["orang"]
	fmt.Println(v2, ok2) // 0 false

	// 要素の存在チェック（if文で応用）
	if v, ok := m4["apple"]; ok {
		fmt.Println("appleキーの値:", v) // appleキーの値: 500
	}

	// マップの要素を繰り返す（range）
	for key, value := range m4 {
		fmt.Println("キー:", key, ", ", "値:", value)
	}

	// マップの要素を繰り返す（キーだけ、値だけも可能）
	for key := range m4 {
		fmt.Println("キー:", key)
	}

}

/*
出力：
マップの要素数: 0
map[]
map[]
map[]
マップの要素数: 3
map[apple:100 banana:50 lemon:150]
100
150
50
0
map[apple:500 banana:50 lemon:150]
map[apple:500 banana:50 lemon:150 orange:200]
map[apple:500 banana:50 orange:200]
200
true
0 false
appleキーの値: 500
キー: banana ,  値: 50
キー: orange ,  値: 200
キー: apple ,  値: 500
キー: apple
キー: banana
キー: orange
*/