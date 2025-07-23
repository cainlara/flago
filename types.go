package flago

import "errors"

var (
	errEmptyInput   = errors.New("no valid input provided")
	errUnknownField = errors.New("no such field [] in provided output struct")
)

const dash = "-"
