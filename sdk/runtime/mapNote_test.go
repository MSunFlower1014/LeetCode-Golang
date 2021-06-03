package runtime

import "testing"

/*
Go 语言字典的键类型不可以是函数类型、字典类型和切片类型。
Go语言规范规定，在键类型的值之间必须可以施加操作符==和！=，即键类型的值必须要支持判等操作。
由于函数类型、字典类型、切片类型不支持判等操作，所以字典的键类型不能是这些类型。
不建议使用高级数据类型（数组、接口、结构体）作为字典的键类型，它们的值求哈希较慢，并且它们的值中存在变数（可以修改其中的元素值），还需要对其内部元素类型进行判断
如键的类型是数组类型，那么还要确保该类型的元素类型不是函数类型、字典类型或切片类型

在字典中添加nil作为key-value时，会panic

1.应该优先考虑哪些类型作为字典的键类型？
求哈希和判等操作的速度越快，对应的类型就越适合作为键类型。

map 类型源码地址：src/runtime/map.go

type hmap struct {
	count     int // 数据总量
	B         uint8  // 容量为2的B次方
	hash0     uint32 // hash种子

	buckets    unsafe.Pointer // 程度为2的B次方的桶数组，当count为0时可能为nil
	oldbuckets unsafe.Pointer // 扩容时旧数据保存桶数组，仅仅当扩容时不为nil
}

make map 流程：
new(hmap) - 计算并设置hash0和B - 创建桶数组

扩容:新建2倍容量的桶数组，并包留旧桶数组
哈希表数据的实际复制是以增量方式完成的
通过growWork（）和excavate（）

get:根据低位hash获取桶数组下标，高八位hash通过比较后获取key-value（每个桶内8个元素，一次比较速度也很快）
hash的高8位存储在了bucket的tophash中，方便比较

先查询旧桶数组，如果找到顺便迁移到新桶（额外迁移一个pair），未找到时查询新桶数组

遍历顺序问题？
每次遍历会取随机数从某一bucket开始遍历
*/

func TestMap(t *testing.T) {
	m := make(map[string]string, 16)
	t.Logf("map length is %v ", len(m))
	if m == nil {
		t.Error("")
	}
	m["1"] = ""
	t.Logf("map length is %v ", len(m))
	//t.Logf("map cap is %v ", cap(m))
}
