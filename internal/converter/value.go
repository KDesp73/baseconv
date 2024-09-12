package converter

import (
	"fmt"
	"strings"
)


type Value struct {
	Dec int64
	Hex string
	Oct string
	Bin string
	Char string
}

func (v *Value) ToString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Dec: %d\n", v.Dec))
	b.WriteString(fmt.Sprintf("Hex: %s\n", v.Hex))
	b.WriteString(fmt.Sprintf("Oct: %s\n", v.Oct))
	b.WriteString(fmt.Sprintf("Bin: %s\n", v.Bin))
	b.WriteString(fmt.Sprintf("Char: %s\n", v.Char))

	return b.String()
}

func (v *Value) UpdateDec(dec int64) error {
	var err error
	v.Dec = dec
	v.Hex = DecToHex(dec)
	v.Bin = DecToBin(dec)
	v.Oct = DecToOct(dec)
	v.Char, err = DecToChar(dec)

	return err
}

func (v *Value) UpdateHex(hex string) error {
	var err error
	v.Dec, err = HexToDec(hex)
	v.Hex = hex
	v.Bin, err = HexToBin(hex)
	v.Oct, err = HexToOct(hex)
	v.Char, err = HexToChar(hex)

	return err
}

func (v *Value) UpdateBin(bin string) error {
	var err error
	v.Dec, err = BinToDec(bin)
	v.Hex, err = BinToHex(bin)
	v.Bin = bin
	v.Oct, err = BinToOct(bin)
	v.Char, err = BinToChar(bin)

	return err
} 

func (v *Value) UpdateOct(oct string) error {
	var err error
	v.Dec, err = OctToDec(oct)
	v.Hex, err = OctToHex(oct)
	v.Bin, err = OctToBin(oct)
	v.Oct = oct
	v.Char, err = OctToChar(oct)

	return err
} 

func (v *Value) UpdateChar(char string) error {
	var err error
	v.Dec, err = CharToDec(char)
	v.Hex, err = CharToHex(char)
	v.Bin, err = CharToBin(char)
	v.Oct, err = CharToOct(char)
	v.Char = char

	return err
}

func (v *Value) Reset() {
	v.Dec = 0
	v.Hex = ""
	v.Bin = ""
	v.Oct = ""
	v.Char = ""
}
