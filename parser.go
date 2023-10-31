package f

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antonmedv/expr"
)

// Parsed contains a parsed template string, ready to be evaluated.
type Parsed struct {
	Template string
	Parts    []Part
}

// Parse parses a template string into a Parsed struct.
func Parse(template string) Parsed {
	return Parsed{
		Template: template,
		Parts:    splitIntoParts(template),
	}
}

// String returns the parsed template parts as a single string.
func (parsed Parsed) String() string {
	var res strings.Builder
	for _, part := range parsed.Parts {
		res.WriteString(part.Value)
	}

	return res.String()
}

// Eval evaluated expressions in the parsed template string.
func (parsed Parsed) Eval(data any) (string, error) {
	for i, part := range parsed.Parts {
		if part.Parsed {
			continue
		}

		exp, err := expr.Compile(part.Value, expr.Env(data))
		if err != nil {
			return "", fmt.Errorf("failed to compile expression: %w", err)
		}

		result, err := expr.Run(exp, data)
		if err != nil {
			return "", fmt.Errorf("failed to evaluate expression: %w", err)
		}

		parsed.Parts[i].Value = fmt.Sprintf("%v", result)
	}

	return parsed.String(), nil
}

// Part is a single part of a template string.
// Can either be a raw string, or an expression.
type Part struct {
	Value  string
	Parsed bool
}

func splitIntoParts(template string) []Part {
	r := regexp.MustCompile(`\$\{([^\}]*)\}`)
	matches := r.FindAllStringSubmatchIndex(template, -1)

	parts := make([]Part, 0, len(matches)+1)

	lastIndex := 0

	for _, match := range matches {
		start, end := match[0], match[1]
		exprStart, exprEnd := match[2], match[3]

		// Add the string before the match, if any
		if start > lastIndex {
			parts = append(parts, Part{
				Value:  template[lastIndex:start],
				Parsed: true,
			})
		}

		// Add the expression
		parts = append(parts, Part{
			Value:  template[exprStart:exprEnd],
			Parsed: false,
		})

		lastIndex = end
	}

	// Add remaining string, if any
	if lastIndex < len(template) {
		parts = append(parts, Part{
			Value:  template[lastIndex:],
			Parsed: true,
		})
	}

	return parts
}
