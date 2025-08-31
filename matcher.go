package main

func match(line, pattern string) bool {
	tokens := tokenize(pattern)
	for i := 0; i+len(tokens) <= len(line); i++ {
		if matchAt(line[i:], i, tokens) {
			return true
		}
	}

	return false
}

func matchAt(text string, i int, tokens []Token) bool {
	if len(tokens) == 0 {
		return true
	}
	t := tokens[0]

	if i >= len(text) {
		switch t.quant {
		case ZeroOrMore, ZeroOrOne:
			return matchAt(text, i, tokens[1:])
		default:
			return false
		}
	}

	switch t.quant {
	case One:
		return matchToken(text[i], t) &&
			matchAt(text, i+1, tokens[1:])

	case ZeroOrMore:
		j := i
		for j < len(text) && matchToken(text[j], t) {
			j++
		}
		for k := j; k >= i; k-- {
			if matchAt(text, k, tokens[1:]) {
				return true
			}
		}
		return false

	case OneOrMore:
		if !matchToken(text[i], t) {
			return false
		}
		j := i + 1
		for j < len(text) && matchToken(text[j], t) {
			j++
		}
		for k := j; k >= i+1; k-- {
			if matchAt(text, k, tokens[1:]) {
				return true
			}
		}
		return false

	case ZeroOrOne:
		if matchToken(text[i], t) {
			if matchAt(text, i+1, tokens[1:]) {
				return true
			}
		}
		return matchAt(text, i, tokens[1:])
	}
	return false
}

func matchToken(ch byte, t Token) bool {
	switch t.kind {
	case Dot:
		return true
	case Word:
		return isWord(ch)
	case Digit:
		return isDigit(ch)
	case Space:
		return isWhiteSpace(ch)
	case Literal:
		return t.ch == ch
	}
	return false
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
