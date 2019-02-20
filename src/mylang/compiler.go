package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Parser keeps track of Tokens we are turning into bytecode
type Parser struct {
	Current   Token
	Previous  Token
	HadError  bool
	PanicMode bool
}

// Precedence is for tracking what operatios are emited first
// Higher first, lower last
type Precedence int

const (
	// PrecNone is is the last thing emited
	PrecNone Precedence = iota
	// PrecAssignment is for =
	PrecAssignment Precedence = iota
	// PrecOr is for or
	PrecOr Precedence = iota
	// PrecAnd is for and
	PrecAnd Precedence = iota
	// PrecEquality is for == !=
	PrecEquality Precedence = iota
	// PrecComparison is for < > <= >=
	PrecComparison Precedence = iota
	// PrecTerm is for + -
	PrecTerm Precedence = iota
	// PrecFactor is for * /
	PrecFactor Precedence = iota
	// PrecUnary is for ! - +
	PrecUnary Precedence = iota
	// PrecCall is for . () []
	PrecCall Precedence = iota
	// PrecPrimary is going to be defined later
	PrecPrimary Precedence = iota
)

// ParseFn is used for functions in ParseRule struct
// This way we can write easily the rule table
type ParseFn func()

// ParseRule is used for rule table
type ParseRule struct {
	Prefix     ParseFn
	Infix      ParseFn
	Precedence Precedence
}

// rules contains parsing rules. initialized in initCompiler()
var rules = []ParseRule{}

var parser = Parser{}

var compilingChunk *Chunk

func currentChunk() *Chunk {
	return compilingChunk
}

func errorAt(token *Token, message string) {
	// Ignore all the rest of the tokens when in panic mode
	if parser.PanicMode {
		return
	}

	parser.PanicMode = true

	fmt.Fprintf(os.Stderr, "[line %d] Error", token.Line)

	if token.Type == TokenEOF {
		fmt.Fprintf(os.Stderr, " at end")
	} else if token.Type == TokenError {
		// Do nothing
	} else {
		fmt.Fprintf(os.Stderr, " at '%s'", token.Value)
	}

	fmt.Fprintf(os.Stderr, ": %s\n", message)
	parser.HadError = true
}

func errorAtCurrent(message string) {
	errorAt(&parser.Current, message)
}

func errorAtPrev(message string) {
	errorAt(&parser.Previous, message)
}

func advanceParser() {
	parser.Previous = parser.Current

	for {
		parser.Current = ScanToken()
		if parser.Current.Type != TokenError {
			break
		}

		errorAtCurrent(parser.Current.Value)
	}
}

func consumeToken(_type TokenType, message string) {
	if parser.Current.Type == _type {
		advanceParser()
		return
	}

	errorAtCurrent(message)
}

func emitByte(_byte uint8) {
	currentChunk().WriteChunk(_byte, parser.Previous.Line)
}

func emitBytes(byte1, byte2 uint8) {
	emitByte(byte1)
	emitByte(byte2)
}

func emitReturn() {
	emitByte(OpReturn)
}

func makeConstant(value Value) uint8 {
	constant := currentChunk().AddConstant(value)
	if constant > math.MaxInt8 {
		errorAtPrev("Too many constant in one chunk")
		return 0
	}

	return uint8(constant)
}

func emitConstant(value Value) {
	emitBytes(OpConstant, makeConstant(value))
}

func endCompiler() {
	emitReturn()
	if DebugPrintCode && !parser.HadError {
		currentChunk().DisassembleChunk("code")
	}
}

func parseBinary() {
	// Remember the operator
	operatorType := parser.Previous.Type

	// Compile the right operand
	rule := getRule(operatorType)
	parsePrecedence(rule.Precedence + 1)

	// Emit the operator instruction
	switch operatorType {
	case TokenBangEqual:
		emitBytes(OpEqual, OpNot)
		break
	case TokenEqualEqual:
		emitByte(OpEqual)
		break
	case TokenGreater:
		emitByte(OpGreater)
		break
	case TokenGreaterEqual:
		emitBytes(OpLess, OpNot)
		break
	case TokenLess:
		emitByte(OpLess)
		break
	case TokenLessEqual:
		emitBytes(OpGreater, OpNot)
		break
	case TokenPlus:
		emitByte(OpAdd)
		break
	case TokenMinus:
		emitByte(OpSubtract)
		break
	case TokenStar:
		emitByte(OpMultiply)
		break
	case TokenSlash:
		emitByte(OpDivide)
		break
	default:
		break // Unreachable
	}
}

