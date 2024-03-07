package demo10_struct

import (
	"fmt"
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
