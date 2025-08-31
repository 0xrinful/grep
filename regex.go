package main

func contains(line, pattern string) bool {
	n := len(pattern)
	if n == 0 {
		return false
	}

	for i := 0; i+n <= len(line); i++ {
		if line[i:i+n] == pattern {
			return true
		}
	}
	return false
}

func match(line, pattern string) bool {
	if pattern == "" {
		return true
	}

	return contains(line, pattern)
}
