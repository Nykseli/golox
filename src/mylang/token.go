package main

// TokenType type for tokens
type TokenType int

// costs are token  types
const (

	// TokenLeftParen is type for '('
	TokenLeftParen = 0
	// TokenRightParen is type for ')'
	TokenRightParen = 1
	// TokenLeftBrace is type for '{'
	TokenLeftBrace = 2
	// TokenRightBrace is type for '}'
	TokenRightBrace = 3
	// TokenComma is type for ','
	TokenComma = 4
	// TokenDot is type for '.'
	TokenDot = 5
	// TokenMinus is type for '-'
	TokenMinus = 6
	// TokenPlus is type for '+'
	TokenPlus = 7
	// TokenSemicolon is type for ';'
	TokenSemicolon = 8
	// TokenSlash is type for '/'
	TokenSlash = 9
	// TokenStar is type for '*'
	TokenStar = 10

	// TokenBang is type for '!'
	TokenBang = 11
	// TokenBangEqual is type for '!='
	TokenBangEqual = 12
	// TokenEqual is type for '='
	TokenEqual = 13
	// TokenEqualEqual is type for '=='
	TokenEqualEqual = 14
	// TokenGreater is type for '>'
	TokenGreater = 15
	// TokenGreaterEqual is type for '>='
	TokenGreaterEqual = 16
	// TokenLess is type for '<'
	TokenLess = 17
	// TokenLessEqual is type for '<='
	TokenLessEqual = 18

	// TokenIdentifier is type for varaible names
	TokenIdentifier = 19
	// TokenString is type for strings
	TokenString = 20
	// TokenNumber is type for numbers
	TokenNumber = 21

	// TokenAnd is type for and keyword
	TokenAnd = 22
	// TokenClass is type for class keyword
	TokenClass = 23
	// TokenElse is type for else keyword
	TokenElse = 24
	// TokenFalse is type for false keyword
	TokenFalse = 25
	// TokenFor is type for for keyword
	TokenFor = 26
	// TokenFun is type for fun keyword
	TokenFun = 27
	// TokenIf is type for if keyword
	TokenIf = 28
	// TokenNil is type for nil keyword
	TokenNil = 29
	// TokenOr is type for or keyword
	TokenOr = 30
	// TokenPrint is type for print keyword
	TokenPrint = 31
	// TokenReturn is type for return keyword
	TokenReturn = 32
	// TokenSuper is type for super keyword
	TokenSuper = 33
	// TokenThis is type for this keyword
	TokenThis = 34
	// TokenTrue is type for true keyword
	TokenTrue = 35
	// TokenVar is type for var keyword
	TokenVar = 36
	// TokenWhile is type for while keyword
	TokenWhile = 37

	// TokenError is type for error tokens
	TokenError = 38

	// TokenEOF is type for end of file token
	TokenEOF = iota
)

// Token holds info for each token
type Token struct {
	Type   TokenType
	Value  string
	Length int
	Line   int
}
