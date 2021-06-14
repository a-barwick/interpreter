package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	FSLASH   = "/"

	BANG  = "!"
	EQ    = "=="
	NOTEQ = "!="
	GT    = ">"
	LT    = "<"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func GetIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
