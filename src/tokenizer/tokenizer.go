package tokenizer

import (
	"html"
	"strings"

	"github.com/mdaushi/bahasago/src/fuzzy"
)

// Tokenization splits a source into words.
// Tokenizer can also normalize the sentence before tokenizing, to do that provide a true as second argument to the method tokenize
func Tokenize(source string, normalize ...bool) []string {
	nm := false
	if normalize != nil {
		nm = normalize[0]
	}

	if nm {
		source = fuzzy.Normalize(source)
	}

	source = html.UnescapeString(source)
	source = strings.TrimSpace(source)
	source = strings.ToLower(source)
	return strings.Fields(source)
}
