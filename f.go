package f

func Format(template string, data any) string {
	res, _ := FormatSafe(template, data)
	return res
}

func FormatSafe(template string, data any) (string, error) {
	parsed := Parse(template)
	return parsed.Eval(data)
}