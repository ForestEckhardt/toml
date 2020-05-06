package toml

func lexComment(l *lexer) stateFn {
	l.acceptFuncRun(commentAcceptance)
	l.ignore()
	return lexTOML
}
