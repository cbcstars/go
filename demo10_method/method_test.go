package demo10_method

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 方法是什么(Demo1)
/*
1:Go 方法是作用在接收者 (receiver) 上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。
2:接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。
  但是接收者不能是一个接口类型（参考第 11 章），因为接口是一个抽象定义，但是方法却是具体实现；如果这样做会引发一个编译错误：invalid receiver type...。
  最后接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。
3:一个类型加上它的方法等价于面向对象中的一个类。一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的。
4:类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的方法集 (method set)。
5:因为方法是函数，所以同样的，不允许方法重载，即对于一个类型只能有一个给定名称的方法。
  但是如果基于接收者类型，是有重载的：具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在，比如在同一个包里这么做是允许的：
		func (a *denseMatrix) Add(b Matrix) Matrix
		func (a *sparseMatrix) Add(b Matrix) Matrix
6:别名类型没有原始类型上已经定义过的方法。
7:定义方法的一般格式如下：
		func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
  在方法名之前，func 关键字之后的括号中指定 receiver。
  如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()。
  如果 recv 是一个指针，Go 会自动解引用。
8:如果方法不需要使用 recv 的值，可以用 _ 替换它，比如：
		func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
  recv 就像是面向对象语言中的 this 或 self，但是 Go 中并没有这两个关键字。随个人喜好，你可以使用 this 或 self 作为 receiver 的名字。
*/

// 如何为其他包中的类型定义方法(Demo2)
/*
1:类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在 int、float32(64) 或类似这些的类型上定义方法。
2:但是有一个间接的方式：
	a.可以先定义该类型（比如：int 或 float32(64)）的别名类型，然后再为别名类型定义方法。
	b.或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。(推荐，类似继承)
*/

// 指针或值作为接收者(Demo3)
/*
1:鉴于性能的原因，recv 最常见的是一个指向 receiver_type 的指针（因为我们不想要一个实例的拷贝，如果按值调用的话就会是这样），特别是在 receiver 类型是结构体时，就更是如此了。
2:如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法。否则，就在普通的值类型上定义方法。
3:go自动转换
	a:指针类型变量 调用 普通接收者方法，会自动解引用  t.Method() -> (*t).Method()
	b:普通类型变量 调用 指针接收者方法，会自动转换    t.Method() -> (&t).Method()
*/

// 方法和未导出字段(Demo4)
/*
1:这可以通过面向对象语言一个众所周知的技术来完成：提供 getter() 和 setter() 方法。对于 setter() 方法使用 Set... 前缀，对于 getter() 方法只使用成员名
*/

// 内嵌类型的方法和继承(Demo5)
/*
1:当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型。
  这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果，也类似 Ruby 中的混入 (mixin)。
2:内嵌将一个已存在类型的字段和方法注入到了另一个类型里：匿名字段上的方法“晋升”成为了外层类型的方法。当然类型可以有只作用于本身实例而不作用于内嵌“父”类型上的方法。
3:可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。
4:结构体内嵌和自己在同一个包中的结构体时，可以彼此访问对方所有的字段和方法。
*/

// 如何在类型中嵌入功能(Demo6)
/*
1:主要有两种方法来实现在类型中嵌入功能：
	a：聚合（或组合）：包含一个所需功能类型的具名字段。
	b：内嵌：内嵌（匿名地）所需功能类型。
2:如果内嵌类型嵌入了其他类型，也是可以的，那些类型的方法可以直接在外层类型中使用。
  因此一个好的策略是创建一些小的、可复用的类型作为一个工具箱，用于组成域类型
*/

// 多重继承(Demo7)
/*
多重继承指的是类型获得多个父类型行为的能力，它在传统的面向对象语言中通常是不被实现的（C++ 和 Python 例外）。
因为在类继承层次中，多重继承会给编译器引入额外的复杂度。但是在 Go 语言中，通过在类型中嵌入所有必要的父类型，可以很简单的实现多重继承。
*/

// 类型的 String() 方法和格式化描述符
/*
1:如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
  还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。
2:当你广泛使用一个自定义类型时，最好为它定义 String()方法。从上面的例子也可以看到，格式化描述符 %T 会给出类型的完全规格，
  %#v 会给出实例的完整输出，包括它的字段（在程序自动生成 Go 代码时也很有用）。
3:不要在 String() 方法里面调用涉及 String() 方法的方法，它会导致意料之外的错误，
  比如下面的例子，它导致了一个无限递归调用（TT.String() 调用 fmt.Sprintf，而 fmt.Sprintf 又会反过来调用 TT.String()），很快就会导致内存溢出
*/

