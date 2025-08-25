package token

//1.一个token得有类型，还得有字面量

//2.一个token的类型是用来区分不同种类的词法单元的

//3.一个token的字面量是它所表示的具体值

//4.一个token的类型和字面量共同构成了一个完整的词法单元

//5.一个token可以被认为是一个不可分割的基本单位

//6.类型可以用常量来包含所有的token类型，所以字符

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//标识符 + 字面量
	IDENT   = "IDENT"
	INT     = "INT"

	//分隔符
	COMMA   = ","
	SEMICOLON = ";"
	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"

	//运算符
	ASSIGN = "="
	PLUS   = "+"

	//关键字
	LET    = "let"
	FUNCTION = "function"
)
