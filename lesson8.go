package main
import "fmt"
var (
	i   int     = 1
	f64 float64 = 1.2
	s   string  = "test"
	//var t bool = true
	//var f bool = false
	t, f bool = true, false
	)

	//varと違ってmain関数の外では宣言できない。
func foo() {
	xi := 1
	//再代入の際はコロンがいらない
	xi = 2
	xf64 := 1.2
	xs := "test"
	xt, xf := true,false
	fmt.Println(xi, xf64, xs, xt, xf)
	//f64の型を確認する
	fmt.Printf("%T\n",f64)
	//バックスラッシュで改行
	fmt.Printf("%T\n",xi)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}