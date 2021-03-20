### 如何判断Redis是否变慢
* 查看Redis的响应延迟
    当发现Redis命令的执行时间突然增长到几秒，基本就认定Redis变慢了
  
* 基于环境下的Redis基线性能
    通过redis-cli命令提供的-intrinsic-latency选项来检测和同级测试期间内的最大延迟，这个延迟可以作为Redis的基线性能
  
```html
./redis-cli --intrinsic-latency 120
```
上面命令会打印120秒内检测到的最大延迟。

### 如何应对Redis变慢
Redis的性能诊断，一般从Redis自身的操作特性、文件系统和操作系统入手。
#### Redis自身操作特性的影响
1. 慢查询命令
慢查询命令，就是指在Redis中执行速度慢的命令，这会导致Redis延迟增加。
   
* 每个命令复杂度不同
* 官方文档对每个命令的复杂度都有介绍
  https://redis.io/commands/
  
* 如何应对慢查询命令
  当发现Redis性能变慢，通过Redis日志，或者latency monitor工具，查询变慢的请求，根据请求对应的具体命令以及官方文档，
  确认下是否采用了复杂度高的慢查询命令。
  
如果的确有大量的慢查询命令，有两种处理方式：
    
* 用其他高效命令代替
* 当你需要执行排序、交集、并集操作时，可以在客户端完成，而不要用SORT、SUNION、SINTER命令，以免拖慢Redis实例

2. 过期key操作
过期key自动删除操作，它是Redis用来回收内存空间的常用机制，应用广泛，本身就会引起Redis操作阻塞，导致性能变慢。
   
Redis键值对的key可以设置过期时间。默认情况下，Redis每100毫秒会删除一些过期key。具体算法如下：

* 采样ACTIVE_EXPIRE_CYCLE_LOOKUPS_PER_LOOP个数的key，并将其中过期的key全部删除
   
    ACTIVE_EXPIRE_CYCLE_LOOKUPS_PER_LOOP是Redis的一个参数，默认是20，那么，一秒内基本有200个过期key会被删除。
    这一策略对清除过期key、释放内存空间很有帮助。

* 如果超过25%的key过期了，则重复删除的过程，直到过期key的比例降至25%以下。

    如果触发了这条算法，Redis就会一直删除以释放内存空间。注意：删除操作是阻塞的（Redis4.0后可以用异步线程机制来减少阻塞影响）。
所以，一旦该条件触发，Redis的线程就会一直执行删除，这样就没办法正常服务其他的键值操作，就会进一步引起其他键值操作的延迟增加。
  
**频繁使用带有相同时间参数的EXPIREAT命令设置过期key**就可能会导致在同一秒内有大量的key同时过期。

### 关于如何分析、排查、解决Redis变慢问题
1. 使用了复杂度过高的命令（例如SORT/SUNION/KEYS）或一次查询全量数据（例如LRANGE key 0 N ，N很大）
分析： a. 查看showlog是否存在这些命令   b. Redis进程CPU使用率是否飙升(聚合运算命令导致)
   
解决：a. 不上复杂度过高的命令或用其他方式代替实现（放在客户端做） b. 数据尽量分批查询（LARNGE key 0 N 建议N<=100，查询全量数据建议使用HSCAN/SSCAN/ZSCAN）

2. 操作bigkey
分析： a. slowlog出现很多SET/DELETE变慢命令 b. 使用redis-cli -h $host -p $port --bigkeys 扫描出很多bigkey
   
解决： a. 优化业务，避免存储bigkey  b. Redis4.0+ 可开启lazy-free机制

3. 大量key集中过期
分析：a. 业务集中EXPIREAT/PEXPIREAT命令  b. Redis info中的expired_keys指标短期突增
   
解决：a. 优化业务，过期增加随机时间，把时间打散，减轻删除过期Key的压力  b. 运维层面，监控expired_keys指标，有短期突增及时报警排查

4. Redis内存达到maxmemory
分析：a. 实例内存达到maxmemory，且写入量大，淘汰key压力变大  b. Redis info中evcited_keys指标短期徒增
   
解决：a. 业务层面，根据情况调整淘汰策略（随机比LRU快）  b) 运维层面，监控evicted_keys指标，有短期突增及时报警 c) 集群扩容，多个实例减轻淘汰key的压力

5. 大量短连接请求
分析：redis处理大量短连接请求，TCP三次握手和四次挥手也会增加耗时
   
解决：使用长连接操作Redis

6. 生成RDB和AOF重写fork耗时严重
分析：a. Redis变慢只发生在生成RDB和AOF重写期间  b. 实例占用内存越大，fork拷贝内存页表越久  c. Redis info中latest_fork_usec耗时变长
   
解决：a) 实例尽量小 b) Redis尽量部署在物理机上 c) 优化备份策略（例如低峰期备份） d) 合理配置repl-backlog和slave client-output-buffer-limit，避免主从全量同步 e) 视情况考虑关闭AOF f) 监控latest_fork_usec耗时是否变长

7. 网卡负载过高

分析： a. TCP/IP层延迟变大，丢包重传变多  b. 是否存在流量过大的实例占满带宽

解决： a. 机器网络资源监控，负载过高及时报警  b. 提前规划部署策略，访问量大的实例隔离部署

总之，Redis的性能与CPU、内存、网络、磁盘读息息相关，任何一处发生问题，都会影响到Redis的性能


