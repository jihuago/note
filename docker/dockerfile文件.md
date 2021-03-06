### dockerfile
Dockerfile用来构建镜像的文本文件，文本内容包含了一条条构建镜像锁需的指令和说明。

* .dockerignore文件
.dockerignore文件中写的是需要排除的文件或路径
  
* Dockerfile文件指令解释

FROM 定制的镜像基于什么，如FROM nginx  这里的nginx就是定制需要的基础镜像

RUN  用于执行后面跟着的命令行命令。shell格式：
```html
    RUN <命令行命令>
    # 命令行命令等同于，在终端操作的shell命令
```

exec格式：
```html
RUN ["可执行文件", "参数1", "参数2"]
```
RUN ["./test.php", "dev", "offline"]等价于 RUN ./test.php dev offline

### 注意
Dockerfile的指令每执行一次都会在docker上新建一层。所以过多无意义的层，会造成镜像过大。例如：
```html
FROM centos
RUN yum install wget
RUN wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz"
RUN tar -xvf redis.tar.gz
```
以上执行会创建3层镜像。可简化为以下格式：
```html
FROM centos
RUN yum install wget \
&& wget -O redis.tar.gz  "http://download.redis.io/releases/redis-5.0.3.tar.gz"
&& tar -xvf redis.tar.gz
```
## 构建镜像
在Dockerfile文件的存放目录下，执行构建动作。
```html
docker build -t nginx:v3 .
```
nginx:v3 镜像名称:镜像标签

### 上下文路径
. 是上下文路径
> docker build -t nginx:v3 .

上下文路径是指docker在构建镜像，有时候想要使用到本机的文件（比如复制），docker build 命令得知这个路径后，会将路劲下的所有内容打包。
如未说明最后一个参数，那么默认上下文路径就是Dockerfile所在的位置。
注意：上下文路径不要放无用的文件，因为会一起打包发送给docker引擎，如果文件过多会造成过程缓慢。

## 指令详解
* FROM 
  格式为`FROM <image>`或FROM <image>:<tag>
  第一条指令必须为FROM指令。如果在同一个Dockerfile创建多个镜像，可以使用多个FROM指令（每个镜像一次）

* MAINTAINER
    格式 MAINTAINER <name> 指定维护者信息
  
* RUN
    格式为 RUN <command> 或 RUN ["executable", "param1", "param2"]
每条RUN指令将在当前镜像基础上执行指定命令，并提交为新的镜像。当命令较长可以使用\来换行。

* CMD
    格式：`CMD ["executable", "param1", "param2"]` 使用exec执行，推荐
    每个Dockerfile只能有一条CMD命令。如果指定了多条命令，只有最后一条被执行。
    如果用户启动容器时指定了运行的命令，则会覆盖掉CMD指定的命令。

* COPY
格式： `COPY <src> <dest>`
  复制本地主机<src>（为Dockerfile所在目录的相对路径）到容器中的<dest>
  
* VOLUME
格式为`VOLUME ["/data"]`
  创建一个可以从本地主机或其他容器挂载的挂载点，一般用来存放数据库和需要保持的数据等
  
* USER
格式为 USER daemon
  指定运行容器时的用户名或UID，后续的RUN 也会使用指定用户
  
* WORKDIR
`WORKDIR /path/to/workdir`
  为后续的RUN、CMD、ENTRYPONT指令配置工作目录
  可以使用多个WORKDIR指令，后续命令如果参数是相对命令，则会基于之前命令指定的路径。
  ```html
WORKDIR /a
WORKDIR b
WORKDIR c
RUN pwd
```
则最终路径为/a/b/c
```
* ONBUILD
`ONBUILD [INSTRUCTION]`
配置当所创建的镜像作为其他新创建镜像的基础镜像时，所执行的操作指令。

* ENV
格式有两种：
  `ENV <key> <value>`
  `ENV <key1>=<value1> ...`
  
ENV指令就是设置环境变量，后面的其他指令，如RUN，可以直接使用这里定义的环境变量。
```html
ENV VERSION=1.0 DEBUG=on \
    NAME="Happy Feet"

# 使用
RUN $VERSION
```

* ARG 构建参数
格式 `ARG <参数名><=默认值>`
  
构建参数和ENV的效果一样，都是设置环境变量。不同的是，ARG所设置的构建环境的环境变量，在将来容器运行时不会存在这些环境变量。

ARG指令有生效范围，如果在FROM指令之前指定，那么只能用于FROM指令中。
```html
# 只在FROM中生效
ARG DOCKER_USERNAME=library
FROM ${DOCKER_USERNAME}/alpine
#要想在FROM之后使用，必须再次指定
ARG DOCKER_USERNAME=library
RUN set -x ; echo ${DOCKER_USERNAME}

```

* EXPORE 暴露端口
格式为`EXPORE <端口1> [<端口2>]`
  EXPORE指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务。
  在DOckerfile中写入这样的声明有两个好处，一个是帮助镜像使用者理解这个镜像服务的守护端口，以方便映射；
  另一个好处则是在运行时使用随机端口映射时，也就是`docker run -P`时，会自动随机映射EXPOSE的端口
  

