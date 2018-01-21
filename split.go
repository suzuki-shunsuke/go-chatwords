// Split a string into substrings
package chatwords

import (
	"fmt"
)

// Split a string into substrings.
// The parameter text is splited to some strings.
// The parameter limit is the max size of returned array.
// If the size of the array is reaced to the limit, the rest text is returned as the second returned value.
// If the parameter limit is negative integer, the limit is ignored, and the second returned value is empty.
// Example:
//   // ["foo", "bar"], ""
//   args, txt := Split(" foo  bar  ", -1)
//   // ["echo", "hello world"], ""
//   args, txt := Split("echo 'hello world'", -1)
//   // ["echo"], " 'hello world'"
//   args, txt = Split("echo 'hello world'", 1)
func Split(text string, limit int) ([]string, string) {
	args := []string{}
	var squoted, dquoted, escaped, spaced, ended bool
	var preChar, word string
	if limit == 0 {
		return args, text
	}
	for i, c := range text {
		switch c {
		case '\'':
			switch {
			case escaped:
				word += string(c)
				escaped = false
			case squoted:
				ended = true
				squoted = false
			case spaced:
				squoted = true
			default:
				word += string(c)
			}
			preChar = string(c)
			spaced = false
		case '"':
			switch {
			case escaped:
				word += string(c)
				escaped = false
			case dquoted:
				ended = true
				dquoted = false
			case spaced:
				dquoted = true
			default:
				word += string(c)
			}
			preChar = string(c)
			spaced = false
		case ' ', '\t':
			switch {
			case squoted || dquoted:
				word += string(c)
			case !spaced:
				args = append(args, word)
				if limit == 1 {
					return args, text[i:]
				}
				limit--
				word = ""
				spaced = true
				escaped = false
			}
			preChar = string(c)
		case '\\':
			if escaped {
				word += string(c)
				escaped = false
			} else {
				escaped = true
			}
			spaced = false
			preChar = string(c)
		default:
			switch {
			case squoted || dquoted:
				word += string(c)
			case ended:
				ended = false
				switch preChar {
				case "'":
					words, txt := Split(fmt.Sprintf("\\'%s'%s", word, text[i:]), limit)
					for _, w := range words {
						args = append(args, w)
					}
					return args, txt
				case "\"":
					words, txt := Split(fmt.Sprintf("\\\"%s\\\"%s", word, text[i:]), limit)
					for _, w := range words {
						args = append(args, w)
					}
					return args, txt
				}
			default:
				word += string(c)
			}
			spaced = false
			escaped = false
			preChar = string(c)
		}
	}
	if word != "" {
		args = append(args, word)
	}
	return args, ""
}
