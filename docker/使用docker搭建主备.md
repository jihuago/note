## 使用docker搭建主备
1主，2备；mysql5.7

### 过程
* 拉取mysql镜像
> docker pull mysql:5.7

* 运行容器（创建主库）
> docker run -itd --name mysql-master -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7

* 创建备库
> docker run -itd --name mysql-slave00 -p 3308:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7

> docker run -itd --name mysql-slave01 -p 3309:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7

参数说明：
    -p 3306:3307  映射容器服务的3306端口到宿主机3307端口
    -itd 让docker启动后能一直运行
   
* 删除重复的容器，如果需要
> docker container rm mysql-master

出现类似以下问题：
> docker: Error response from daemon: Conflict. The container name "/mysql-master" is already in use by container

* 配置主从

* 将mysql-master的配置文件复制到宿主，然后修改配置文件后，将修改后的配置文件覆盖docker里的配置文件

> docker cp mysql-master:/etc/mysql/mysql.conf.d/mysqld.cnf /usr/local/mysql/master/mysqld.cnf

> docker cp /usr/local/mysql/master/mysqld.cnf mysql-master:/etc/mysql/mysql.conf.d/mysqld.cnf

* 重启mysql-master

> docker restart mysql-master

### 参考链接
> https://my.oschina.net/u/3773384/blog/1810111

