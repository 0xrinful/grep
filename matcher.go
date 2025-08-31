package main

func match(line, pattern string) bool {
	tokens := tokenize(pattern)
	for i := 0; i+len(tokens) <= len(line); i++ {
		if matchAt(line[i:], tokens) {
			return true
		}
	}

	return false
}

func matchAt(text string, tokens []Token) bool {
	if len(tokens) > len(text) {
		return false
	}

	if len(tokens) == 0 {
		return true
	}

	for i, t := range tokens {
		ch := text[i]
		switch t.kind {
		case Dot:
			continue
		case Word:
			if !isWord(ch) {
				return false
			}
		case Digit:
			if !isDigit(ch) {
				return false
			}
		case Space:
			if !isWhiteSpace(ch) {
				return false
			}
		case Literal:
			if t.ch != ch {
				return false
			}
		}
	}
	return true
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isWord(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') ||
		ch == '_'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t'
}
