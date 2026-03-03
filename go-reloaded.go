package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		return
	}

	text := string(content)

	tokens := Tokenize(text)
	tokens = ApplyModifiers(tokens)
	tokens = FixAAn(tokens)
	result := Rebuild(tokens)

	result += "\n"

	os.WriteFile(outputFile, []byte(result), 0644)
}

//////////////////////////////////////////////////
// TOKENIZER
//////////////////////////////////////////////////

func Tokenize(text string) []string {
	var tokens []string
	current := ""

	for _, r := range text {

		if unicode.IsSpace(r) {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
		} else if strings.ContainsRune(".,!?;:()'", r) { // added '
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, string(r))
		} else {
			current += string(r)
		}
	}

	if current != "" {
		tokens = append(tokens, current)
	}

	return tokens
}

//////////////////////////////////////////////////
// MODIFIERS
//////////////////////////////////////////////////

func ApplyModifiers(tokens []string) []string {
	var result []string

	for i := 0; i < len(tokens); i++ {

		// Simple modifiers (up) (low) (cap) (hex) (bin)
		if tokens[i] == "(" && i+2 < len(tokens) && tokens[i+2] == ")" {

			if len(result) > 0 {
				last := len(result) - 1

				switch tokens[i+1] {
				case "hex":
					if val, err := strconv.ParseInt(result[last], 16, 64); err == nil {
						result[last] = strconv.Itoa(int(val))
					}
				case "bin":
					if val, err := strconv.ParseInt(result[last], 2, 64); err == nil {
						result[last] = strconv.Itoa(int(val))
					}
				case "up":
					result[last] = strings.ToUpper(result[last])
				case "low":
					result[last] = strings.ToLower(result[last])
				case "cap":
					result[last] = Capitalize(result[last])
				}
			}

			i += 2
			continue
		}

		// Modifiers with number (up, 2)
		if tokens[i] == "(" && i+4 < len(tokens) &&
			tokens[i+2] == "," && tokens[i+4] == ")" {

			mode := tokens[i+1]
			count, err := strconv.Atoi(tokens[i+3])

			if err == nil {
				for j := 0; j < count && len(result)-1-j >= 0; j++ {
					index := len(result) - 1 - j

					switch mode {
					case "up":
						result[index] = strings.ToUpper(result[index])
					case "low":
						result[index] = strings.ToLower(result[index])
					case "cap":
						result[index] = Capitalize(result[index])
					}
				}
			}

			i += 4
			continue
		}

		result = append(result, tokens[i])
	}

	return result
}

//////////////////////////////////////////////////
// FIX a → an
//////////////////////////////////////////////////

func FixAAn(tokens []string) []string {
	for i := 0; i < len(tokens)-1; i++ {

		if strings.ToLower(tokens[i]) == "a" {

			if len(tokens[i+1]) > 0 {
				first := unicode.ToLower(rune(tokens[i+1][0]))

				if strings.ContainsRune("aeiouh", first) {
					if tokens[i] == "A" {
						tokens[i] = "An"
					} else {
						tokens[i] = "an"
					}
				}
			}
		}
	}
	return tokens
}

//////////////////////////////////////////////////
// REBUILD TEXT
//////////////////////////////////////////////////

func Rebuild(tokens []string) string {
	var result string
	insideQuote := false

	for i := 0; i < len(tokens); i++ {

		// Handle ...
		if tokens[i] == "." && i+2 < len(tokens) &&
			tokens[i+1] == "." && tokens[i+2] == "." {

			result = strings.TrimRight(result, " ")
			result += "..."
			i += 2
			continue
		}

		// Handle !?
		if tokens[i] == "!" && i+1 < len(tokens) && tokens[i+1] == "?" {

			result = strings.TrimRight(result, " ")
			result += "!?"
			if i+2 < len(tokens) {
				result += " "
			}
			i++
			continue
		}

		// Handle apostrophe '
		if tokens[i] == "'" {

			result = strings.TrimRight(result, " ")
			result += "'"

			if insideQuote {
				// closing quote
				if i < len(tokens)-1 {
					result += " "
				}
				insideQuote = false
			} else {
				// opening quote
				insideQuote = true
			}
			continue
		}

		// Normal punctuation
		if strings.Contains(".,!?;:", tokens[i]) {
			result = strings.TrimRight(result, " ")
			result += tokens[i]

			if i < len(tokens)-1 {
				result += " "
			}

		} else {
			result += tokens[i]

			if i < len(tokens)-1 {
				result += " "
			}
		}
	}

	return strings.TrimSpace(result)
}

//////////////////////////////////////////////////
// CAPITALIZE
//////////////////////////////////////////////////

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
