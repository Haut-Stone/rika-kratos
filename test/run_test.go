package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	fmt.Println("多测试保证代码质量")
}

// * 测试反射
func TestReflect(t *testing.T) {
	type cat struct {
	}
	ins := &cat{}
	typeA := reflect.TypeOf(ins) // - 这里类型是指针
	fmt.Println(typeA.Name(), typeA.Kind())
	typeA = typeA.Elem() // - 这里是取指针所对应的元素的类型
	fmt.Println(typeA.Name(), typeA.Kind())
}

// * 通过反射获取结构体成员
func TestReflectStructField(t *testing.T) {

	type cat struct { // 声明一个空结构体
		Name string
		Type int `json:"type" id:"100"` // 带有结构体tag的字段
	}

	ins := cat{Name: "mimi", Type: 1} // - 创建cat的实例, 注意这里不是指针
	typeOfCat := reflect.TypeOf(ins)
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)                                                                // 获取每个成员的结构体字段类型
		fmt.Printf("name: %v  tag: '%v'  index: %v\n", fieldType.Name, fieldType.Tag, fieldType.Index) // 输出成员名和tag
	}

	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}

// * 泛型的东西太多了, 实在消化不了
func TestT(t *testing.T) {
	type Slice[T int | float32 | float64] []T
	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T", a) //输出：Type Name: Slice[int]

	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T", b) //输出：Type Name: Slice[float32]

}
