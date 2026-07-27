package main // 声明当前文件属于 main 包，Go 程序入口必须是 main 包

import (
	"flag" // 导入 flag 包，用于解析启动时传入的命令行参数
	"log"
	"pc-agent/system"

	"net/http" // 导入 http 包，提供 HTTP 状态码等功能

	"github.com/gin-gonic/gin" // 导入 Gin 框架，用于快速创建 HTTP 服务
	// 导入自己项目里的 system 包，用于执行关机操作
)

// Response 定义统一返回给客户端的 JSON 数据结构
type Response struct {
	Success bool   `json:"success"` // success 字段表示请求是否成功
	Code    int    `json:"code"`
	Message string `json:"message"` // message 字段表示返回提示信息

}

func main() { // Go 程序入口函数，程序从这里开始执行

	// 定义一个名为 port 的启动参数
	// 使用方式：pc-agent.exe --port 8080
	// 如果没有传递参数，则默认使用 5666
	port := flag.String(
		"port", // 参数名称，对应命令行里的 --port

		"5666", // 参数默认值

		"HTTP server port", // 参数说明，会显示在 --help 中

	)

	// 解析命令行参数
	// 执行后，用户传入的 --port 值会覆盖默认值
	flag.Parse()

	// 创建 Gin 路由引擎
	// gin.Default() 会创建默认 Engine
	// 同时开启 Logger 和 Recovery 中间件
	r := gin.Default()

	// 注册 POST 请求接口
	// 当客户端请求 POST /shutdown 时执行下面的函数
	r.POST(
		"/shutdown",

		func(c *gin.Context) {

			// 启动一个新的 goroutine 异步执行关机
			// 不阻塞当前 HTTP 请求
			// HTTP 会立即返回响应
			go func() {

				// 调用 system 包里的 Shutdown 方法
				// 执行实际关机逻辑
				err := system.Shutdown()

				if err != nil {
					log.Println(
						"shutdown error:",
						err,
					)
				}

			}()
			// 返回 HTTP 200 状态码
			// 并返回 JSON 数据
			c.JSON(
				http.StatusOK,
				Response{
					Success: true, // 表示接口调用成功
					Code:    200,
					Message: "电脑将在60秒后关机", // 返回提示消息
				},
			)

		},
	)

	// 拼接监听地址
	// 例如：
	// port=5666
	// addr 就等于 ":5666"
	addr := ":" + *port

	// 启动 HTTP 服务
	// Gin 会监听指定端口
	// 例如 http://localhost:5666
	r.Run(addr)

}
