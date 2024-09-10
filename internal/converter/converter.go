package converter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Value struct {
	Dec int64
	Hex string
	Oct string
	Bin string
}

func (v *Value) ToString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Dec: %d\n", v.Dec))
	b.WriteString(fmt.Sprintf("Hex: %s\n", v.Hex))
	b.WriteString(fmt.Sprintf("Oct: %s\n", v.Oct))
	b.WriteString(fmt.Sprintf("Bin: %s\n", v.Bin))

	return b.String()
}

func (v *Value) UpdateDec(dec int64) error {
	v.Dec = dec
	v.Hex = DecToHex(dec)
	v.Bin = DecToBin(dec)
	v.Oct = DecToOct(dec)

	return nil
}

func (v *Value) UpdateHex(hex string) error {
	var err error
	v.Dec, err = HexToDec(hex)
	v.Hex = hex
	v.Bin, err = HexToBin(hex)
	v.Oct, err = HexToOct(hex)

	return err
}

func (v *Value) UpdateBin(bin string) error {
	var err error
	v.Dec, err = BinToDec(bin)
	v.Hex, err = BinToHex(bin)
	v.Bin = bin
	v.Oct, err = BinToOct(bin)

	return err
} 

func (v *Value) UpdateOct(oct string) error {
	var err error
	v.Dec, err = OctToDec(oct)
	v.Hex, err = OctToHex(oct)
	v.Bin, err = OctToBin(oct)
	v.Oct = oct

	return err
} 

func (v *Value) Reset() {
	v.Dec = 0
	v.Hex = "0"
	v.Bin = "0"
	v.Oct = "0"
}

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
