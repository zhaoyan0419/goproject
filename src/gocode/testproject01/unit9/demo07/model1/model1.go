package model1

type Student struct {
	Name string
	Age  int
}

type persion struct {
	Name string
	Age  int
}
type myFirstStruct struct {
	name string
}

// 工厂函数，由于myFirstStruct不可导出，所以当其他包想使用该结构体时，需要调用工厂函数来创建实例
func NewmyFirstStruct(n string) *myFirstStruct {
	return &myFirstStruct{n}
}

// 封装，只能获取结构体字段的值，而不能修改
func (m myFirstStruct) GetName() string {
	return m.name
}
func NewPersion(n string, a int) *persion {
	return &persion{n, a}
}
