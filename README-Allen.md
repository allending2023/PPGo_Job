
## 补充说明

1. 增加 pgsql 数据库支持
2. 增加 Windows service支持

### 编译

在Windows下，执行./package.sh -a amd64 -p windows -v v2.8.0
需要使用git bash的终端类型执行

先

1. 执行 go env -w GO111MODULE=on [ 可选， 若存在go.mod就不用执行 ]
2. 执行 go mod edit -module=example.com/mod [ 可选， 若存在go.mod就不用执行 ]
3. 执行 go mod tidy
4. ./package.sh -a amd64 -p windows -v v2.8.0

编译时，因git bash不支持 zip命令，无法完成多文件打包


### 安装服务

#### 直接安装

在管理员命令行模式下：执行

./PPGo_Job.exe install | Uninstall | start | stop

#### 通过Bat安装

install.bat
sc create PPGo_Job binPath=D:/programs/PPGo_Job2.8/PPGo_Job.exe type=share start=auto

uninstall.bat
sc delete PPGo_Job
pause


