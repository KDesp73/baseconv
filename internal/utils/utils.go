package utils

import (
	"baseconv/internal/converter"
)

func PrintableForm(s string) (string) {
	if converter.IsCharacter(s) != nil{
		return ""
	}

	if len(s) == 1 {
		return s
	}

	switch s {
	case "\\a":
		return "\\a"
	case "\\b":
		return "\\b"
	case "\\f":
		return "\\f"
	case "\\n":
		return "\\n"
	case "\\r":
		return "\\r"
	case "\\t":
		return "\\t"
	case "\\v":
		return "\\v"
	case "\\'":
		return "\\'"
	case "\\\"":
		return "\\\""
	case "\\\\":
		return "\\\\"
	case "\\0":
		return "\\0"
	default:
		return ""
	}
}
