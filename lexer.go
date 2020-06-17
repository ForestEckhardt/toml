package toml

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type lexer struct {
	input string
	items chan item
	start int
	pos   int
	width []int

	lexStack []stateFn
}

type item struct {
	typ itemType
	val string
}

type stateFn func(*lexer) stateFn
type acceptFn func(rune) bool

func (l *lexer) next() rune {
	if l.pos == len(l.input) {
		l.width = append(l.width, 0)
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = append(l.width, w)
	l.pos += w
	return r
}

func (l *lexer) nextBlock(n int) {
	for i := 0; i < n; i++ {
		l.next()
	}
}

func (l *lexer) backup() {
	length := len(l.width) - 1
	l.pos -= l.width[length]
	l.width = l.width[:length]
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) peekBlock(n int) string {
	var runes []rune

	for i := 0; i < n; i++ {
		runes = append(runes, l.next())
	}

	for i := 0; i < n; i++ {
		l.backup()
	}

	return string(runes)
}

func (l *lexer) emit(t itemType) {
	l.items <- item{typ: t, val: l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func (l *lexer) acceptFunc(acceptFn acceptFn) bool {
	if acceptFn(l.next()) {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptFuncRun(acceptFn acceptFn) {
	for acceptFn(l.next()) {
	}
	l.backup()
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) skip() {
	l.next()
	l.ignore()
}

func (l *lexer) skipBlock(n int) {
	for i := 0; i < n; i++ {
		l.next()
		l.ignore()
	}
}

func (l *lexer) errorf(format string, values ...interface{}) stateFn {
	l.items <- item{typ: itemError, val: fmt.Sprintf(format, values...)}
	return nil
}

func (l *lexer) pushLexStack(fn stateFn) {
	l.lexStack = append(l.lexStack, fn)
}

func (l *lexer) popLexStack() stateFn {
	var fn stateFn
	length := len(l.lexStack) - 1
	fn, l.lexStack = l.lexStack[length], l.lexStack[:length]
	return fn
}

func (l *lexer) nextItem() item {
	return <-l.items
}

func lex(input string) *lexer {
	l := &lexer{
		input: input,
		items: make(chan item),
	}

	go l.run()
	return l
}

func (l *lexer) run() {
	for state := lexTOML; state != nil; {
		state = state(l)
	}

	close(l.items)
}

func lexTOML(l *lexer) stateFn {
	switch r := l.peek(); {
	case isWhitespace(r) || isNewline(r):
		l.skip()
		return lexTOML
	case r == commentStart:
		l.skip()
		l.pushLexStack(lexTOML)
		return lexComment
	case r == eof:
		if l.pos > l.start {
			panic("Unexpected eof")
			// return l.errorf("Unexpected eof")
		}
		l.emit(itemEOF)
		return nil
	default:
		l.pushLexStack(lexTOML)
		return lexKeyStart
	}
}

func lexValue(l *lexer) stateFn {
	switch r := l.peek(); {
	case isWhitespace(r):
		l.skip()
		return lexValue
	case r == stringDelimiter:
		if l.peekBlock(3) == `"""` {
			l.skipBlock(3)
			return lexMultiLineBasicString
		}
		l.skip()
		return lexBasicString
	case r == rawStringDelimiter:
		if l.peekBlock(3) == `'''` {
			l.skipBlock(3)
			return lexMultiLineRawString
		}
		l.skip()
		return lexRawString
	case unicode.IsDigit(r) || r == 'i' || r == 'n' || r == '+' || r == '-':
		return lexStartNumber
	case r == 't' || r == 'f':
		return lexBoolean
	default:
		panic("lexValue")
		// return l.errorf("expected value found %q", r)
	}
}

//Rune type checking functions
func isWhitespace(r rune) bool {
	return r == '\t' || r == ' '
}

func isNewline(r rune) bool {
	return r == '\n' || r == '\r'
}

func isHexadecimal(r rune) bool {
	return (r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'f') ||
		(r >= 'A' && r <= 'F')
}

func isOctal(r rune) bool {
	return (r >= '0' && r <= '7')
}

func isBinary(r rune) bool {
	return r == '0' || r == '1'
}

func allowedEscapeCharacter(r rune) bool {
	switch r {
	case '"':
		return true
	case '\\':
		return true
	default:
		return strings.ContainsRune("btnfr", r)
	}
}

func isDateTimeCharacter(r rune) bool {
	//The space is a valid character
	return strings.ContainsRune(`-:TZ+. `, r)
}

// Run functions
func commentAcceptance(r rune) bool {
	return !isNewline(r) && r != eof
}

func keyAcceptance(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-'
}

func basicStringAcceptance(r rune) bool {
	return r != stringDelimiter && r != escapeCharacter && !isNewline(r)
}

func rawStringAcceptance(r rune) bool {
	return r != rawStringDelimiter && !isNewline(r)
}

func integerAcceptance(r rune) bool {
	return unicode.IsDigit(r)
}

func hexadecimalAcceptance(r rune) bool {
	return isHexadecimal(r)
}

func octalAcceptance(r rune) bool {
	return isOctal(r)
}

func binaryAcceptance(r rune) bool {
	return isBinary(r)
}

func dateTimeAcceptance(r rune) bool {
	return unicode.IsDigit(r) || isDateTimeCharacter(r)
}
