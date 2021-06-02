package base

/**
每个goroutine的栈默认为2k，当goroutine执行结束且未进行栈扩容，则放置到所属P的空闲G队列中，进行重用，栈大于2k的G，进行销毁。
*/
