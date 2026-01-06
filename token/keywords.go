//go:generate stringer -type=TokenType

package token

const (
	ILLEGAL TokenType = iota
	EOF
	// Identifiers + literals
	IDENT
	INT

	// Operators
	ASSIGN   // =
	PLUS     // +
	MINUS    // -
	BANG     // !
	ASTERISK // *
	SLASH    // /
	LT       // <
	GT       // >
	EQ		// ==
	NOT_EQ	// !=

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)