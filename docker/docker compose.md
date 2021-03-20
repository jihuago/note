## Docker Compose
Docker Compose 是Docker官方编排项目之一，负责快速在集群中部署分布式应用。
Dockerfile可以让用户管理一个单独的应用容器；而Compose则允许用户在一个模板YAML格式定义一组相关联的应用容器，例如一个Web
服务容器再加上后端的数据库服务容器等。
Docker Compose由Python编写，实际上调用了Docker提供的API来实现

### 安装
* PIP安装
```html
$ sudo pip install -U docker-compose
```

* bash补全命令
```html
$ curl -L https://raw.githubusercontent.com/docker/compose/1.27.4/contrib/completion/bash/docker-compose > /etc/bash_completion.d/docker-compose
```

### 使用
1. 术语
* 服务：一个应用容器，实际上可以运行多个相同镜像的实例
* 项目：由一组关联的应用容器组成的一个完整业务单元

