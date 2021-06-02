# Container包
container  
## heap
可用来实现[优先队列](https://github.com/golang/go/blob/master/src/container/heap/example_pq_test.go)  
push或pop元素后会根据元素大小调整到对应位置（堆排序）  

## list
双向链表实现
```go
type List struct {
	root Element // 根元素
	len  int     // 链表长度
}

type Element struct {
	//前驱节点及后续节点
	next, prev *Element

	// 所属List
	list *List

	// 当前节点value
	Value interface{}
}
```
PushFront或PushBack方法可触发延迟加载，自动初始化链表  

## ring
```go
type Ring struct {
    //前驱和后继元素
	next, prev *Ring
	Value      interface{} 
}
```
循环链表  
需要指定长度  
[ring示例](https://github.com/golang/go/blob/master/src/container/ring/example_test.go)  
Do(f func(interface{})) 方法执行遍历操作  
Link(s *Ring) *Ring  方法可将两个循环链表连接  
Unlink(n int) *Ring  方法从下一元素开始删除n个元素  