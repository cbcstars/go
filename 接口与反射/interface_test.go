package 接口与反射

import (
	"bytes"
	"fmt"
	"testing"
)

// 接口格式
/*
type Namer interface {
	Method1(param_list) return_type
	Method2(param_list) return_type
	...
}
*/

// 接口命名规则
/*
1: 只包含一个方法的接口的名字由方法名加 er 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。
2: 还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）
*/

// 接口特性(Demo1、Demo2)
/*
1:类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
2:实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
3:一个类型可以实现多个接口。
4:接口类型可以包含一个实例的引用,该实例的类型实现了此接口（接口是动态类型）。
5:即使接口在类型之后才定义，二者处于不同的包中，被单独编译：只要类型实现了接口中的方法，它就实现了此接口。
*/

// 接口嵌套接口(Demo3)
/*
1:一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
*/

// 类型断言(Demo4)
/*
1:varI 必须是一个接口变量,可以使用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值
v := varI.(T)       // unchecked type assertion
2:类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是它不可能预见所有的可能性。如果转换在程序运行时失败会导致错误发生。更安全的方式是使用以下形式来进行类型断言
if v, ok := varI.(T); ok {  // checked type assertion
    Process(v)
    return
}
// varI is not of type T
*/

// 类型判断：type-switch（Demo5）
/*
1:接口变量的类型也可以使用一种特殊形式的 switch 来检测：type-switch
2:可以用 type-switch 进行运行时类型分析，但是在 type-switch 不允许有 fallthrough
*/

// 使用方法集与接口(Demo6)
/*
1:在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以根据具体类型 P 直接辨识的：
	a:指针方法可以通过指针调用
	b:值方法可以通过值调用
	c:接收者是值的方法可以通过指针调用，因为指针会首先被解引用
	d:接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。

2:Go 语言规范定义了接口方法集的调用规则：
	a:类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
	b:类型 T 的可调用方法集包含接受者为 T 的所有方法
	c:类型 T 的可调用方法集不包含接受者为 *T 的方法
*/

// 空接口(Demo7)
/*
1:空接口或者最小接口 不包含任何方法，它对实现不做任何要求：
	type Any interface {}
2:可以给一个空接口类型的变量 var val interface {} 赋任何类型的值。
*/

// 实战:使用Sorted排序(Demo8)

// Demo1
// Shaper 图形
type Shaper interface {
	// Area 面积
	Area() float32
}

// Square 正方形
type Square struct {
	// Side 边
	Side float32
}

func (s *Square) Area() float32 {
	return s.Side * s.Side
}

// Rectangle 长方形
type Rectangle struct {
	// 长，宽
	Length, Width float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

func TestShaper(t *testing.T) {
	squarePointer := &Square{Side: 4.5}
	rectangle := Rectangle{Length: 4.5, Width: 5.5}

	shapers := []Shaper{squarePointer, rectangle}
	for _, shaper := range shapers {
		area := shaper.Area()
		fmt.Printf("square area is %f \n", area)
	}
}

// Demo2: 所有实现了 valuable 接口的类型都可以用这个函数
type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (sp stockPosition) getValue() float32 {
	return sp.sharePrice * sp.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Println(asset.getValue())
}

func TestValuable(t *testing.T) {
	var asset valuable = stockPosition{ticker: "ESOP", sharePrice: 88.99, count: 10}
	showValue(asset)
	asset = car{make: "Ben", model: "BMW", price: 459800}
	showValue(asset)
}

// Demo3:接口嵌套接口
type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

type Word struct {
}

func (w *Word) Read(b bytes.Buffer) bool {
	return true
}

func (w *Word) Write(b bytes.Buffer) bool {
	return true
}

func (w *Word) Lock() {

}

func (w *Word) Unlock() {

}

func (w *Word) Close() {

}

func TestWord(t *testing.T) {
	w := new(Word)
	w.Read(bytes.Buffer{})
	w.Write(bytes.Buffer{})
	w.Lock()
	w.Unlock()
	w.Close()
}

// Demo4:类型断言
func TestShaper4(t *testing.T) {
	rectanglePtr := &Rectangle{Length: 4.5, Width: 3.5}
	var shaper Shaper
	shaper = rectanglePtr

	rectangle := shaper.(*Rectangle)
	fmt.Println(rectangle.Area(), rectangle.Length, rectangle.Width)

	if v, ok := shaper.(*Rectangle); ok {
		fmt.Println(v.Area())
	}
}

// Demo5:类型判断
func TestTypeSwitch(t *testing.T) {
	var shaper Shaper
	//var shaper Shaper = &Square{Side: 5}

	switch t := shaper.(type) {
	case *Square:
		fmt.Printf("type square %T with value %v \n", t, t)
	case *Rectangle:
		fmt.Printf("type rectangle %T with value %v \n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check \n")
	default:
		fmt.Printf("unexpected type %T \n", t)
	}

	switch shaper.(type) {
	case *Square:
		// pass
	case *Rectangle:
		// pass
	default:
		// pass
	}
}

// Demo6:接口方法集的调用规则
type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i < end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

type List []int

func (listPtr *List) Append(i int) {
	*listPtr = append(*listPtr, i)
}

func (listVal List) Len() int {
	return len(listVal)
}

func TestList(t *testing.T) {
	listVal := List{1, 2}
	listVal.Append(3)
	fmt.Println(listVal.Len())
	CountInto(&listVal, 5, 10)
	LongEnough(listVal)
	fmt.Println(listVal)

	listPtr := &List{1, 2}
	listPtr.Append(3)
	fmt.Println(listPtr.Len())
	CountInto(listPtr, 5, 10)
	LongEnough(listPtr)
	fmt.Println(listPtr)
}

// Demo7:空接口
type Any interface {
}

func TestAny(t *testing.T) {
	var a Any
	a = 1
	a = "string"
	a = struct {
		Name string
		Age  int
	}{"小陈", 18}
	fmt.Println(a)
}

func TestAny2(t *testing.T) {
	var a Any = "1234"
	if v, ok := a.(int); ok {
		fmt.Println(v)
	}
	fmt.Println("测试")
}

// Demo8:使用Sorter接口排序
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 冒泡排序
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i, i+1) {
				data.Swap(i, i+1)
			}
		}
	}
}

type IntArray []int

func (a IntArray) Len() int {
	return len(a)
}

func (a IntArray) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a IntArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func TestIntArray(t *testing.T) {
	intSlice := IntArray{3, 5, 6, 2, 1}
	Sort(intSlice)
	fmt.Println(intSlice)
}
