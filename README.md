# go_dgut_yqdk
**作者: ChovyChen(陈世旗), huoayi(霍一慷)**

**特别鸣谢: Charlexchen(陈灏嵘)**

#### 系统使用
仅使用本系统可访问[dgut自动打卡](http://keyvon.club/yqdk/yqfk_ui.html)

#### 介绍
Dgut自动打卡系统，`GO`语言开发

#### 特点
使用了协程机制, 不管多少账号都可以同时打卡

#### 使用说明
1. `git clone`
2. 在`main.go`同级目录下新建`config.yaml`, 输入账号密码, 模板如下:
```yaml
# 主要配置不可忽略
# 协程最大数量
func_max: 10
# 数据源, 0 = 文件, 1 = 数据库
data_source: 0

# data_source = 0 时使用以下配置
UserAccount:
  - username: "your dgut username, e.g.:201841413404"
    password: "your dgut password, e.g.:xxxxxxxxxxxx"
# 如果想多个账号同时打卡，把注释去掉即可
# - username: "your dgut username e.g.:201841413404"
#   password: "your dgut password e.g.:xxxxxxxxxxxx"
# - username: "your dgut username e.g.:201841413404"
#   password: "your dgut password e.g.:xxxxxxxxxxxx"
# - username: "your dgut username e.g.:201841413404"
#   password: "your dgut password e.g.:xxxxxxxxxxxx"
# - username: "your dgut username e.g.:201841413404"
#   password: "your dgut password e.g.:xxxxxxxxxxxx"

# data_source = 1 时使用以下配置
## DB 信息
#DB:
#  # 数据库主机地址
#  DB_host: "127.0.0.1"
#  # 数据库端口
#  DB_port: "3306"
#  # 数据库名称
#  DB_name: "db_dk"
#  # 数据表名
#  DB_table: "tb_dk"
#  # 账号
#  DB_username: "root"
#  # 密码
#  DB_password: "123456"
``` 
3. 命令行
```shell
-- windows上执行, 打开cmd后输入以下命令
> run.sh // 选择Git Bash运行即可
-- linux / macOX 执行
> run.sh // 直接执行即可
```
4. `log`都输出在`cmd`/终端上, 若查看往期日志, 可查看`log`文件夹
5. 可根据自己需求定制定时任务, 如`linux`上可用`crontab`

#### 目录结构
- `yqdk` 主目录
    - `base` 实体类与工具类
    - `log` 日志文件
    - `src` 主要代码逻辑
    - `main.go` 项目入口
    - `run.sh` 启动脚本