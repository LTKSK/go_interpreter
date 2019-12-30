package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y...
	INT   = "INT"

	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	COMMA     = ","
	SEMICOLON = ";"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// user定義のものがあったらIDENTを、fnやletならそれらを返す
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
