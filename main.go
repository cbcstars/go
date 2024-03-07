package _go

import (
	"fmt"
	"github.com/cbcstars/go/demo10_struct"
)

func main() {
	// Demo5:强制使用工厂方法
	m := demo10_struct.NewMatrix()
	fmt.Println(m, m.Name)
}
