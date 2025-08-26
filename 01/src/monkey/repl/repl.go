package repl

//REPL 是指Read-Eval-Print Loop（读取（求值求打印循环），

import (
	"bufio"        // 缓冲I/O，用于读取输入
	"fmt"          // 格式化输出
	"io"           // I/O接口
	"monkey/lexer" // 自定义词法分析器
	"monkey/token" // 自定义Token定义
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
		scannered := scanner.Scan()
		if !scannered {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok) // %+v 对于结构体包含字段名
		}
	}
}
