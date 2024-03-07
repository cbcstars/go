package demo11_reflect

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
)

// 方法和类型的反射(Demo1)
/*
1:变量的最基本信息就是类型和值：反射包的 Type 用来表示一个 Go 类型，反射包的 Value 为 Go 值提供了反射接口。
2:两个简单的函数，reflect.TypeOf 和 reflect.ValueOf，返回被检查对象的类型和值。例如，x 被定义为：var x float64 = 3.4，那么 reflect.TypeOf(x) 返回 float64，reflect.ValueOf(x) 返回 <float64 Value>
	实际上，反射是通过检查一个接口的值，变量首先被转换成空接口。这从下面两个函数签名能够很明显的看出来：
		func TypeOf(i interface{}) Type
		func ValueOf(i interface{}) Value
*/

// 通过反射修改值(Demo2)
/*
1:当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 v = reflect.ValueOf(&x)。
2:通过 Type() 我们看到 v 现在的类型是 *float64 并且仍然是不可设置的。
3:要想让其可设置我们需要使用 Elem() 函数，这间接地使用指针：v = v.Elem()
4:现在 v.CanSet() 返回 true 并且 v.SetFloat(3.1415) 设置成功了！
*/

// 反射结构(Demo3)
/*
1:有些时候需要反射一个结构类型。NumField() 方法返回结构内的字段数量；通过一个 for 循环用索引取得每个字段的值 Field(i)。
2:我们同样能够调用签名在结构上的方法，例如，使用索引 n 来调用：Method(n).Call(nil)。
3:结构体中只有被导出字段（首字母大写）才可以设置值。
*/

// Printf和反射(Demo4)
/*
Printf() 中的 ... 参数为空接口类型。Printf() 使用反射包来解析这个参数列表。所以，Printf() 能够知道它每个参数的类型。
因此格式化字符串中只有 %d 而没有 %u 和 %ld，因为它知道这个参数是 unsigned 还是 long。这也是为什么 Print() 和 Println() 在没有格式字符串的情况下还能如此漂亮地输出。
*/

// Demo1: Kind()总是返回底层类型
func TestReflect(t *testing.T) {
	type MyInt int
	var i MyInt
	typeOf := reflect.TypeOf(i)
	fmt.Println("type:", typeOf.Name())
	fmt.Println("under type:", typeOf.Kind())

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}

// Demo2: 通过反射设置值
func TestSetValue(t *testing.T) {
	var v float64 = 3.4
	vRef := reflect.ValueOf(v)
	//vRef.SetFloat(3.1415) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Println("can set:", vRef.CanSet())

	vRef = reflect.ValueOf(&v)
	fmt.Println("type of v:", vRef.Type())
	fmt.Println("can set:", vRef.CanSet())
	vRef = vRef.Elem()
	fmt.Println("elem of v:", vRef.Type())
	fmt.Println("can set:", vRef.CanSet())
	vRef.SetFloat(3.1415)
	fmt.Println(vRef.Float())
}

// Demo3: 反射结构体
type Person struct {
	Name string
	age  int
}

func (p Person) String() {
	fmt.Printf("name:%v, age:%v \n", p.Name, p.age)
}

var p = Person{Name: "小陈", age: 28}

func TestReflectStruct(t *testing.T) {
	refT := reflect.TypeOf(p)
	refV := reflect.ValueOf(p)

	fmt.Println("kind type:", refT.Kind())
	fmt.Println("value:", refV)

	for i := 0; i < refV.NumField(); i++ {
		field := refV.Field(i)
		fmt.Printf("field type: %v, value: %v \n", field.Type(), field)
	}

	refV.Method(0).Call(nil)

	// 只有导出的字段可以修改值
	elem := reflect.ValueOf(&p).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if field.CanSet() {
			field.SetString("1234")
		}
	}
	fmt.Println(p)
}

// Demo4: Printf和反射
type Stringer interface {
	String() string
}

type Day int

var dayNames = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Day) String() string {
	return dayNames[d]
}

func Print(args ...any) {
	for index, arg := range args {
		if index > 0 {
			os.Stdout.WriteString(" ")
		}
		switch a := arg.(type) {
		case Stringer:
			os.Stdout.WriteString(a.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
		default:
			os.Stdout.WriteString("???")
		}
	}
	fmt.Println()
}

func TestPrint(t *testing.T) {
	Print(Day(1), "happy", 1)
}
