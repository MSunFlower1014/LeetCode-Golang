# 链表
三种最常见的链表结构，它们分别是：单链表、双向链表和循环链表。
  
数组简单易用，在实现上使用的是连续的内存空间，可以借助 CPU 的缓存机制，预读数组中的数据，所以访问效率更高。  
而链表在内存中并不是连续存储，所以对 CPU 缓存不友好，没办法有效预读。数组的缺点是大小固定，一经声明就要占用整块连续内存空间。  
如果声明的数组过大，系统可能没有足够的连续内存空间分配给它，导致“内存不足（out of memory）”。  
如果声明的数组过小，则可能出现不够用的情况。  
这时只能再申请一个更大的内存空间，把原数组拷贝进去，非常费时。  
链表本身没有大小的限制，天然地支持动态扩容，我觉得这也是它与数组最大的区别。

[Redis的缓存淘汰策略LRU与LFU](https://www.jianshu.com/p/c8aeb3eee6bc)  
LRU (Least recently used) 最近最少使用，如果数据最近被访问过，那么将来被访问的几率也更高。  
Redis维护了一个24位时钟，可以简单理解为当前系统的时间戳，每隔一定时间会更新这个时钟。每个key对象内部同样维护了一个24位的时钟，当新增key对象的时候会把系统的时钟赋值到这个内部对象时钟。比如我现在要进行LRU，那么首先拿到当前的全局时钟，然后再找到内部时钟与全局时钟距离时间最久的（差最大）进行淘汰，这里值得注意的是全局时钟只有24位，按秒为单位来表示才能存储194天，所以可能会出现key的时钟大于全局时钟的情况，如果这种情况出现那么就两个相加而不是相减来求最久的key。  
LFU (Least frequently used) 最不经常使用，如果一个数据在最近一段时间内使用次数很少，那么在将来一段时间内被使用的可能性也很小。  
LFU把原来的key对象的内部时钟的24位分成两部分，前16位还代表时钟，后8位代表一个计数器。16位的情况下如果还按照秒为单位就会导致不够用，所以一般这里以时钟为单位。而后8位表示当前key对象的访问频率，8位只能代表255，但是redis并没有采用线性上升的方式，而是通过一个复杂的公式，通过配置两个参数来调整数据的递增速度。  
FIFO (Fist in first out) 先进先出， 如果一个数据最先进入缓存中，则应该最早淘汰掉。  

[单链表反转](https://github.com/MSunFlower1014/LeetCode-Golang/tree/master/leetcode/201-299/206-reverseList.go)  
[链表中环的检测](https://github.com/MSunFlower1014/LeetCode-Golang/tree/master/leetcode/101-200/141-hasCycle.go)  
[两个有序的链表合并](https://github.com/MSunFlower1014/LeetCode-Golang/tree/master/leetcode/1-100/21-mergeTwoLists.go)  
[删除链表倒数第 n 个结点](https://github.com/MSunFlower1014/LeetCode-Golang/tree/master/leetcode/1-100/19-removeNthFromEnd.go)  
[求链表的中间结点](https://github.com/MSunFlower1014/LeetCode-Golang/tree/master/leetcode/800-899/876-middleNode.go)
