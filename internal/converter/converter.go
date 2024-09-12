package converter

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)


func conversionPanic(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s (%v)", msg, err)
	os.Exit(1)
}

func HexToDec(hex string) (int64, error) {
	return strconv.ParseInt(hex, 16, 64)
}

func DecToHex(dec int64) (string) {
	return fmt.Sprintf("%X", dec)
}

func BinToDec(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 64)
}

func DecToBin(dec int64) string {
	return fmt.Sprintf("%b", dec)
}

func HexToBin(hex string) (string, error) {
	dec, err := HexToDec(hex)

	if err != nil {
		return "", err
	}

	return DecToBin(dec), nil
}

func BinToHex(bin string) (string, error) {
	dec, err := BinToDec(bin)

	if err != nil {
		return "", err
	}

	return DecToHex(dec), nil
}

func DecToOct(decimal int64) string {
	return fmt.Sprintf("%o", decimal)
}

func OctToDec(octalStr string) (int64, error) {
	return strconv.ParseInt(octalStr, 8, 64)
}

func OctToBin(octalStr string) (string, error) {
	decimalValue, err := OctToDec(octalStr)
	if err != nil {
		return "", err
	}
	return DecToBin(decimalValue), nil
}

func BinToOct(binaryStr string) (string, error) {
	decimalValue, err := BinToDec(binaryStr)
	if err != nil {
		return "", err
	}
	return DecToOct(decimalValue), nil
}

func HexToOct(hexStr string) (string, error) {
	decimalValue, err := HexToDec(hexStr)
	if err != nil {
		return "", err
	}
	return DecToOct(decimalValue), nil
}

func OctToHex(octalStr string) (string, error) {
	decimalValue, err := OctToDec(octalStr)
	if err != nil {
		return "", err
	}
	return DecToHex(decimalValue), nil
}

func DecToChar(dec int64) (string, error) {
	if dec < 0 || dec > 127 {
		return "", fmt.Errorf("Decimal value out of ASCII range: %d", dec)
	}
	if dec == 0 {
		return "\\0", nil
	}
	return string(dec), nil
}

func CharToDec(s string) (int64, error) {
	if IsCharacter(s) != nil {
		return 0, fmt.Errorf("'%s' is not a valid character or escape code", s)
	}

	// If it's a single character, return its ASCII value
	if len(s) == 1 {
		return int64(s[0]), nil
	}

	// Handle escape codes
	switch s {
	case "\\a":
		return 7, nil  // Bell
	case "\\b":
		return 8, nil  // Backspace
	case "\\f":
		return 12, nil // Form feed
	case "\\n":
		return 10, nil // Newline
	case "\\r":
		return 13, nil // Carriage return
	case "\\t":
		return 9, nil  // Horizontal tab
	case "\\v":
		return 11, nil // Vertical tab
	case "\\'":
		return 39, nil // Single quote
	case "\\\"":
		return 34, nil // Double quote
	case "\\\\":
		return 92, nil  // Backslash
	case "\\0":
		return 0, nil // Null terminator
	default:
		return -1, fmt.Errorf("unknown escape code: %s", s)
	}
}

func HexToChar(hex string) (string, error) {
	dec, err := HexToDec(hex)
	if err != nil {
		return "", err
	}
	return DecToChar(dec)
}

func CharToHex(char string) (string, error) {
	dec, err := CharToDec(char)
	if err != nil {
		return "", err
	}

	return DecToHex(dec), nil
}

func CharToBin(char string) (string, error) {
	dec, err := CharToDec(char)
	if err != nil {
		return "", err
	}

	return DecToBin(dec), nil
}

func CharToOct(char string) (string, error) {
	dec, err := CharToDec(char)
	if err != nil {
		return "", err
	}

	return DecToOct(dec), nil
}

func BinToChar(bin string) (string, error) {
	dec, err := BinToDec(bin)
	if err != nil {
		return "", err
	}
	return DecToChar(dec)
}

func OctToChar(oct string) (string, error) {
	dec, err := OctToDec(oct)
	if err != nil {
		return "", err
	}
	return DecToChar(dec)
}

func IsDecimal(s string) error {
	_, err := strconv.ParseInt(s, 10, 64)
	return err
}

func IsHexadecimal(s string) error {
	_, err := strconv.ParseInt(s, 16, 64)
	return err
}

func IsOctal(s string) error {
	_, err := strconv.ParseInt(s, 8, 64)
	return err
}

func IsBinary(s string) error {
	_, err := strconv.ParseInt(s, 2, 64)
	return err
}

func IsCharacter(s string) error {
	if len(s) == 1 {
		return nil
	}

	escapeCodePattern := `^\$$abfnrtv'"\$$$`
	var err error
	var matched bool
	if matched, err = regexp.MatchString(escapeCodePattern, s); matched {
		return nil
	} 

	return err
}
