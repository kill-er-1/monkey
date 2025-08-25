// 准备 (Setup): 定义输入字符串和一份详细的“标准答案”（期望的 Token 序列）。
// 执行 (Execution): 创建词法分析器实例，并循环调用它的 NextToken() 方法来生成实际的 Token。
// 验证 (Verification): 在每次循环中，将实际生成的 Token 与“标准答案”中对应的 Token 进行比较。
// 报告 (Reporting): 如果有任何不匹配，就使用测试框架的 t.Fatalf 函数报告一个详细的错误并立即失败。如果整个循环顺利完成都没有失败，则测试通过。
package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5; 
	let ten = 10;

	 let add = fn(x, y)
	  { 
		 x + y; 
	 }; 

	 let result = add(five, ten); 
	 `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"}, {token.IDENT, "five"}, {token.ASSIGN, "="}, {token.INT, "5"}, {token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "ten"}, {token.ASSIGN, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "add"}, {token.ASSIGN, "="}, {token.FUNCTION, "fn"}, {token.LPAREN, "("}, {token.IDENT, "x"}, {token.COMMA, ","}, {token.IDENT, "y"}, {token.RPAREN, ")"}, {token.LBRACE, "{"}, {token.IDENT, "x"}, {token.PLUS, "+"}, {token.IDENT, "y"}, {token.SEMICOLON, ";"}, {token.RBRACE, "}"}, {token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "result"}, {token.ASSIGN, "="}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "five"}, {token.COMMA, ","}, {token.IDENT, "ten"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"}, {token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - expected[%q],got[%q]",i,tt.expectedType,tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - expected[%q],got[%q]",i,tt.expectedLiteral,tok.Literal)
		}
// 		%q 作用：在字符串的两端加上双引号 "。
// 对字符串内部的特殊字符进行转义 (escape)，使其能够安全地显示。
	}
}
