# MYSQL

## 建议
1.在联合索引中将选择性最高的列放在索引最前面。  
以age 和gender为索引，显然age要放在前面，因为性别就两种选择男或女，选择性不如age  
2.永远设置主键且最好设置为int且自增，字符串的排序比int复杂很多  
3.多使用explain查看SQL复杂度  
4.尽量设置not null，存在null列会使索引失效
5.拆分大的insert和delete
6.明确只有一条的时候加 LIMIT 1  
7.like条件 % 不能放在前面  
8.尽量不要使用内置函数（根据耗时统计，通过函数计算不会使用索引，直接入库时记录才好）  
9.Inner join 、left join、right join，优先使用Inner join，如果是left join，左边表结果尽量小。  
Inner join 内连接，在两张表进行连接查询时，只保留两张表中完全匹配的结果集；  
left join 在两张表进行连接查询时，会返回左表所有的行，即使在右表中没有匹配的记录；  
right join 在两张表进行连接查询时，会返回右表所有的行，即使在左表中没有匹配的记录。  
10.不要使用!=  
11.count(字段)<count(主键 id)<count(1)≈count(*)，所以我建议你，尽量使用 count(*)


## 慢查询
表设计问题，是否无索引
SQL问题，是否未用到索引
索引失效
join的表太多了
数据量太大，是否需要分库分表
数据库服务器IO磁盘或磁盘性能不行
网络是否波动
请求量是否过大，是否可以增加缓存

## SQL执行流程
客户端 -> 连接器（管理连接，权限验证） -> 分析器（语法分析） -> 优化器（执行计划生成，索引选择） - > 执行器（操作引起，返回结果） -> 存储引擎（存储数据，提供读写接口）  
[参考连接](https://time.geekbang.org/column/article/68319)

## 长连接
为减少建立连接的损耗，建议尽量使用长连接  
长连接内存消耗可能过快，是由于执行过程中使用的内存在断开连接时才释放，长期累积，可能导致OOM  
解决：1.定时或执行大查询后断开连接2.执行mysql_reset_connection 来重新初始化连接资源  

## ACID
* 原子性（atomicity)  
一个事务要么全部提交成功，要么全部失败回滚，不能只执行其中的一部分操作，这就是事务的原子性
* 一致性（consistency)  
事务的执行不能破坏数据库数据的完整性和一致性，一个事务在执行之前和执行之后，数据库都必须处于一致性状态。  
如果数据库系统在运行过程中发生故障，有些事务尚未完成就被迫中断，这些未完成的事务对数据库所作的修改有一部分已写入物理数据库，这是数据库就处于一种不正确的状态，也就是不一致的状态  
* 隔离性（isolation）  
事务的隔离性是指在并发环境中，并发的事务时相互隔离的，一个事务的执行不能不被其他事务干扰。  
不同的事务并发操作相同的数据时，每个事务都有各自完成的数据空间，即一个事务内部的操作及使用的数据对其他并发事务时隔离的，并发执行的各个事务之间不能相互干扰。  
* 持久性（durability）  
一旦事务提交，那么它对数据库中的对应数据的状态的变更就会永久保存到数据库中。--即使发生系统崩溃或机器宕机等故障，只要数据库能够重新启动，那么一定能够将其恢复到事务成功结束的状态
## 隔离级别
并发可能导致：脏读:读到其他事务未提交的数据；不可重复读：前后读取的记录内容不一致；幻读：前后读取的记录数量不一致。  
读未提交是指，一个事务还没提交时，它做的变更就能被别的事务看到。  
读提交是指，一个事务提交之后，它做的变更才会被其他事务看到。  
可重复读是指，一个事务执行过程中看到的数据，总是跟这个事务在启动时看到的数据是一致的。当然在可重复读隔离级别下，未提交变更对其他事务也是不可见的。  
串行化，顾名思义是对于同一行记录，“写”会加“写锁”，“读”会加“读锁”。当出现读写锁冲突的时候，后访问的事务必须等前一个事务执行完成，才能继续执行。 

## 锁
* 全局锁  
 MySQL 提供了一个加全局读锁的方法，命令是 Flush tables with read lock (FTWRL)。使整个库处于只读状态，用于全库逻辑备份。  
* 表级锁  
表锁的语法是 lock tables … read/write  
另一类表级的锁是 MDL（metadata lock)  
* 行锁
MySQL 的行锁是在引擎层由各个引擎自己实现的。但并不是所有的引擎都支持行锁，比如 MyISAM 引擎就不支持行锁。
* 间隙锁
for update 不止会锁住所在行，也会锁住与上下的间隙  