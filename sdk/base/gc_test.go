package base

import (
	"runtime"
	"testing"
	"time"
)

/*
设置环境变量开启 gc日志
GOGCTRACE=1;GODEBUG=gctrace=1

linux中
GODEBUG=gctrace=1 执行文件即可

格式

gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
含义

	gc #        GC次数的编号，每次GC时递增
	@#s         距离程序开始执行时的时间
	#%          GC占用的执行时间百分比
	#+...+#     GC使用的时间
	#->#-># MB  GC开始，结束，以及当前活跃堆内存的大小，单位M
	# MB goal   全局堆内存大小
	# P         使用processor的数量
如果每条信息最后，以(forced)结尾，那么该信息是由runtime.GC()调用触发

垃圾回收的触发是由一个gcpercent的变量控制的，当新分配的内存占已在使用中的内存的比例超过gcprecent时就会触发。
比如，gcpercent=100，当前使用了4M的内存，那么当内存分配到达8M时就会再次gc。如果回收完毕后，内存的使用量为5M，那么下次回收的时机则是内存分配达到10M的时候。
也就是说，并不是内存分配越多，垃圾回收频率越高，这个算法使得垃圾回收的频率比较稳定，适合应用的场景。

gcpercent的值是通过环境变量GOGC获取的，如果不设置这个环境变量，默认值是100。如果将它设置成off，则是关闭垃圾回收。

基于标记清除算法，第一步stop world 将root区（包括当前运行时栈数据和全局数据区域）标记为灰色，通过灰色节点一次遍历，将其引用数据标记为灰色，自身改为黑色，
重复此过程，知道没有灰色节点存在，则白色节点为待清理数据，标记结束后执行第二步：清除所有未引用数据。

内存分配：每个M都持有局部缓存MCache，分配小对象直接从MCache中分配，无需加锁，提升效率，如果内存不足，像MCentral申请，如果MCentral也不足，则像MHeap申请空间，MHeap空间不足，则像系统申请空间
大于32k的对象直接分配到MHeap中
*/
func TestTc(t *testing.T) {
	var s []int
	for i := 0; i < 1000; i++ {
		s = append(s, i)
	}

	t.Logf("force gc")
	runtime.GC()
	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	t.Logf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
	time.Sleep(time.Second * 3600)
}
