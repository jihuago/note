## 到底用不用join

```mysql

CREATE TABLE `t1` (
 `id` int(11) NOT NULL,
  `a` int(11) DEFAULT NULL, 
  `b` int(11) DEFAULT NULL, PRIMARY KEY (`id`), 
  KEY `a` (`a`)
) ENGINE=InnoDB;

CREATE TABLE `t2` (
 `id` int(11) NOT NULL,
  `a` int(11) DEFAULT NULL, 
  `b` int(11) DEFAULT NULL, PRIMARY KEY (`id`), 
  KEY `a` (`a`)
) ENGINE=InnoDB;

```

> select * from t1 straight_join t2 on (t1.a=t2.a);


t1是驱动表，t2是被驱动表


如果可以使用被驱动表的索引：

* 使用join语句，性能比拆成多个单表执行SQL语句的性能要好
* 如果要使用join语句，需要让小表做驱动表

### straight_join
straight_join功能类似inner join，但能让左边的表来驱动右边的表，能改表优化器对于联表的执行顺序。

* straight join只适用于内连接，因为left join、right join已经知道了哪个表作为驱动表

## Block Nested-Loop Join
> select * from t1 straight_join t2 on (t1.a=t2.b);

当被驱动表上没有可用的索引，算法流程如下：
1. 把表t1的数据读入线程内存join_buffer中，由于我们这个语句写的是select *，因此是把整个表t1放入了内存
2. 扫描表t2，把表t2中的每一行取出来，跟join_buffer中的数据做对比，满足join条件的，作为结果集的一部分返回

以上情况，选择大表还是小表做驱动表，执行耗时是一样的。

* 要是驱动表是一个大表，join_buffer放不下如何处理？
* join_buffer大小由参数join_buffer_size设定，默认256k。如果放不下驱动表的所有数据，策略就是分段放

* join_buffer_size越大，一次可放入的行越多，分成的段数就越少，对被驱动表的全表扫描次数就越少
* join语句很慢，可以尝试将join_buffer_size改大

## 能不能使用Join语句
1. 可以用上被驱动表上的索引的，可以使用join
2. 如果使用Block Nested-Loop Join算法，扫描次数就会过多。尤其是在大表的join操作，这样可能要扫描被驱动表很多次，
会占用大量的系统资源。这种join尽量不要用

## 什么是小表
在决定哪个表做驱动表的时候，应该是两个表按照各自的条件过滤，过滤完成之后，计算参与join的各个字段的总数据量，数据量小的那个表
就是“小表”，应该作为驱动表。






