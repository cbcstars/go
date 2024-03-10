package demo8_map

import (
	"fmt"
	"testing"
)

// map概念(Demo1)
/*
1:map 是引用类型，可以使用如下声明：
	var map1 map[keytype]valuetype
	var map1 map[string]int
	（[keytype] 和 valuetype 之间允许有空格，但是 gofmt 移除了空格）
2:在声明的时候不需要知道 map 的长度，map 是可以动态增长的。
3:未初始化的 map 的值是 nil。
4:key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float32(64)。
  所以数组、切片和结构体不能作为 key (译者注：含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），
  但是指针和接口类型可以。如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。
5:value 可以是任意类型的；通过使用空接口类型（详见第 11.9 节），我们可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言（详见第 11.3 节）。
6:map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。
  通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果你很在乎性能的话还是建议用切片来解决问题。
7:map 也可以用函数作为自己的值，这样就可以用来做分支结构（详见第 5 章）：key 用来选择要执行的函数。
8:如果 key1 是 map1 的 key，那么 map1[key1] 就是对应 key1 的值，就如同数组索引符号一样（数组可以视为一种简单形式的 map，key 是从 0 开始的整数）。
	key1 对应的值可以通过赋值符号来设置为 val1：map1[key1] = val1。
	令 v := map1[key1] 可以将 key1 对应的值赋值给 v；如果 map 中没有 key1 存在，那么 v 将被赋值为 map1 的值类型的空值。
9:常用的 len(map1) 方法可以获得 map 中的 pair 数目，这个数目是可以伸缩的，因为 map-pairs 在运行时可以动态添加和删除。
10:map的初始化
	a:map 可以用 {key1: val1, key2: val2} 的描述方法来初始化，就像数组和结构体一样。
	b:map 是 引用类型 的： 内存用 make() 方法来分配。
		map 的初始化：var map1 = make(map[keytype]valuetype)。
		或者简写为：map1 := make(map[keytype]valuetype)。
11:不要使用 new()，永远用 make() 来构造 map
   注意 如果你错误地使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	mapCreated := new(map[string]float32)
	接下来当我们调用：mapCreated["key1"] = 4.5 的时候，编译器会报错：
	invalid operation: mapCreated["key1"] (index of type *map[string]float32).
*/

// map容量(Demo2)
/*
1:和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。
  但是你也可以选择标明 map 的初始容量 capacity，就像这样：make(map[keytype]valuetype, cap)。例如：
	map2 := make(map[string]float32, 100)
2:当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。
*/

// 测试键值对是否存在及删除元素
/*

 */

func TestMap(t *testing.T) {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{
		"one": 1,
		"two": 2,
	}
	mapCreated := make(map[string]int)
	mapAssigned = mapLit

	fmt.Println(mapLit["one"])
	fmt.Println(mapLit["two"])
	mapAssigned["two"] = 3
	fmt.Println(mapLit["two"])
	fmt.Println(mapCreated["one"])

	// value 可以是任意类型
	mf := map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
		3: func() int {
			return 30
		},
	}
	fmt.Println(mf)
}
