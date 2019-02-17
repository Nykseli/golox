package main

// FileEOF is constant for end of file
// use 0 instead of -1 to work with unsigned values
const FileEOF = 0

// Scanner is struct for scanning source code to tokens
type Scanner struct {
	Source   string
	StartPos int

	CurrentPos int
	Line       int
}

var scanner = Scanner{}

// InitScanner initilizes the scanner for reading
func InitScanner(source string) {
	scanner.Source = source
	scanner.StartPos = 0
	scanner.CurrentPos = 0
	scanner.Line = 1
}

// ScanToken returns the next token from source code
func ScanToken() Token {
	skipWhitespace()

	scanner.StartPos = scanner.CurrentPos

	if isAtEnd() {
		return makeToken(TokenEOF)
	}

	c := advance()
	if isAlpha(c) {
		return identifier()
	}

	if isDigit(c) {
		return number()
	}

	switch c {
	case '(':
		return makeToken(TokenLeftParen)
	case ')':
		return makeToken(TokenRightParen)
	case '{':
		return makeToken(TokenLeftBrace)
	case '}':
		return makeToken(TokenRightBrace)
	case ';':
		return makeToken(TokenSemicolon)
	case ',':
		return makeToken(TokenComma)
	case '.':
		return makeToken(TokenDot)
	case '-':
		return makeToken(TokenMinus)
	case '+':
		return makeToken(TokenPlus)
	case '/':
		return makeToken(TokenSlash)
	case '*':
		return makeToken(TokenStar)
	case '!':
		if match('=') {
			return makeToken(TokenBangEqual)
		}
		return makeToken(TokenBang)
	case '=':
		if match('=') {
			return makeToken(TokenEqualEqual)
		}
		return makeToken(TokenEqual)
	case '<':
		if match('=') {
			return makeToken(TokenLessEqual)
		}
		return makeToken(TokenLess)
	case '>':
		if match('=') {
			return makeToken(TokenGreaterEqual)
		}
		return makeToken(TokenGreater)
	case '"':
		return stringToken()

	}

	return errorToken("Unexpected character.")
}

func isAlpha(c uint8) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func isAtEnd() bool {
	return scanner.Source[scanner.CurrentPos] == FileEOF
}

func advance() uint8 {
	scanner.CurrentPos++
	return scanner.Source[scanner.CurrentPos-1]
}

func peek() uint8 {
	return scanner.Source[scanner.CurrentPos]
}

func peekNext() uint8 {
	if isAtEnd() {
		return FileEOF
	}
	return scanner.Source[scanner.CurrentPos+1]
}

func match(expected uint8) bool {
	if isAtEnd() {
		return false
	}
	if scanner.Source[scanner.CurrentPos] != expected {
		return false
	}
	scanner.CurrentPos++
	return true
}

func makeToken(_type TokenType) Token {
	var token = Token{}
	token.Type = _type
	token.Length = scanner.CurrentPos - scanner.StartPos
	token.Value = scanner.Source[scanner.StartPos:scanner.CurrentPos]
	token.Line = scanner.Line

	return token
}

func errorToken(message string) Token {
	var token = Token{}
	token.Type = TokenError
	token.Value = message
	token.Length = len(message)
	token.Line = scanner.Line

	return token
}

func skipWhitespace() {
	for {
		c := peek()

		if c == ' ' || c == '\r' || c == '\t' {
			advance()
		} else if c == '\n' {
			scanner.Line++
			advance()
		} else if c == '/' {
			if peekNext() == '/' {
				// A comment goes until the end of the line.
				for peek() != '\n' && !isAtEnd() {
					advance()
				}
				// Finally advance once more to consume the newline character
				advance()
			} else {
				return
			}
		} else {
			return
		}
	}
}

func checkKeyword(start int, length int, rest string, _type TokenType) TokenType {
	if scanner.CurrentPos-scanner.StartPos == start+length &&
		string(scanner.Source[scanner.StartPos+start:scanner.StartPos+length+start]) == rest {
		return _type
	}

	return TokenIdentifier
}

func identifierType() TokenType {
	switch scanner.Source[scanner.StartPos] {
	case 'a':
		return checkKeyword(1, 2, "nd", TokenAnd)
	case 'c':
		return checkKeyword(1, 4, "lass", TokenClass)
	case 'e':
		return checkKeyword(1, 3, "lse", TokenElse)
	case 'f':
		if scanner.CurrentPos-scanner.StartPos > 1 {
			switch scanner.Source[scanner.StartPos+1] {
			case 'a':
				return checkKeyword(2, 3, "lse", TokenFalse)
			case 'o':
				return checkKeyword(2, 1, "r", TokenFor)
			case 'u':
				return checkKeyword(2, 1, "n", TokenFun)
			}
		}
		break
	case 'i':
		return checkKeyword(1, 1, "f", TokenIf)
	case 'n':
		return checkKeyword(1, 2, "il", TokenNil)
	case 'o':
		return checkKeyword(1, 1, "r", TokenOr)
	case 'p':
		return checkKeyword(1, 4, "rint", TokenPrint)
	case 'r':
		return checkKeyword(1, 5, "eturn", TokenReturn)
	case 's':
		return checkKeyword(1, 4, "uper", TokenSuper)
	case 't':
		if scanner.CurrentPos-scanner.StartPos > 1 {
			switch scanner.Source[scanner.StartPos+1] {
			case 'h':
				return checkKeyword(2, 2, "is", TokenThis)
			case 'r':
				return checkKeyword(2, 2, "ue", TokenTrue)
			}
		}
		break
	case 'v':
		return checkKeyword(1, 2, "ar", TokenVar)
	case 'w':
		return checkKeyword(1, 4, "hile", TokenWhile)
	}

	return TokenIdentifier
}

func identifier() Token {
	for isAlpha(peek()) || isDigit(peek()) {
		advance()
	}

	return makeToken(identifierType())
}

func number() Token {
	for isDigit(peek()) {
		advance()
	}

	if peek() == '.' && isDigit(peekNext()) {
		// Consume the "."
		advance()
		for isDigit(peek()) {
			advance()
		}
	}

	return makeToken(TokenNumber)
}

func stringToken() Token {
	for peek() != '"' && !isAtEnd() {
		if peek() == '\n' {
			scanner.Line++
		}
		advance()
	}

	if isAtEnd() {
		return errorToken("Unterminated string.")
	}

	// Consume the closing '"'
	advance()
	return makeToken(TokenString)
}
