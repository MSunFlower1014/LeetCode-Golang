package runtime

import (
	"testing"
	"time"
)

/*

一个通道相当于一个先进先出（FIFO）的队列

对通道的发送和接收操作都有哪些基本的特性？
1.对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
2.发送操作和接收操作中对元素值的处理都是不可分割的。
3.发送操作在完全完成之前会被阻塞。接收操作也是如此。

元素值从外界进入通道时会被复制。更具体地说，进入通道的并不是在接收操作符右边的那个元素值，而是它的副本。

接收表达式的结果同时赋给两个变量时，第二个变量的类型就是一定bool类型。
它的值如果为false就说明通道已经关闭，并且再没有元素值可取了。

如果通道关闭时，里面还有元素值未被取出，那么接收表达式的第一个结果，仍会是通道中的某一个元素值，而第二个结果值一定会是true。
由于通道的收发操作有上述特性，所以除非有特殊的保障措施，我们千万不要让接收方关闭通道，而应当让发送方做这件事。

chan
1.为nil的通道，发送和接受都会阻塞
2.被关闭的通道，发送会panic，接受会接收到类型的零值
3.关闭nil的通道，会panic
4.关闭已关闭的通道，会panic
5.无缓冲的通道，发送和接受会阻塞到双方完成连接
6.有缓冲的通道，在发送或接受缓冲队列满时进行发送或者接受会阻塞挂起

通道的长度代表通道当前包含的元素个数，容量就是初始化时你设置的那个数。

底层是双向链表

select中屏蔽已关闭的通道，可将通道置为nil，屏蔽该case
*/
func TestChan(t *testing.T) {
	m := make(chan int)

	chanSend(m)

	i := chanGet(m)

	if i != 1 {
		t.Fatalf("chan get value is not 1")
	}
}

/*
此参数管道只能用来发送消息，单向发送管道
*/
func chanSend(c chan<- int) {
	c <- 1
}

/*
此参数管道只能用来接受消息，单向接受管道
*/
func chanGet(c <-chan int) int {
	i := <-c
	return i
}

/*
无缓冲管道，发送和接受都会阻塞，直到双方成功建起连接
*/
func TestNoBufferChan(t *testing.T) {
	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		c <- 1
	}()
	before := time.Now()
	i := <-c
	end := time.Now()

	t.Logf("c get i = %v", i)
	t.Logf("c get time is  %v", end.UnixNano()-before.UnixNano())
}

/*
nil管道，发送和接受都会阻塞
*/
func DeadTestNilRecvChan(t *testing.T) {
	c := make(chan int)
	c = nil
	i := <-c
	t.Logf("c get i = %v", i)
}

/*
nil管道，发送和接受都会阻塞
*/
func DeadTestNilSendChan(t *testing.T) {
	c := make(chan int)
	c = nil
	c <- 1
}

/*
有缓冲管道，当发送缓冲满了，后续的发送回阻塞挂起
*/
func TestBufferSendChan(t *testing.T) {
	c := make(chan int, 4)
	for i := 0; i < 4; i++ {
		c <- 1
	}

	go func() {
		time.Sleep(time.Second * 5)
		<-c
	}()
	before := time.Now()
	c <- 1
	end := time.Now()
	t.Logf("c get time is  %v", end.Unix()-before.Unix())
}

/*
有缓冲管道，当接受缓冲为空，会阻塞挂起
*/
func TestBufferRecvChan(t *testing.T) {
	c := make(chan int, 4)
	c <- 1
	before := time.Now()
	i, flag := <-c
	end := time.Now()
	t.Logf("no recv get i = %v , flag = %v ", i, flag)
	t.Logf("c get time is  %v", end.Unix()-before.Unix())
}

/*
已关闭的管道，接受会获得管道类型的零值，flag 为 false
发送会发生panic
*/
func TestCloseChan(t *testing.T) {
	c := make(chan int, 4)
	close(c)
	i, flag := <-c
	t.Logf("no recv get i = %v , flag = %v ", i, flag)

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("cant recover paic ")
		}
		t.Logf("send to closed chan get panic : %v", err)
	}()
	c <- 1
}
