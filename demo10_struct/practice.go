package demo10_struct

import "fmt"

// Demo5:强制使用工厂方法
type matrix struct {
	Name string
	side float32
}

func NewMatrix() *matrix {
	return new(matrix)
}

func main() {
	m := NewMatrix()
	fmt.Println(m, m.Name, m.side)
}
