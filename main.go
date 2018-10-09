package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	_ "golang.org/x/tour/wc"
	"math"
	"strings"
)

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int  {
	return x*10 + 1
}

func needFloat(x float64) float64  {
	return  x*0.1
}
func main() {

	//fmt.Println(needInt(Small))
	//fmt.Println(needFloat(Small))
	//fmt.Println(needFloat(Big))
	//
	////For loop
	//sum := 0
	//for i :=0; i<10 ; i++  {
	//	sum += i
	//}
	//
	//fmt.Println("Sum is ",sum)
	//
	////Special case for loop
	//sum = 1
	//for sum < 1000	 {
	//	sum += sum
	//}
	//fmt.Println(sum)
	//
	//// If - condition
	//fmt.Println(sqrt(2),sqrt(-4))
	//
	//// If else - condition
	//fmt.Println(
	//	pow(3,2,10),
	//	pow(3,3,20),
	//	)

	//fmt.Println("-------------------------------")


	//Exercise - give x find z where z^2 = x
	//fmt.Println(sqrt_Newton_method(8.0))

	//fmt.Println("-------------------------------")

	// Switch condition
	//fmt.Print("Go runs on ")
	//switch os:= runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux")
	//default:
	//	fmt.Printf("%s.", os)
	//}

	//fmt.Println("-------------------------------")
	//Switch example
	//fmt.Println("When's Saturday?")
	//today := time.Now().Weekday()
	//switch time.Saturday {
	//case today + 0:
	//	fmt.Println("Today.")
	//case today + 1:
	//	fmt.Println("Tomorrow")
	//case today + 2:
	//	fmt.Println("Two days")
	//default:
	//	fmt.Println("Too far away")
	//
	//}

	//Switch without condition
	//t := time.Now()
	//switch  {
	//case t.Hour() <12:
	//	fmt.Println("Good morning")
	//case t.Hour() < 17:
	//	fmt.Println("Goood afternoon")
	//default:
	//	fmt.Println("Good evening")
	//}

	//Defer statement
	//The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns

	//defer fmt.Println("world")
	//
	//fmt.Println("Hello")

	//Result: Hello world



	//Deferred function calls are pushed onto a stack
	//fmt.Println("counting")
	//
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//
	//fmt.Println("done")

	//Pointer Example
	//i,j := 42,2701
	//
	//p := &i
	//fmt.Println(*p)
	//*p = 21
	//fmt.Println(i)
	//
	//p = &j
	//*p = *p / 37
	//fmt.Println(j)


	//Struct example
	//Can declare inside our outside function
	//
	//v := Vertex{3,5}
	//
	//// Create pointer to struct
	//p := &v
	//p.X = 1e9
	//fmt.Println(p.X)


	//
	//Slice example
	//When you change slice, the array will change
	//names := [4]string{
	//	"John",
	//	"Paul",
	//	"George",
	//	"Ringo",
	//}
	//fmt.Println(names)
	//
	//a := names[0:2]
	//b := names[1:3]
	//fmt.Println(a, b)
	//
	//b[0] = "XXX"
	//fmt.Println(a, b)
	//fmt.Println(names)

	//Slice length and capacity example
	//s := []int{2, 3, 5, 7, 11, 13}
	//printSlice(s)
	//
	//// Slice the slice to give it zero length.
	//s = s[:0]
	//printSlice(s)
	//
	//// Extend its length.
	//s = s[:4]
	//printSlice(s)
	//
	//// Drop its first two values.
	//s = s[2:]
	//printSlice(s)

	//Create slide with make
	//a := make([]int, 5)
	//printSlice2("a", a)
	//
	//
	//b := make([]int,2,8)
	//printSlice2("b", b)


	//// Create a tic-tac-toe board.
	//board := [][]string{
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//}
	//
	//// The players take turns.
	//board[0][0] = "X"
	//board[2][2] = "O"
	//board[1][2] = "X"
	//board[1][0] = "O"
	//board[0][2] = "X"
	//
	//for i := 0; i < len(board); i++ {
	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
	//}


	//Append slice use append method
	//var s []int
	//printSlice(s)
	//
	//// append works on nil slices.
	//s = append(s, 0)
	//printSlice(s)
	//
	//// We can add more than one element at a time.
	//s = append(s, 2, 3, 4)
	//printSlice(s)

	////Range in for loop
	//When ranging over a slice,
	//two values are returned for each iteration.
	//	The first is the index, and
	//the second is a copy of the element at that index.
	//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//for i, v := range pow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}


	//If you want to ignore index or value in range use this
	//pow := make([]int, 10)
	//for i := range pow {
	//	pow[i] = 1 << uint(i) // equals to == 2**i
	//}
	//for _, value := range pow {
	//	fmt.Printf("%d\n", value)
	//}


	//Slices of slices Exercise
	//Important
	//pic.Show(Pic)


	//Map Examples
	//type Geo struct {
	//	Lat,Long float64
	//}
	//
	//var maps = map[string] Geo{
	//	"Bell Labs": {40.68433, -74.39967,},
	//		"Google": {37.42202, -122.08408,},
	//}
	////maps = make(map[string]Geo)
	//maps["HCM"] = Geo{
	//	40.68433, -74.39967,
	//}
	//
	//
	////fmt.Println(maps["HCM"])
	//fmt.Println(maps)
	//


	//Check if a key is present in maps
	//m := make(map[string]int)
	//
	//m["Answer"] = 42
	//fmt.Println("The value:", m["Answer"])
	//
	//m["Answer"] = 48
	//fmt.Println("The value:", m["Answer"])
	//
	//delete(m, "Answer")
	//fmt.Println("The value:", m["Answer"])
	//
	//v, ok := m["Answer"]
	//fmt.Println("The value:", v, "Present?", ok)
	////If key is in m, ok is true. If not, ok is false.
	//
	////If key is not in the map,
	//// then elem is the zero value for the map's element type.


	//Exercise Maps
	//Need to see solution
	wc.Test(WordCount)


	fmt.Println("-------------------------------")


	//fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//
	//temp := strings.Fields("  foo bar  baz   ")
	//sort.Strings(temp)
	//fmt.Println(temp)
}
func WordCount(s string) map[string]int {

	//Old version , sort then count
	//This is not optimize
	//Don't use this
	//
	//temp := strings.Fields(s)
	//sort.Strings(temp)
	//
	//maps := make(map[string]int)
	//
	//cur := temp[0]
	//count := 0
	//for index:=0;index< len(temp);index++{
	//	if temp[index] == cur {
	//		count++
	//	} else {
	//		maps[cur] = count
	//
	//		cur = temp[index]
	//		count = 1
	//	}
	//
	//	maps[cur] = count
	//}
	//
	////return map[string]int{"x": 1}
	//return maps

	//New version
	//Use this
	maps := make(map[string]int)
	for _,value :=range strings.Fields(s){
		maps[value]++
	}
	return maps



}
func Pic(dx,dy int) [][]uint8  {
	myPic := make([][]uint8,dy)
	for y := range myPic {
		// y is index in myPic
		myPic[y] = make([]uint8, dx)

		for x := range myPic[y] {
			//x is index in myPic
			//myPic[y][x] = uint8( (x+y)/2 )
			myPic[y][x] = uint8( x*y )
			//myPic[y][x] = uint8( x^y )

		}
	}
	return myPic
}

func imageFun(x, y int) uint8 {
	return uint8(x) ^ uint8(y)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


type Vertex struct {
	X int
	Y int
}

func sqrt_Newton_method(x float64) float64  {
	z := 1.0
	for z<x {
		old := z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
		if math.Abs(old - z) < 0.0000000001 {
			break
		}
	}
	return z
}

func sqrt(x float64) string  {
	if x<0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x,n,lim float64) float64  {
	if v:= math.Pow(x,n); v<lim{
		return v
	} else {
		fmt.Printf("%g >= %g \n",v,lim)
	}
	return lim
}