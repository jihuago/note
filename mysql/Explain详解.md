### Explain详解
一条查询语句在经过MySQL查询优化器的各种基于成本和规则的优化后生成一个所谓的执行计划，这个执行计划展示了接下来具体执行查询的方式，比如多表链接的顺序是什么，
多余每个表采用什么访问方法来具体执行查询等。

> MySQL提供了Explain语句来帮助我们查看某个查询语句具体执行计划。搞懂Explain语句的各个输出项有利于提升我们查询语句的性能

id: 在一个大的查询语句中每个select关键字都对应一个唯一的id
select_type:select关键字对应的那个查询的类型
table：表名
partitions：匹配的分区信息
type：针对单表的访问方法
possible_leys：可能用到的索引
key：实际上使用的索引
key_len:实际使用到的索引长度
rows:预估的需要读取的记录条数
filtered:某个表经过搜索条件过滤后记录条数的百分比
Extra：一些额外的信息

#### 执行计划输出中各列详解
* table

不论查询语句有多复杂，包含了多少个表，到最后也是需要对每个表进行单表访问的，所以设计MySQL的规定Explain语句输出的每条记录都对应着某个单表的访问方法，
该条记录的table列代表着该表的表名。

* id

查询语句中每出现一个select关键字，MySQL就会为它分配一个唯一的id值。
> 在链接查询的执行计划中，每个表都会对应一条记录，这些记录的id列的值是相同的，出现在前边的表表示驱动表，出现在后边的表是被驱动表。

mysql> EXPLAIN SELECT * FROM s1 INNER JOIN s2;
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+---------------------------------------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra                                 |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+---------------------------------------+
|  1 | SIMPLE      | s1    | NULL       | ALL  | NULL          | NULL | NULL    | NULL | 9688 |   100.00 | NULL                                  |
|  1 | SIMPLE      | s2    | NULL       | ALL  | NULL          | NULL | NULL    | NULL | 9954 |   100.00 | Using join buffer (Block Nested Loop) |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+---------------------------------------+
2 rows in set, 1 warning (0.01 sec)

上面，查询优化器让s1作为驱动表，s2为被驱动表。

> 包含子查询的查询语句，有多个select，所以在包含子查询的查询语句的执行计划中，每个select关键字都会对应一个唯一的id值
> 查询优化器可能对子查询涉及子查询的查询语句进行重写，转换为链接查询。

* possible_keys 和 key
possible_keys列表示在某个查询语句中，对某个表执行单表查询时可能用到的索引有哪些，key列表示实际用到的索引有哪些。
  
possible_keys列中的值并不是越多越好，可能使用的索引越多，查询优化器计算查询成本时就得花费更长时间，所以如果可以的话，尽量删除那些用不到的索引。

* filtered

MySQL在计算驱动表扇出时采用的一个策略：
* 如果使用的是全表扫描的方式执行的单表查询，那么计算驱动表扇出时需要估计出满足搜索条件的记录到底有多少条
* 如果使用的是索引执行的单表扫描，那么计算驱动表扇出的时候需要估计出满足除使用到对应索引的搜索条件外的其他搜索条件的记录有多少条

* Extra

Extra列是用来说明一些额外信息的，可以通过这些额外信息来更准确的理解MySQL到底将如何执行给定的查询语句

    * No tables user
    * Impossible WHERE
        查询语句的WHERE子句永远为FALSE时将会提示该额外信息
    * Using index
        当查询列表以及搜索条件中只包含属于某个索引的列，也就是在可以使用索引覆盖的情况下，在Extra列将会提示该额外信息。
    * Using index condition
        有些搜索条件中虽然出现了索引列，但却不能使用到索引，比如下边这个查询：
> select * from s1 where key1 > 'z' and key1 like '%a';

其中的key1 > 'z'可以使用索引，但是key1 like '%a'却无法使用到索引，在以前版本的MySQL，是按照下边步骤来执行这个查询的：
    * 先根据key1 > 'z'，从二级索引idx_key1中获取到对应的二级索引记录
    * 根据上一步得到的二级索引记录中的主键值进行回表，找到完整的用户记录再检测该记录是否符合key1 like '%a' 这个条件，将符合条件的记录加入到最后的结果集
