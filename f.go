package f

func Format(template string, data ...any) string {
	res, _ := FormatSafe(template, data)
	return res
}

func FormatSafe(template string, data ...any) (string, error) {
	var d any

	if len(data) > 0 {
		d = data[0]
	}

	parsed := Parse(template)

	return parsed.Eval(d)
}
