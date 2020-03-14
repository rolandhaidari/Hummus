package lexer

type TokenType uint8

type Token struct {
	Value string    `json:"value"`
	Type  TokenType `json:"type"`
	Line  uint      `json:"line"`
}

const (
	OPEN_BRACE        TokenType = 1
	CLOSE_BRACE       TokenType = 2
	IDENTIFIER        TokenType = 3
	INT               TokenType = 4
	FLOAT             TokenType = 5
	STRING            TokenType = 6
	BOOL              TokenType = 7
	ATOM              TokenType = 8
	IDENTIFIER_DEF    TokenType = 9
	IDENTIFIER_FN     TokenType = 10
	IDENTIFIER_STRUCT TokenType = 11
	IDENTIFIER_FOR    TokenType = 12
)