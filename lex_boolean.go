package toml

func lexBoolean(l *lexer) stateFn {
	if l.peekBlock(4) == `true` {
		l.nextBlock(4)
		l.emit(itemBoolean)
		return l.popLexStack()
	}

	if l.peekBlock(5) == `false` {
		l.nextBlock(5)
		l.emit(itemBoolean)
		return l.popLexStack()
	}

	panic("lexBoolean")
}
