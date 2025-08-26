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
	l.skipWhitespace()
	switch l.ch {
		case '=':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.ASSIGN, l.ch)
			}
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
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '!':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.BANG, l.ch)
			}
		case '*':
			tok = newToken(token.ASTERISK, l.ch)
		case '/':
			tok = newToken(token.SLASH, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case '<':
			tok = newToken(token.LT, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(l.ch) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok //直接返回
			} else if isDigit(l.ch) {
				tok.Literal = l.readNumber()
				tok.Type = token.INT
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}
	l.readChar()
	return tok
}

//判断是否为字母的函数 ？标识符与关键词都是字母，_ 标识符中间可以有下划线
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]//左闭右开
}

//跳过空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

//判断数字
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

//读取数字
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

//窥探下一个字符但是不移动position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}