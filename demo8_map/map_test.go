package demo8_map

import (
	"fmt"
	"sort"
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

// 测试键值对是否存在及删除元素(Demo3)
/*
1:测试 map1 中是否存在 key1：
  在例子 8.1 中，我们已经见过可以使用 val1 = map1[key1] 的方法获取 key1 对应的值 val1。如果 map 中不存在 key1，val1 就是一个值类型的空值。
  这就会给我们带来困惑了：现在我们没法区分到底是 key1 不存在还是它对应的 value 就是空值。
  为了解决这个问题，我们可以这么用：val1, isPresent = map1[key1]
  isPresent 返回一个 bool 值：如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent 为 true；
  如果 key1 不存在，val1 就是一个空值，并且 isPresent 会返回 false。
2:如果你只是想判断某个 key 是否存在而不关心它对应的值到底是多少，你可以这么做：
  _, ok := map1[key1] // 如果key1存在则ok == true，否则ok为false
  或者和 if 混合使用：
  if _, ok := map1[key1]; ok {
  	  // ...
  }
3:从 map1 中删除 key1：
  直接 delete(map1, key1) 就可以。
  如果 key1 不存在，该操作不会产生错误。
*/

// for-range循环(Demo4)
/*
1:可以使用 for 循环读取 map：
	for key, value := range map1 {
		...
	}
第一个返回值 key 是 map 中的 key 值，第二个返回值则是该 key 对应的 value 值；
这两个都是仅 for 循环内部可见的局部变量。其中第一个返回值 key 值是一个可选元素。
2:如果你只关心值，可以这么使用：
	for _, value := range map1 {
		...
	}
3:如果只想获取 key，你可以这么使用：
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}
4:注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
  map 的本质是散列表，而 map 的增长扩容会导致重新进行散列，这就可能使 map 的遍历结果在扩容前后变得不可靠，
  Go 设计者为了让大家不依赖遍历的顺序，每次遍历的起点--即起始 bucket 的位置不一样，即不让遍历都从某个固定的 bucket0 开始，所以即使未扩容时我们遍历出来的 map 也总是无序的。
*/

// map类型的切片(Demo5)
/*
1:假设我们想获取一个 map 类型的切片，我们必须使用两次 make() 函数，第一次分配切片，第二次分配切片中每个 map 元素
2:需要注意的是，应当像 A 版本那样通过索引使用切片的 map 元素。在 B 版本中获得的项只是 map 值的一个拷贝而已，所以真正的 map 元素没有得到初始化
*/

// map的排序(Demo6)
/*
1:map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
2:如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort 包），然后可以使用切片的 for-range 方法打印出所有的 key 和 value。
*/

// Demo1:概念
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

// Demo2:容量
func TestCapacity(t *testing.T) {
	// 定义map的初始容量
	m := make(map[string]int, 100)
	// 键值对数量
	fmt.Println(len(m))
	// cap() 函数仅适用于切片（slice）、数组（array）和通道（channel），并不适用于 map
	//fmt.Println(cap(m))
}

// Demo3:判断键是否存在&删除键值
func TestIsPresentAndDelete(t *testing.T) {
	m := make(map[string]int, 10)
	m["chen"] = 28
	m["yun"] = 30

	if v, ok := m["chen"]; ok {
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println("map key don't present")
		}
	}

	delete(m, "chen")

	fmt.Println(m)
}

// Demo4: for range 循环
func TestForRange(t *testing.T) {
	m := map[string]int{
		"chen": 28,
		"yun":  30,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	for k := range m {
		fmt.Println(k)
	}
}

// Demo5:map类型的切片
func TestMapSlice(t *testing.T) {
	// Version A:
	items := make([]map[int]int, 5)
	for k := range items {
		items[k] = make(map[int]int, 1)
		items[k][1] = 2
	}
	fmt.Println(items)

	// Version B: NOT GOOD!
	item2 := make([]map[int]int, 5)
	for _, item := range item2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element
		item[1] = 2                 // This item will be lost on th next iteration
	}
	fmt.Println(item2)
}

// Demo6:map的排序
func TestMapSort(t *testing.T) {
	var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
