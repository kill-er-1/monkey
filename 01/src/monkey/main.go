package main

import (
	"fmt"         // 格式化输出包
	"monkey/repl" // 自定义REPL包
	"os"          // 操作系统接口包
	"os/user"     // 用户信息包
)

// main函数是程序的入口点
func main() {
	// 获取当前系统用户信息
	user, err := user.Current()

	// 检查是否获取用户信息时出错
	if err != nil {
		panic(err) // 如果出错，终止程序并打印错误信息
	}

	// 打印个性化欢迎信息，包含用户名
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)

	// 打印使用提示
	fmt.Printf("Feel free to type in commands\n")

	// 启动REPL交互式环境
	// os.Stdin: 标准输入（键盘输入）
	// os.Stdout: 标准输出（屏幕输出）
	repl.Start(os.Stdin, os.Stdout)
}
