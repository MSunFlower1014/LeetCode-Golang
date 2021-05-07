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