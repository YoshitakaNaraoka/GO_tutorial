
package lib

import (
	"fmt"
	"reflect"
	"strconv"
)

func DataType() {

	i := 123
	f := 1.23
	s := "abc"
	b := true

	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(f, reflect.TypeOf(f))
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(b, reflect.TypeOf(b))

	/* ①数値同士のキャスト */
	var f2 = float64(i) // 123
	var i2 = int(f)     // 1

	fmt.Println(f2, reflect.TypeOf(f2))
	fmt.Println(i2, reflect.TypeOf(i2))

	/* ②数値型→ブール型 */
	var bool2 bool = i != 0 // true
	var bool3 bool = f != 0 // true
	fmt.Println(bool2, reflect.TypeOf(bool2))
	fmt.Println(bool3, reflect.TypeOf(bool3))

	/* ③Format系（各データ型から文字列型に変換）*/
	// 数値型→文字列型
	s2 := strconv.FormatInt(int64(i), 10)              // "123"
	s3 := strconv.FormatFloat(float64(f), 'E', -1, 64) // "1.23"
	fmt.Println(s2, reflect.TypeOf(s2))
	fmt.Println(s3, reflect.TypeOf(s3))

	// ブール型→文字列型
	s4 := strconv.FormatBool(b) // "true"
	fmt.Println(s4, reflect.TypeOf(s4))

	/* ④Parse系（文字列型から別のデータ型に変換）*/
	// 文字列型→数値型、ブール型
	i3, err := strconv.ParseInt("123", 10, 0) // 123
	fmt.Println(i3, reflect.TypeOf(i3), err, reflect.TypeOf(err))

	f4, err := strconv.ParseFloat("1.23", 64) // 1.23
	fmt.Println(f4, reflect.TypeOf(f4), err, reflect.TypeOf(err))

	b4, err := strconv.ParseBool("true") // true
	fmt.Println(b4, reflect.TypeOf(b4), err, reflect.TypeOf(err))

	// Parse系は、変換できなかった場合にエラーを変数に返す。（ここではerr変数で受け取っている）
	// b4, err := strconv.ParseBool("true")
	// fmt.Println(b4, reflect.TypeOf(b4), err, reflect.TypeOf(err))

}
/*
出力：
123 int
1.23 float64
abc string
true bool
123 float64
1 int
true bool
true bool
123 string
1.23E+00 string
true string
123 int64 <nil> <nil>
1.23 float64 <nil> <nil>
true bool <nil> <nil>
*/