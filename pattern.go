package main

type TokenType uint8

const (
	Literal TokenType = iota
	Dot
	Word
	Digit
	Space
)

type Token struct {
	kind TokenType
	ch   byte
}

func tokenize(pattern string) []Token {
	var tokens []Token
	n := len(pattern)

	for i := 0; i < n; i++ {
		switch pattern[i] {
		case '.':
			tokens = append(tokens, Token{kind: Dot})
		case '\\':
			if i+1 < n {
				switch pattern[i+1] {
				case 'w':
					tokens = append(tokens, Token{kind: Word})
				case 'd':
					tokens = append(tokens, Token{kind: Digit})
				case 's':
					tokens = append(tokens, Token{kind: Space})
				default:
					tokens = append(tokens, Token{kind: Literal, ch: pattern[i+1]})
				}
				i++
			}
		default:
			tokens = append(tokens, Token{kind: Literal, ch: pattern[i]})
		}
	}

	return tokens
}
