package f

// Format formats the template string.
func Format(template string, data ...any) string {
	res, _ := FormatSafe(template, data...)
	return res
}

// FormatSafe formats the template string and returns an additional, optional error, if something goes wrong.
func FormatSafe(template string, data ...any) (string, error) {
	var d any

	if len(data) > 0 {
		d = data[0]
	}

	parsed := Parse(template)

	return parsed.Eval(d)
}
