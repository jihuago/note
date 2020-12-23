## MySQL主备的基本原理
备库B跟主库A之间维持了一个长连接。主库A内部有一个线程，专门用于服务备库B的这个长连接。
一个事务日志同步的完整过程是这样的：

1. 在备库B上通过change master命令，设置主库A的Ip、端口、用户名、密码，以及要从哪个位置开始请求binlog，
这个文件包含文件名和日志偏移量
2. 在备库B上执行start slave命令，这时候备库会启动两个线程，就是io_thread和sql_thread，其中io_thread负责与主库建立连接
3. 主库A校验完用户名、密码后，开始按照备库B传过来的位置，从本地读取binlog，发给B
4. 备库B拿到binlog后，写到本地文件，称为中转日志relay log
5. sql_thread读取中转日志，解析出日志里的命令，并执行

## binlog的三种格式对比
1. statement格式（一般不适用statement）
statement格式，记录到binlog里的是sql语句原文。
优点：占用空间小；缺点：有些statement格式的binlog可能会导致主备不一致。

2. row格式（一般采用这种方式）
当binlog_format使用row格式时，binlog里面记录了真实删除行主键id，这样binlog传到备库的时候，肯定是不存在主备删除不同行的问题

缺点：row很占空间，比如用一个delete语句删除10万行数据，用statement就是一个SQL语句被记录到binlog，占用几十个字节的空间。
但如果用row格式的binlog，就要把这10万条记录都写到binlog中。
这样做，不仅会占用更大的空间，同时写binlog也要耗费IO资源，影响执行速度。

3. mixed格式
mixed格式：mysql会判断SQL语句是否可能引起主备不一致，如果可能，就用row格式，否则就用statement格式。

### binlog配置设置
* 修改my.cnf

```markdown

binlog_format = mixed
log_bin = mysql-bin
server-id = 12


```




