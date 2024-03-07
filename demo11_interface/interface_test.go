package demo11_interface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

// 实战:Go的动态类型(Demo9)
/*
1:Go 没有类：数据（结构体或更一般的类型）和方法是一种松耦合的正交关系。
2:Go 中的接口跟 Java/C# 类似：都是必须提供一个指定方法集的实现。但是更加灵活通用：任何提供了接口方法实现代码的类型都隐式地实现了该接口，而不用显式地声明。
3:和其它语言相比，Go 是唯一结合了接口值，静态类型检查（是否该类型实现了某个接口），运行时动态转换的语言，并且不需要显式地声明类型是否满足某个接口。该特性允许我们在不改变已有的代码的情况下定义和使用新接口。
4:接收一个（或多个）接口类型作为参数的函数，其实参可以是任何实现了该接口的类型的变量。 实现了某个接口的类型可以被传给任何以此接口为参数的函数。
5:类似于 Python 和 Ruby 这类动态语言中的动态类型 (duck typing)；这意味着对象可以根据提供的方法被处理（例如，作为参数传递给函数），而忽略它们的实际类型：它们能做什么比它们是什么更重要。
*/

// 实战:动态方法调用(Demo10)
/*
1:当变量被赋值给一个接口类型的变量时，编译器会检查其是否实现了该接口的所有函数。
2:如果方法调用作用于像 interface{} 这样的“泛型”上，你可以通过类型断言（参见 11.3 节）来检查变量是否实现了相应接口。
*/

// 实战:接口的提取(Demo1)
/*
你不用提前设计出所有的接口；
整个设计可以持续演进，而不用废弃之前的决定。类型要实现某个接口，它本身不用改变，你只需要在这个类型上实现新的方法
*/

// 实战:接口的继承(Demo11)
/*
当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，这个类型就可以使用（另一个类型）所有的接口方法。
*/

// 实战:结构体、集合和高阶函数(Demo12)
/*
高阶函数，实际上也就是把函数作为定义所需方法（其他函数）的参数
*/

// 总结
/*
我们总结一下前面看到的：Go 没有类，而是松耦合的类型、方法对接口的实现。
OO 语言最重要的三个方面分别是：封装、继承和多态，在 Go 中它们是怎样表现的呢？
	封装（数据隐藏）：和别的 OO 语言有 4 个或更多的访问层次相比，Go 把它简化为了 2 层（参见 4.2 节的可见性规则）:
		1）包范围内的：通过标识符首字母小写，对象只在它所在的包内可见
		2）可导出的：通过标识符首字母大写，对象对所在包以外也可见
类型只拥有自己所在包中定义的方法。
	继承：用组合实现：内嵌一个（或多个）包含想要的行为（字段和方法）的类型；多重继承可以通过内嵌多个类型实现
	多态：用接口实现：某个类型的实例可以赋给它所实现的任意接口类型的变量。类型和接口是松耦合的，并且多重继承可以通过实现多个接口实现。
		           Go 接口不是 Java 和 C# 接口的变体，而且接口间是不相关的，并且是大规模编程和可适应的演进型设计的关键。
*/

// Demo1
// Shaper 图形
type Shaper interface {
	// Area 面积
	Area() float32
}

// TopologicalGenus 拓扑级
type TopologicalGenus interface {
	// Rank 等级
	Rank() int
}

// Square 正方形
type Square struct {
	// Side 边
	Side float32
}

func (s *Square) Area() float32 {
	return s.Side * s.Side
}

func (s *Square) Rank() int {
	return 1
}

// Rectangle 长方形
type Rectangle struct {
	// 长，宽
	Length, Width float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

func (r Rectangle) Rank() int {
	return 2
}

func TestShaper(t *testing.T) {
	squarePointer := &Square{Side: 4.5}
	rectangle := Rectangle{Length: 4.5, Width: 5.5}

	shapers := []Shaper{squarePointer, rectangle}
	for _, shaper := range shapers {
		area := shaper.Area()
		fmt.Printf("square area is %f \n", area)
	}

	topologicalGenus := []TopologicalGenus{squarePointer, rectangle}
	for _, topology := range topologicalGenus {
		fmt.Println(topology.Rank())
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

// Demo9:Go的动态类型
type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
}

func (b Bird) Quack() {
	fmt.Println("I am quacking")
}

func (b Bird) Walk() {
	fmt.Println("I am walking")
}

func TestDuck(t *testing.T) {
	b := Bird{}
	DuckDance(b)
}

// Demo10:动态方法调用
type xmlWriter interface {
	WriteXML(w io.Writer) error
}

func StreamXml(v any, w io.Writer) error {
	if xw, ok := v.(xmlWriter); ok {
		return xw.WriteXML(w)
	}
	return EncodeToXML(v, w)
}

func EncodeToXML(v any, w io.Writer) error {
	return nil
}

// Demo11:接口的继承
type Task struct {
	Command string
	*log.Logger
}

func NewTask(command string, logger *log.Logger) *Task {
	return &Task{command, logger}
}

// 当 log.Logger 实现了 Log() 方法后，Task 的实例 task 就可以调用该方法：
//task.Log()
//类型可以通过继承多个接口来提供像多重继承一样的特性：
//type ReaderWriter struct {
//	*io.Reader
//	*io.Writer
//}

// Demo12:结构体、集合和高阶函数
type Car struct {
	Module       string
	Manufacturer string
	BuildYear    int
}

type Cars []*Car

func (cs Cars) Process(f func(c *Car)) {
	for _, c := range cs {
		f(c)
	}
}

func (cs Cars) FindAll(f func(c *Car) bool) Cars {
	cars := make([]*Car, 0)

	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})

	return cars
}

func TestCar(t *testing.T) {
	cars := Cars{
		&Car{Module: "1", Manufacturer: "BMW", BuildYear: 2024},
		&Car{Module: "1", Manufacturer: "BYD", BuildYear: 2024},
	}

	all := cars.FindAll(func(c *Car) bool {
		return c.Manufacturer == "BYD" && c.BuildYear > 2020
	})

	marshal, _ := json.Marshal(all)
	fmt.Println(string(marshal))
}
