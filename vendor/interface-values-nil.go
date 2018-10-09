package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()


	//Switch with interface
	//
	////var inter interface{} = 21
	////var inter interface{} = "Hello"
	//var inter interface{} = true
	//
	//switch v := inter.(type) {
	//case int:
	//	fmt.Printf("Twice %v is %v\n", v, v*2)
	//case string:
	//	fmt.Printf("%q is %v bytes long\n", v, len(v))
	//default:
	//	fmt.Printf("I don't know about type %T!\n", v)
	//}


}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
