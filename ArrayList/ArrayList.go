package ArrayList

import (
	"errors"
	"fmt"
)

// 定义接口
type List interface {
	Size() int                               // 获取大小
	Get(index int) (interface{}, error)                 // 抓取第几个元素
	Set(index int, newVal interface{})    // 修改指定位置的元素
	Insert(index int, newVal interface{}) // 在第几个位置插入元素
	Append(newVal interface{})            // 追加元素

	String() string // 变成字符串
}

// 数据结构
type ArrayList struct {
	// 保存数据的数组
	dataStorage []interface{}
	// 当前数组长度
	theSize int
}

// 构造函数
func NewArrayList() *ArrayList {
	list := new(ArrayList)
	// 设置值
	list.dataStorage = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

// 打印到屏幕
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStorage)
}

// 追加值
func (list *ArrayList) Append(newVal interface{}) {
	list.dataStorage = append(list.dataStorage, newVal)
	list.theSize++
}

// 抓取第几个元素
func (list *ArrayList) Get(index int) (interface{}, error){
	if index < 0 || index > list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStorage[index], nil
}
// 获取数组大小
func (list *ArrayList) Size() int {
	return list.theSize
}
