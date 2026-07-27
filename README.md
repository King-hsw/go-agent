# go-agent README
配套 ha-pc-switch 的 Windows 端 Go 监听程序，接收 HomeAssistant 下发的局域网关机指令，执行 Windows 系统关机操作。
> 搭配 HA 插件 ha-pc-switch 使用，单独运行无任何作用。

## 项目简介
`go-agent` 轻量 Go 后台监听程序，部署在被控 Windows 电脑本地，功能单一纯粹：
1. 监听指定局域网端口，接收 HA `ha-pc-switch` 下发的关机网络指令
2. 收到合法指令后，执行 Windows 系统关机命令
3. 轻量化无多余依赖，支持设置开机自启后台静默运行

## 功能清单
- ✅ 局域网端口监听，接收 HA 下发关机指令
- ✅ 收到指令自动执行 Windows 系统关机
- ✅ 轻量编译，单exe文件，无需Go运行环境
- ✅ 支持后台静默运行，可配置Windows开机自启
- ✅ 简单端口自定义配置，与HA插件端口保持一致即可

## 前置依赖与环境
1. 系统：仅支持 Windows（关机指令适配Windows系统命令）
2. 配套 HA 插件：`ha-pc-switch`
3. 网络：电脑与 HomeAssistant 设备处于同一局域网
4. 防火墙：放行程序监听端口，允许局域网入站访问

## 编译与安装部署
### 方式1：直接使用预编译exe（推荐）
1. 进入项目 Releases 下载已编译好的 Windows exe 程序
2. 新建文件夹存放exe，避免系统盘临时目录
3. 配置监听端口（两种方式任选其一）
   - 方式A：启动命令指定端口
     ```cmd
     go-agent.exe -port 8090
     ```
   - 方式B：修改配置文件指定端口
4. 放行Windows防火墙对应端口入站规则
5. 设置程序开机自启，后台静默运行

### 方式2：本地Go源码自行编译
1. 本地安装 Go 语言环境（Go1.18+）
2. 克隆项目源码
   ```bash
   git clone https://github.com/King-hsw/go-agent.git
   cd go-agent
   ```
3. Windows交叉编译生成exe程序
   ```bash
    GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o pc-agent.exe
   ```
4. 编译完成后得到单文件 `go-agent.exe`，复制到Windows电脑部署

## 配套 HA 插件配置注意事项
1. `go-agent` 监听端口必须与 `ha-pc-switch` 集成内填写的客户端端口完全一致
2. HA插件填写被控电脑静态内网IP，保证局域网可正常访问该端口
3. WOL开机无需本程序，仅远程关机功能依赖 `go-agent` 后台常驻运行

## 开机自启配置教程
### 方法1：Windows 任务计划程序（推荐，稳定后台运行）
1. 打开「任务计划程序」，创建基本任务
2. 触发器选择「计算机启动时」
3. 操作选择「启动程序」
4. 程序路径选择 `go-agent.exe`，参数填写 `-port 8090`（替换你的端口）
5. 勾选「不管用户是否登录都要运行」，完成配置

### 方法2：启动文件夹快捷方式
1. Win+R 输入 `shell:startup` 打开开机启动文件夹
2. 创建 `go-agent.exe` 快捷方式，右键快捷方式-属性，目标末尾添加端口参数
3. 确定保存，电脑重启后自动启动监听程序

## 常见问题
### 1. HA下发关机指令，电脑无反应
1. `go-agent` 程序未启动，未后台常驻运行
2. HA配置端口与程序监听端口不一致
3. Windows防火墙拦截监听端口，HA无法发送指令
4. 电脑IP变更，HA插件内IP未同步修改

### 2. 程序启动后闪退
1. 端口被其他软件占用，更换未占用端口重新启动
2. 无管理员权限运行，右键exe选择「以管理员身份运行」

### 3. 电脑重启后程序不自动启动
1. 自启配置设置错误，优先使用任务计划程序配置自启
2. 杀毒软件拦截程序开机启动，添加程序文件夹至杀毒白名单

## 项目目录结构
```
go-agent/
├── main.go                 # Go程序主入口源码
├── go.mod                  # Go依赖管理文件
├── go.sum
├── .gitignore
└── README.md               # 客户端说明文档
```

## 配套 HomeAssistant 插件仓库
HA 集成插件源码仓库：
`https://github.com/King-hsw/ha-pc-switch`

## 开源协议
本项目开源免费，仅用于个人智能家居学习使用，禁止商用。

## 反馈与问题
使用过程中出现监听异常、关机失效等问题，可在本仓库提交 Issues，持续迭代维护。