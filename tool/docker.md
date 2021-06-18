# Docker

## 环境隔离-NameSpace
包含 PID-NameSpace 隔离进程id信息
Mount-Namespace 隔离挂载点信息
NetWork-namespace 隔离网络设备和配置


## 资源限制-Linux Control Group
Linux Cgroups 的全称是 Linux Control Group。它最主要的作用，就是限制一个进程组能够使用的资源上限，包括 CPU、内存、磁盘、网络带宽等等。


## docker exec
一个进程，可以选择加入到某个进程已有的 Namespace 当中，从而达到“进入”这个进程所在容器的目的，这正是 docker exec 的实现原理。

## k8s的pod
抽象管理节点概念，支持单容器或多容器编排，并为多个容器提供组概念简化维护    
利于容器之间的组合；管理编排和资源共享