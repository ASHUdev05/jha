package token

type TokenType int

type Token struct {
	Type TokenType
	Literal string
}