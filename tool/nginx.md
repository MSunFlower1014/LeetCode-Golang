# 进程模型
一个master进程  
多个worker进程  
* Master 进程：管理 Worker 进程。  
对外接口：接收外部的操作（信号）；  
对内转发：根据外部操作的不同，通过信号管理 Worker；  
监控：监控 Worker 进程的运行状态，Worker 进程异常终止后，自动重启 Worker 进程。  
* Worker 进程：所有 Worker 进程都是平等的。  
实际处理：网络请求，由 Worker 进程处理。  
Worker 进程数量：在 nginx.conf 中配置，一般设置为核心数，充分利用 CPU 资源，同时，避免进程数量过多，避免进程竞争 CPU 资源，增加上下文切换的损耗。  
多路复用 + epoll 事件驱动  

# 热加载原理
1.master重新加载配置文件  
2.master启动新的worker进程  
3.停止老worker进程  

# 备注
Nginx 反向代理时，会建立 Client 的连接和后端 Web Server 的连接，占用 2 个连接。   
* poll机制： 
解决了连接数限制：poll 中将 select 中的 fd_set 替换成了一个 pollfd 数组，解决 fd 数量过小的问题。  
数据复制：用户空间和内核空间，复制连接就绪状态信息。  
* epoll机制  
事件机制：避免线性扫描，为每个 fd，注册一个监听事件，fd 变更为就绪时，将 fd 添加到就绪链表。  
fd 数量：无限制（OS 级别的限制，单个进程能打开多少个 fd）。  

## 处理连接
首先，nginx在启动时，会解析配置文件，得到需要监听的端口与ip地址，然后在nginx的master进程里面，  
先初始化好这个监控的socket(创建socket，设置addrreuse等选项，绑定到指定的ip地址端口，再listen)，  
然后再fork(一个现有进程可以调用fork函数创建一个新进程。由fork创建的新进程被称为子进程 )出多个子进程出来，然后子进程会竞争accept新的连接。  

## 负载均衡策略
轮询，随机，最少链接，ip哈希，url哈希，权重  

## 配置
1、全局块：配置影响nginx全局的指令。一般有运行nginx服务器的用户组，nginx进程pid存放路径，日志存放路径，配置文件引入，允许生成worker process数等。  
2、events块：配置影响nginx服务器或与用户的网络连接。有每个进程的最大连接数，选取哪种事件驱动模型处理连接请求，是否允许同时接受多个网路连接，开启多个网络连接序列化等。  
3、http块：可以嵌套多个server，配置代理，缓存，日志定义等绝大多数功能和第三方模块的配置。如文件引入，mime-type定义，日志自定义，是否使用sendfile传输文件，连接超时时间，单连接请求数等。  
4、server块：配置虚拟主机的相关参数，一个http中可以有多个server。  
5、location块：配置请求的路由，以及各种页面的处理情况。  

## 参考连接
[Nginx 配置详解](https://www.runoob.com/w3cnote/nginx-setup-intro.html)  