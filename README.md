#Hawk Project

##1 目录结构

* conf: 配置文件目录
* src: 源代码目录
    * eye: 鹰眼服务,接受原始信号上报
    * brain: 鹰脑服务,运行定位算法,存储定位结果

##2 编译方法
* 进入hawk目录后执行make命令

##3 执行方法
* 执行eye程序,指定服务监听的本地IP和端口:

    `./bin/eye -host 0.0.0.0 -port 3333`
