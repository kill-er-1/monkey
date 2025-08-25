package lexer

//lexer词法分析器

//1.定义lexer的结构

//2.实现New方法 创建实例
import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 当前字符的位置
	readPosition int  // 下一个字符的位置
	ch          byte // 当前字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() //让lexer准备好读取输入
	return  l
}

func (l *Lexer) readChar()  {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII 码中的 NUL 字符，表示文件结束
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

//定义一个创建新token的函数
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

//定义读取下一个token的函数
func (l *Lexer)NextToken()token.Token {
	var tok token.Token
	switch l.ch {
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
	}
	l.readChar()
	return tok
}