但是虽然key1 like '%a'不能组成范围区间参与range访问方法的执行，但这个条件毕竟只涉及到了key1列，所以MySQL将步骤改进了一下：
    * 先根据key1 > 'z'这个条件，定位到二级索引idx_key1中对应的二级索引记录
    * 对于指定的二级索引记录，先不回表，而是先检测一下改记录是否满足key1 like '%a'这个条件，如果这个条件不满足，则该二级索引记录不回表
    * 对于满足key1 like '%a'这个条件的二级索引记录执行回表操作

**回表操作其实是一个随机IO，比较耗时**，上述修改虽然只改进了一点，但是省去好多回表操作的成本。这个改进称之为**索引条件下推(Index Condition Pushdown)**

如果在查询语句的执行过程中将要使用**索引条件下推**这个特性，在Extra列中将会显示Using index condition

* Using where

    * 当我们使用全表扫描来执行对某个表的查询，并且该语句的WHERE子句中有针对该表的搜索条件时
    * 当使用索引访问来执行对某个表的查询，并且该语句的WHERE子句有除了该索引包含的列之外的其他搜索条件时，在Extra列中中也会提示Using Where
    
* Using join buffer (Block Nested Loop)

在链接查询执行过程中，当被驱动表不能有效的利用索引加快访问速度，MySQL一般会为其分配一块名为join buffer的内存块来加快查询速度，也就是**基于块的嵌套循环算法**

* Using filesort

有一些情况下对结果集中的记录进行排序是可以使用到索引的，但是很多情况下排序操作无法使用到索引，只能在内存（记录较少的时候）或者磁盘中进行排序，MySQL
把这种在内存中或磁盘上进行排序的方式统称为文件排序（filesort）,如果查询中需要使用filesort的方式进行排序的记录非常多，那么这个过程是很耗费性能的，
最好将使用文件排序的执行方式改为使用索引进行排序。

* Using temporary

在许多查询的执行过程中，MySQL可能会借助临时表来完成一些功能，比如去重、排序之类的，比如在执行包含distinct、group by、Union等子句的查询过程中，
如果不能有效利用索引来完成查询，MySQL很有可能寻求通过建立内部的临时表来执行查询。

**MySQL会在包含GROUP BY子句的查询中默认添加上order by子句，如果并不想group by子句的查询进行排序，需要显式的写上order by null**
这下执行计划中就没有Using filesort的提示了，也就意味着执行查询时可以省去对记录进行文件排序的成本。

#### Json格式的执行计划
Explain语句输出中缺少了一个衡量执行计划好坏的重要属性：成本。MySQL为我们提供一种查看某个执行计划花费的成本的方式：

* 在explain单次和真正的查询语句中间加上FORMAT=JSON

{
"query_block": {
"select_id": 1,
"cost_info": {
"query_cost": "11593.60"
},
"table": {
"table_name": "gp_order",
"access_type": "ALL",
"rows_examined_per_scan": 52028,
"rows_produced_per_join": 5202,
"filtered": "10.00",
"cost_info": {
"read_cost": "10553.04",
"eval_cost": "1040.56",
"prefix_cost": "11593.60",
"data_read_per_join": "88M"
},
"used_columns": [
"id",
"order_no",
"wait_no",
"user_id",
"employee_id",
"store_id",
"store_name",
"serve_id",
"serve_name",
"serve_realated_type",
"sub_serve_id",
"sub_serve_name",
"color_card_id",
"color_card_name",
"hair_length",
"hair_quality",
"img",
"img_list",
"price",
"pay_price",
"use_balance_price",
"use_wechat_price",
"expected_price",
"pay_img",
"pay_type",
"coupon_id",
"coupon_price",
"hongbao_ids",
"hongbao_price",
"vip_discount_price",
"vip_use_first_price",
"vip_level",
"group_buy",
"time_card_id",
"time_card_name",
"user_time_card_id",
"user_time_card_price",
"type",
"comment_id",
"comment_status",
"order_time",
"start_time",
"end_time",
"expected_duration",
"duration",
"cancel_reason",
"cancel_type",
"sub_coming",
"sub_coming_time",
"sub_wait_pay",
"sub_wait_pay_time",
"sub_pass",
"sub_pass_time",
"formid",
"transfer",
"transfer_time",
"material_extract",
"material_extract_time",
"passed",
"passed_time",
"deleted",
"status",
"status_time",
"update_time",
"create_time",
"refund_id",
"hair_length_options",
"fake"
],
"attached_condition": "(`gipin`.`gp_order`.`price` = 2)"
}
}
}