// 垃圾回收和SetFinalizer
/*
1:Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，即垃圾收集器 (GC)，会处理这些事情，它搜索不再使用的变量然后释放它们的内存。
  可以通过 runtime 包访问 GC 进程。
2:通过调用 runtime.GC() 函数可以显式的触发 GC，但这只在某些罕见的场景下才有用，比如当内存资源不足时调用 runtime.GC()，
  它会在此函数执行的点上立即释放一大片内存，此时程序可能会有短时的性能下降（因为 GC 进程在执行）。
3:如果想知道当前的内存状态，可以使用：
	// fmt.Printf("%d\n", runtime.MemStats.Alloc/1024)
	// 此处代码在 Go 1.5.1下不再有效，更正为
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	上面的程序会给出已分配内存的总量，单位是 Kb。进一步的测量参考 文档页面。
4:如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：
	runtime.SetFinalizer(obj, func(obj *typeObj))
	func(obj *typeObj) 需要一个 typeObj 类型的指针参数 obj，特殊操作会在它上面执行。func 也可以是一个匿名函数。
	在对象被 GC 进程选中并从内存中移除以前，SetFinalizer 都不会执行，即使程序正常结束或者发生错误。
*/

// Demo1
type TwoInt struct {
	a, b int
}

func (ti *TwoInt) AddThem() int {
	return ti.a + ti.b
}

func (ti *TwoInt) AddToParam(p int) int {
	return ti.AddThem() + p
}

func TestMethod(t *testing.T) {
	ti := &TwoInt{1, 2}
	fmt.Println(ti.AddThem())
	fmt.Println(ti.AddToParam(3))

	twoInt := TwoInt{2, 3}
	fmt.Println(twoInt.AddThem())
	fmt.Println(twoInt.AddToParam(3))
}

type IntVector []int

func (iv IntVector) sum() (s int) {
	for _, v := range iv {
		s += v
	}
	return
}

func TestMethod2(t *testing.T) {
	iv := IntVector{1, 2, 3}
	fmt.Println(iv.sum())
}

// Demo2
type myTime struct {
	time.Time
}

func (mt myTime) first4Chars() string {
	return mt.Time.String()[0:4]
}

type myTime2 time.Time

func (mt myTime2) first4Chars() string {
	return time.Time(mt).String()
}

func TestMyTime(t *testing.T) {
	m := myTime{time.Now()}
	fmt.Println(m.first4Chars())
	fmt.Println(m.String())

	m2 := myTime2{}
	fmt.Println(m2.first4Chars())
	fmt.Println(time.Time(m2).String())
}

// Demo3:指针或值作为接受者
type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}

func (b B) change2() {
	b.thing = 2
}

func (b B) write() {
	fmt.Println(b.thing)
}

func TestB(t *testing.T) {
	b1 := B{}
	b1.change()
	b1.change2()
	b1.write()

	b2 := new(B)
	b2.change()
	b2.change2()
	b2.write()
}

// Demo4: 方法和未导出字段
type Ren struct {
	name string
}

func (r *Ren) Name() string {
	return r.name
}

func (r *Ren) SetName(name string) {
	r.name = name
}

func TestRen(t *testing.T) {
	r := new(Ren)
	r.SetName("小陈")
	fmt.Println(r.Name())
}

// Demo5:匿名类型的方法和继承
type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func (c *Car) GoToWorkIn() {
	c.Start()
	c.Stop()
	c.Engine.Start()
	c.Engine.Stop()
}

// Demo6:在类型中嵌入功能
type Log struct {
	msg string
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

type Customer struct {
	Name string
	log  *Log
}

func (c *Customer) Log() *Log {
	return c.log
}

type Customer2 struct {
	Name string
	Log
}

func (c *Customer2) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log.String())
}

func TestLog(t *testing.T) {
	// 组合
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Log().Add("2 - After me the world will be a better place!")
	//fmt.Println(c.log)
	fmt.Println(c.Log())

	// 内嵌
	c2 := &Customer2{"Barak Obama", Log{"1 - Yes we can!"}}
	c2.Add("2 - After me the world will be a better place!")
	fmt.Println(c2)
}

type As struct {
	A string
	Bs
}

func (as *As) MethodA() {

}

type Bs struct {
	B string
	Cs
}

func (bs *Bs) MethodB() {

}

type Cs struct {
	C string
}

func (cs *Cs) MethodC() {

}

func TestAS(t *testing.T) {
	as := new(As)
	fmt.Println(as.A, as.B, as.C)
	as.MethodA()
	as.MethodB()
	as.MethodC()
}

// Demo7:多重继承
type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func TestCameraPhone(t *testing.T) {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
}

// Demo8:类型的String()方法
type TwoInts struct {
	a int
	b int
}

func (t *TwoInts) String() string {
	return "(" + strconv.Itoa(t.a) + "/" + strconv.Itoa(t.b) + ")"
}

func TestString(t *testing.T) {
	ti := TwoInts{1, 2}
	fmt.Println(ti.String())
}