func parseLiteral() {
	switch parser.Previous.Type {
	case TokenFalse:
		emitByte(OpFalse)
		break
	case TokenNil:
		emitByte(OpNil)
		break
	case TokenTrue:
		emitByte(OpTrue)
		break
	default:
		return // Unreachable
	}
}

func parseExpression() {
	parsePrecedence(PrecAssignment)
}

func parseGrouping() {
	parseExpression()
	consumeToken(TokenRightParen, "Expect ')' after expression")
}

func parseNumber() {
	value, _ := strconv.ParseFloat(parser.Previous.Value, 64)
	val := NumberVal(value)
	emitConstant(val)
}

func parseUnary() {
	operatorType := parser.Previous.Type

	// Compile the operand
	parsePrecedence(PrecUnary)

	switch operatorType {
	case TokenBang:
		emitByte(OpNot)
		break
	case TokenMinus:
		emitByte(OpNegate)
		break
	default:
		return // Unreachable
	}

}

func parsePrecedence(precedence Precedence) {
	advanceParser()
	prefixRule := getRule(parser.Previous.Type).Prefix

	if prefixRule == nil {
		errorAtPrev("Expect expression")
		return
	}

	prefixRule()

	for precedence <= getRule(parser.Current.Type).Precedence {
		advanceParser()
		infixRule := getRule(parser.Previous.Type).Infix
		infixRule()
	}

}

func getRule(_type TokenType) *ParseRule {
	return &rules[_type]
}

func initCompiler() {
	// Init parse rule table
	rules = []ParseRule{
		{parseGrouping, nil, PrecCall},      // TokenLeftParen
		{nil, nil, PrecNone},                // TokenRightParen
		{nil, nil, PrecNone},                // TokenLeftBrace
		{nil, nil, PrecNone},                // TokenRightBrace
		{nil, nil, PrecNone},                // TokenComma
		{nil, nil, PrecCall},                // TokenDot
		{parseUnary, parseBinary, PrecTerm}, // TokenMinus
		{nil, parseBinary, PrecTerm},        // TokenPlus
		{nil, nil, PrecNone},                // TokenSemicolon
		{nil, parseBinary, PrecFactor},      // TokenSlash
		{nil, parseBinary, PrecFactor},      // TokenStar
		{parseUnary, nil, PrecNone},         // TokenBang
		{nil, parseBinary, PrecEquality},    // TokenBangEqual
		{nil, nil, PrecNone},                // TokenEqual
		{nil, parseBinary, PrecEquality},    // TokenEqualEqual
		{nil, parseBinary, PrecComparison},  // TokenGreater
		{nil, parseBinary, PrecComparison},  // TokenGreaterEqual
		{nil, parseBinary, PrecComparison},  // TokenLess
		{nil, parseBinary, PrecComparison},  // TokenLessEqual
		{nil, nil, PrecNone},                // TokenIdentifier
		{nil, nil, PrecNone},                // TokenString
		{parseNumber, nil, PrecNone},        // TokenNumber
		{nil, nil, PrecAnd},                 // TokenAnd
		{nil, nil, PrecNone},                // TokenClass
		{nil, nil, PrecNone},                // TokenElse
		{parseLiteral, nil, PrecNone},       // TokenFalse
		{nil, nil, PrecNone},                // TokenFor
		{nil, nil, PrecNone},                // TokenFun
		{nil, nil, PrecNone},                // TokenIf
		{parseLiteral, nil, PrecNone},       // TokenNil
		{nil, nil, PrecNone},                // TokenOr
		{nil, nil, PrecNone},                // TokenPrint
		{nil, nil, PrecNone},                // TokenReturn
		{nil, nil, PrecNone},                // TokenSuper
		{nil, nil, PrecNone},                // TokenThis
		{parseLiteral, nil, PrecNone},       // TokenTrue
		{nil, nil, PrecNone},                // TokenVar
		{nil, nil, PrecNone},                // TokenWhile
		{nil, nil, PrecNone},                // TokenError
		{nil, nil, PrecNone},                // TokenEOF
	}
}

// Compile the source code
func Compile(source string, chunk *Chunk) bool {
	initCompiler()
	InitScanner(source)

	compilingChunk = chunk
	parser.HadError = false
	parser.PanicMode = false

	advanceParser()
	parseExpression()

	consumeToken(TokenEOF, "Expect end of expression")
	endCompiler()
	return !parser.HadError
}
