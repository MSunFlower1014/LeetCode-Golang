---
title: TCP连接那些事
date: 2021-04-19 11:18:23
tags:
---
# TCP
传输控制协议
## 可靠性保证
1.应用数据被分割成TCP认为最合适发送的数据块（所以需要syn确认报文时序）  
2.当TCP发出一个段后，启动一个定时器，等待收到确认ACK，如果没收到将重发这个报文  
3.当TCP收到另一端的请求，将推迟几分之一秒后发送确认（批量确认）  
4.TCP将保持首部和数据的校验和，检测数据在传输中的任何变化  
5.如果报文在IP层失序，TCP将对收到的数据进行重新排序  
6.TCP丢弃重复的数据  
7.TCP提供流量控制（滑动窗口针对接收端和慢启动针对发送端）  

## TCP数据格式
|16位源端口号|16位目的端口号|||
|---|---|---|---|
|32位序号|
|32位ack|
|4位首部长度|6位保留位|6位标志位|16位窗口大小|
|16位检验和|16位紧急指针|
|选项|
|数据|

每个TCP段都包含源端和目的端的端口号，用于寻找发端和收端应用进程。这两个值加上IP首部中的源端IP地址和目的端IP地址唯一确定一个TCP连接。
窗口大小是一个16bit字段，所以最大位65635字节  
标注位：  
 
|标志|3字符缩写|描述|
|---|---|---|
|S|SYN|同步序号|
|F|FIN|发送方完成数据发送|
|R|RST|复位连接|
|P|PSH|尽可能将数据送往接受进程|
|.| |以上四个标注为均置0|

## 半连接和全连接
当服务器接收到客户端的SYN后，会将连接放入半连接队列，回复ACK及自身SEQ，当收到客户端回复的ACK时，将连接从半连接队列移除放入全连接队列。  
如果全连接队列满了，会丢弃客户端回复，并定时重发 回复ACK及自身SEQ 。  
[参考连接](https://www.jianshu.com/p/6a0fcb1008d6)

## 三次握手
1.客户端发送SYN，携带初始序号ISN到目的端服务器  
2.目的端收到后，回复 SEQ+1 的ACK以及自身的初始序号ISN2  
3.客户端回复 ISN2 +1 的ACK  

![TCP_OPEN](https://github.com/MSunFlower1014/LeetCode-Golang/blob/master/protocol/img/TCP.png?raw=true)

ISN随时间变化，因此每个连接都具有不同的ISN，RFC 793 [Postel 1981c]指出ISN可看作是一个32比特的计数器，每4ms加1。  
这样选择序号的目的在于防止在网络中被延迟的分组在以后又被传送，而导致某个连接的一方对它作错误的解释，保证应用层接收到的数据不会因为网络上的传输的问题而乱序

## 四次挥手
1.客户端发送FIN，告诉服务器已完成数据发送  
2.服务器回复ACK，告诉客户端已收到  
3.服务器发送FIN，告诉客户端已完成数据发送  
4.客户端回复ACK，告诉服务端已收到  

![TCP_CLOSE](https://github.com/MSunFlower1014/LeetCode-Golang/blob/master/protocol/img/TCP_CLOSE.png?raw=true)

一个TCP连接是全双工（即数据在两个方向上能同时传递），因此每个方向必须单独地进行关闭。  
这原则就是当一方完成它的数据发送任务后就能发送一个FIN来终止这个方向连接。   
当一端收到一个FIN，它必须通知应用层另一端几经终止了那个方向的数据传送。发送FIN通常是应用层进行关闭的结果。  

### 2MSL等待时间
MSL表示报文段最大生存时间，是任何报文段被丢弃前在网络内的最长时间，现实中通常为30s或1分钟  

当TCP执行一个主动关闭，并发回最后一个ACK，该连接必须在TIME_WAIT状态停留的时间为2倍的MSL。  
1.当ACK丢失时，对方会重发FIN，等待2SML可以保证对方收到了ACK（如果2MSL未重发表示收到ACK）。  
2.使本连接持续的时间所产生的所有报文段都从网络中消失。这样就可以使下一个新的连接中不会出现这种旧的连接请求的报文段。   
在连接处于2MSL等待时，任何迟到的报文段将被丢弃。 

## 粘包
TCP为提高传输效率，发送方往往要收集到足够多的数据后才发送一个TCP段。  
若连续几次需要send的数据都很少，通常TCP会根据优化算法把这些数据合成一个TCP段后一次发送出去，这样接收方就收到了粘包数据。  
解决：为字节流加上自定义固定长度报头，报头中包含字节流长度，然后一次send到对端，对端在接收时，先从缓存中取出定长的报头，然后再取真实数据。  
比如ESB自定义协议：8位报文长度 + xml格式报文，长度参数可以解决粘包和拆包报文  

## 快速重传
在收到一个失序的报文段时，TCP会发送重复的ACK让对方知道收到一个失序报文段，并告诉对方自己希望收到的序号  
此时会直接重新重传丢失的数据报文段，而不是等超时定时器溢出

## TCP/IP的分层
|协议|层级|
|---|---|
|HTTP|应用层|
|TCP/UDP|传输层|
|IP|网络层|
|硬件|链路层|