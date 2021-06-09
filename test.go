package main

import "fmt"

// func main() {
// 	var str = "Hello World!"
// 	fmt.Println(str)
// }

// func main() {

// 	var a string = "initial"
// 	var b, c int = 1, 2
// 	var d = true
// 	var e int
// 	f := "short"

// 	fmt.Println(a)
// 	fmt.Println(b, c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// 	fmt.Println(f)

// }

// func main() {
// 	const s string = "this is a test sentence"
// 	fmt.Println(s)
// 	const n = 500000000
// 	const d = 3e20 / n
// 	fmt.Println((d))
// 	fmt.Println(int(d))
// 	fmt.Println(math.Sin(n))
// }

// func main() {

// 	i := 1
// 	for i <= 3 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}

// 	for j := 7; j <= 9; j++ {
// 		fmt.Println(j)
// 	}

// 	for {
// 		fmt.Println("loop")
// 		break
// 	}

// 	for n := 0; n <= 5; n++ {
// 		if n%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(n)
// 	}

// }

// func main() {

// 	if 7%2 == 0 {
// 		fmt.Println("7 is even")
// 	} else {
// 		fmt.Println("7 is odd")
// 	}

// 	if 8%4 == 0 {
// 		fmt.Println(("8 is divisible by 4"))
// 	}

// 	if num := 9; num < 0 {
// 		fmt.Println("is negative")
// 	} else if num < 10 {
// 		fmt.Println(num, "has 1 digit")
// 	} else {
// 		fmt.Println(num, "has multiple digits")
// 	}

// }

// func main() {

// 	i := 2
// 	fmt.Print("Write ", i, " as ")
// 	switch i {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	}

// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It's the weekend")
// 	default:
// 		fmt.Println("It's a weekday")
// 	}

// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("It's before noon")
// 	default:
// 		fmt.Println("It's after noon")
// 	}

// 	whatAmI := func(i interface{}) {
// 		switch t := i.(type) {
// 		case bool:
// 			fmt.Println("I'm a bool")
// 		case int:
// 			fmt.Println("I'm an int")
// 		default:
// 			fmt.Printf("Don't know type %T\n", t)
// 		}
// 	}

// 	whatAmI(true)
// 	whatAmI(1)
// 	whatAmI("hey")

// }

// func main() {

// 	// 通过make()函数创建切片，需要传入一个参数来指定切片的长度
// 	slice := make([]int, 3, 5) //3是切片长度，5是切片容量，切片长度不能大于切片容量

// 	// 通过字面量创建切片
// 	// 字符串切片
// 	//myStr := []string{"Jack", "Mark", "Nick"}

// 	// 整型切片
// 	myNum := []int{10, 20, 30, 40}

// 	// 使用空字符串初始化第100个元素
// 	myStr := []string{99: "100"}

// 	fmt.Println(myNum)
// 	fmt.Println(myStr)
// 	fmt.Println(slice)

// 	// []运算符内指定一个值，创建的是数组
// 	// 创建一个整型数组
// 	myArray := [5]int{10, 20, 30}
// 	fmt.Println(myArray)

// 	// []运算符内不指定值，创建的是切片
// 	// 创建一个整型切片
// 	mySlice := []int{10, 20, 30}
// 	fmt.Println(mySlice)

// 	// 创建nil整型切片
// 	var myNum1 []int
// 	fmt.Println(myNum1)

// 	// 使用make创建空的整型切片
// 	myNum2 := make([]int, 0)
// 	fmt.Println(myNum2)

// }

// func isSlash(r rune) bool {
// 	return r == '\\' || r == '/'
// }

// func main() {
// 	s := "test\\cd/d1231"
// 	ss := strings.FieldsFunc(s, isSlash)
// 	fmt.Printf("%q\n", ss)
// }

//    go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

// type Person struct {
// 	name string
// 	sex  string
// 	age  int
// }

// type Student struct {
// 	Person
// 	id   int
// 	addr string
// 	name string
// }

// func main() {
// 	var s Student
// 	s.name = "asdaksd"

// 	fmt.Println(s.name)
// 	// 初始化
// 	// s1 := Student{Person{"ctres", "man", 20}, 1, "bj"}
// 	// fmt.Println(s1)

// 	s2 := Student{Person: Person{"5lmh", "man", 20}}
// 	fmt.Println(s2)

