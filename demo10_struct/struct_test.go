package demo10_struct

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// 结构体定义(Demo1)
/*
结构体定义的一般方式如下：
type identifier struct {
    field1 type1
    field2 type2
    ...
}
type T struct {a, b int} 也是合法的语法，它更适用于简单的结构体。
结构体里的字段都有 名字，像 field1、field2 等，如果字段在代码中从来也不会被用到，那么可以命名它为 _。
结构体的字段可以是任何类型，甚至是结构体本身（参考第 10.5 节），也可以是函数或者接口（参考第 11 章）。可以声明结构体类型的一个变量，然后像下面这样给它的字段赋值：
	var s T
	s.a = 5
	s.b = 8
数组可以看作是一种结构体类型，不过它使用下标而不是具名的字段。
使用 new()
使用 new() 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)，如果需要可以把这条语句放在不同的行（比如定义是包范围的，但是分配却没有必要在开始就做）。
	var t *T
	t = new(T)
写这条语句的惯用方法是：t := new(T)，变量 t 是一个指向 T 的指针，此时结构体字段的值是它们所属类型的零值。
声明 var t T 也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型 T 。在这两种方式中，t 通常被称做类型 T 的一个实例 (instance) 或对象 (object)。
一个导出的结构体类型中有些字段是导出的(字段首字母大写)，另一些不是(字段首字母小写)，这是可能的。
*/

// 结构体初始化(Demo2)
/*
混合字面量语法 (composite literal syntax) &struct1{a, b, c} 是一种简写，底层仍然会调用 new()，这里值的顺序必须按照字段顺序来写。
在下面的例子中能看到可以通过在值的前面放上字段名来初始化字段的方式。表达式 new(Type) 和 &Type{} 是等价的。
*/

// 结构体的内存布局
/*
Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。
不像 Java 中的引用类型，一个对象和它里面包含的对象可能会在不同的内存空间中，这点和 Go 语言中的指针很像。
*/

// 结构体转换(Demo3)
/*
Go 中的类型转换遵循严格的规则。当为结构体定义了一个 alias 类型时，此结构体类型和它的 alias 类型都有相同的底层类型，
它们可以如示例 Demo3 那样互相转换，同时需要注意其中非法赋值或转换引起的编译错误
*/

// 结构体工厂(Demo4)
/*
Go 中实现 “构造子工厂”方法。为了方便通常会为类型定义一个工厂，按惯例，工厂的名字以 new... 或 New... 开头。
如果想知道结构体类型 T 的一个实例占用了多少内存，可以使用：size := unsafe.Sizeof(T{})。
*/

// 强制使用工厂方法(practice-Demo5)
/*
通过应用可见性规则就可以禁止使用 new() 函数，强制用户使用工厂方法，从而使类型变成私有的，就像在面向对象语言中那样。
*/

// map 和 struct vs new() 和 make() (Demo6)
/*
试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，但是 new() 一个 map 并试图向其填充数据，将会引发运行时错误！
因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。
*/

// 带标签的结构体(Demo7)
/*
1:结构体中的字段除了有名字和类型外，还可以有一个可选的标签 (tag)：它是一个附属于字段的字符串，可以是文档或其他的重要标记。
2:标签的内容不可以在一般的编程中使用，只有包 reflect 能获取它。
3:我们将在下一章（第 11.10 节中深入的探讨 reflect 包，它可以在运行时自省类型、属性和方法，
比如：在一个变量上调用 reflect.TypeOf() 可以获取变量的正确类型，如果变量是一个结构体类型，就可以通过 Field 来索引结构体的字段，然后就可以使用 Tag 属性。
*/

// 匿名字段(Demo8)
/*
1:结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。
2:可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。Go 语言中的继承是通过内嵌或组合来实现的，所以可以说，在 Go 语言中，相比较于继承，组合更受青睐。
3:在一个结构体中对于每一种数据类型只能有一个匿名字段。
*/

// 命名冲突(Demo9)
/*
10.5.3 命名冲突
当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？
	外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
	如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
例子：
	type A struct {a int}
	type B struct {a, b int}
	type C struct {A; B}
	var c C
规则 2：使用 c.a 是错误的，到底是 c.A.a 还是 c.B.a 呢？会导致编译器错误：ambiguous DOT reference c.a disambiguate with either c.A.a or c.B.a。

	type D struct {B; b float32}
	var d D
规则1：使用 d.b 是没问题的：它是 float32，而不是 B 的 b。如果想要内层的 b 可以通过 d.B.b 得到。
*/

// Demo1:结构体定义
type struct1 struct {
	i1  int
	f1  float32
	str string
}

func TestStruct(t *testing.T) {
	s := new(struct1)
	s.i1 = 1
	s.f1 = 3.14
	s.str = "hello struct"

	fmt.Println(s, s.i1, s.f1, s.str)
}

// Demo2:结构体初始化
type Interval struct {
	start int
	end   int
}

func TestInit(t *testing.T) {
	// pointer
	interval := new(Interval)
	interval = &Interval{start: 1, end: 2}
	// value
	interval2 := Interval{1, 2}
	interval2 = Interval{end: 2}
	fmt.Println(interval)
	fmt.Println(interval2)
}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func TestPerson(t *testing.T) {
	// 1-struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2—struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // 这是合法的
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}

// Demo3:结构体转换
type number struct {
	f float32
}

type nr number

func TestNumber(t *testing.T) {
	n1 := number{5.5}
	n2 := nr{5.5}
	fmt.Println(nr(n1))
	fmt.Println(number(n2))
}

// Demo4:结构体工厂
type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

// Demo6: map and struct -> new vs make
func TestMake(t *testing.T) {
	type MyMap map[string]any

	type Person struct {
		Name string
		Age  int
	}

	myMap := make(MyMap)
	myMap["hello"] = "world"
	fmt.Println(myMap)

	m := new(MyMap)
	(*m)["hello"] = "world" // panic: assignment to entry in nil map
	fmt.Println(m)

	// person := make(Person) // 编译错误

	p := new(Person)
	p.Name = "stars"
	fmt.Println(p)
}

// Demo7: 带标签的结构体
type TagType struct {
	field  int    `hello golang`
	field2 string `hello world`
	field3 bool   `how much`
}

func TestTag(t *testing.T) {
	tag := TagType{1, "test", true}
	value := reflect.TypeOf(tag)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fmt.Println(field.Tag)
	}
}

// Demo8:匿名字段
type innerS struct {
	int1 int
	int2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

func TestAnonymity(t *testing.T) {
	var o outerS
	o.b = 1
	o.c = 3.14
	o.int = 2
	o.int1 = 3
	o.int2 = 4

	fmt.Println(o, o.b, o.c, o.int, o.int1, o.int2, o.innerS, o.innerS.int1, o.innerS.int2)

	o2 := outerS{b: 1, c: 3.14, int: 2, innerS: innerS{int1: 3, int2: 4}}
	fmt.Println(o2)

	o3 := outerS{1, 3.14, 2, innerS{1, 2}}
	fmt.Println(o3)
}
