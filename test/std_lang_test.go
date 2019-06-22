package test

import (
	"fmt"
	"hello/model"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func Test1equal1(t *testing.T) {
	assert.Equal(t, 1, 1)
}

/*
引用cale模块,调用Add方法
*/
func Test_Mod_Import_Cale_Add(t *testing.T) {
	a := 1
	b := 2
	fmt.Printf("type of a is %T , size of a is %d", a, unsafe.Sizeof(a))
	fmt.Println()
	fmt.Printf("type of a is %T , size of a is %d", b, unsafe.Sizeof(b))
	fmt.Println()
	var c, d = model.Add(a, b)
	assert.Equal(t, c, 3)
	assert.Equal(t, d, 1)

}

/*
创建一个简单的切片
*/
func Test_SimpleSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	length := len(a)
	assert.Equal(t, length, 7)

	b := a[1:4]
	assert.Equal(t, len(b), 3)
}

func Test_MakeSlice(t *testing.T) {
	s := make([]int, 5, 5)
	s = append(s, 6)
	assert.Equal(t, len(s), 6)
}

func change(s ...string) []string {
	s[0] = "Go"
	s = append(s, "Playground")
	return s
}

/*
测试字符串数组切片，并修改其中的值。
*/
func Test_ChangeSlice(t *testing.T) {

	welcome := []string{"Hello", "World"}
	welcome = change(welcome...)
	assert.Equal(t, welcome, []string{"Go", "World", "Playground"})
}

/*
测试Map类型创建与取值
*/
func Test_MakeMap(t *testing.T) {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	employee := "jamie"

	assert.Equal(t, personSalary[employee], 15000)
	assert.Equal(t, personSalary["joe"], 0)

	_, ok := personSalary["joe"]

	assert.Equal(t, ok, false)
}

/*
测试指针操作。
*/
func Test_Op_SimplePoint1(t *testing.T) {
	b := 255
	var a *int = &b
	assert.Equal(t, *a, 255)

	*a++
	assert.Equal(t, b, 256)
}

/*
操作数组尽可能使用切片，不推荐使用指针
*/
func modiflyArrayValueFor0Index(slice []int) {
	slice[0] = 10
}

func modiflyPointValueFor0Index(arr *[]int) {
	(*arr)[0] = 10
	// ==  arr[0] = 10
}

/*
测试修改数组时，使用切片与指针是等价的。
*/
func Test_ArrayElementSliceChanged(t *testing.T) {
	array := []int{11, 20, 30}
	assert.Equal(t, array, []int{11, 20, 30})
	//modiflyArrayValueFor0Index(array)
	modiflyPointValueFor0Index(&array)
	assert.Equal(t, array, []int{10, 20, 30})

}