// 	s3 := Student{Person: Person{name: "123123"}}
// 	fmt.Println(s3)
// }

//数组
// func create(a string, b int) int {
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	c := 100

// 	return c
// }

// var arr0 = [...]int{1, 2, 3, 4}
// var arr1 = [...]int{1: 100, 2: 100, 4: 200}
// var arr2 = [...]string{0: "this is the second", 10: "this is the fifth"}

// func main() {
// 	c := create("sadasd", 1231)
// 	fmt.Println(arr0)
// 	fmt.Println(len(arr2))
// 	fmt.Println(cap(arr2))
// 	fmt.Println(c)
// }

//多维数组遍历
// func main() {
// 	var arr [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
// 	for k1, v1 := range arr {
// 		for k2, v2 := range v1 {
// 			fmt.Println(k1 + k2)
// 			fmt.Println(v1)
// 			fmt.Println(v2)
// 			fmt.Printf("(%d,%d) = %d ", k1, k2, v2)
// 		}
// 		fmt.Println()
// 	}
// }

// func findStudentByID(a string, b int) string {

// 	student := [...]struct {
// 		name string
// 		age  int
// 	}{
// 		{"test", 10},
// 		{"test1", 20},
// 	}

// 	fmt.Println(student)
// 	return "123"
// }

// // func main() {
// // 	findStudentByID("123", 1)
// // }

// func myTest(a [5]int, target int) {
// 	// 遍历数组
// 	for i := 0; i < len(a); i++ {
// 		other := target - a[i]
// 		// 继续遍历
// 		for j := i + 1; j < len(a); j++ {
// 			if a[j] == other {
// 				fmt.Printf("(%d,%d)\n", i, j)
// 			}
// 		}
// 	}
// }

// func main() {
// 	b := [5]int{1, 3, 5, 8, 7}
// 	myTest(b, 8)

// 	findStudentByID("123", 1)
// }

// func main() {
// 	//1.声明切片
// 	var s1 []int
// 	if s1 == nil {
// 		fmt.Println("是空")
// 	} else {
// 		fmt.Println("不是空")
// 	}
// 	// 2.:=
// 	s2 := []int{}
// 	// 3.make()
// 	var s3 []int = make([]int, 0)
// 	fmt.Println(s1, s2, s3)
// 	// 4.初始化赋值
// 	var s4 []int = make([]int, 0, 0)
// 	fmt.Println(s4)
// 	s5 := []int{1, 2, 3}
// 	fmt.Println(s5)
// 	// 5.从数组切片
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	var s6 []int
// 	// 前包后不包
// 	s6 = arr[1:4]
// 	fmt.Println(s6)
// }

// Student 数据类型
// type Student struct {
// 	name  string
// 	age   int
// 	Class string
// }

// type student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	CLass string `json:"class"`
// }

// func newStu(name1 string, age1 int, class1 string) *Student {
// 	return &Student{name: name1, age: age1, Class: class1}
// }

// type intener struct {
// 	name string
// 	student
// 	age int
// }

// func main() {
// 	var stu Student
// 	stu.name = "add"
// 	stu.age = 18
// 	stu.Class = "level1"

// 	fmt.Println(stu.name)

// 	var stu1 = newStu("jack", 18, "level2")
// 	fmt.Println(stu1)

// 	var inn intener
// 	var stu111 = new(student)
// 	stu111.Name = "asdasd"
// 	inn.name = "new test"

// 	fmt.Println(stu111)
// }

// var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// var slice0 []int = arr[1:6]
// var slice1 []int = arr[:5]
// var slice2 []int = arr[2:]
// var slice3 []int = arr[:]
// var slice4 = arr[:len(arr)-1] //去掉切片的最后一个元素

// func main() {
// 	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
// 	slice5 := arr2[2:6]
// 	slice6 := arr2[:5]
// 	slice7 := arr2[2:]
// 	slice8 := arr2[:]
// 	slice9 := arr2[:len(arr2)-1] //去掉切片的最后一个元素

// 	for k1, v1 := range arr2 {
// 		fmt.Println(k1)
// 		fmt.Println((v1))
// 	}

// 	fmt.Println(slice5)
// 	fmt.Println(slice6)
// 	fmt.Println(slice7)
// 	fmt.Println(slice8)
// 	fmt.Println(slice9)
// 	fmt.Println(len(arr2))
// 	fmt.Println(cap(arr2))
// }

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
func main() {
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)
}
