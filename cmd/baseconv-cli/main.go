package main

import (
	"baseconv/internal/converter"
	"baseconv/internal/logging"
	"flag"
	"fmt"
	"strconv"
)

const (
	CONVERSION_HEX = "HEX"
	CONVERSION_OCT = "OCT"
	CONVERSION_BIN = "BIN"
	CONVERSION_DEC = "DEC"
	CONVERSION_ALL = "ALL"
)

func checkValidity(value string, dec, hex, oct, bin bool) {
	if dec && converter.IsDecimal(value) != nil {
		logging.Panic("Value '%s' is not in decimal form", value)
	}
	if hex && converter.IsHexadecimal(value) != nil {
		logging.Panic("Value '%s' is not in hexadecimal form", value)
	}
	if oct && converter.IsOctal(value) != nil {
		logging.Panic("Value '%s' is not in octal form", value)
	}
	if bin && converter.IsBinary(value) != nil {
		logging.Panic("Value '%s' is not in binary form", value)
	}
}

func main() {
	var value = ""
	var conv = ""
	var dec = false
	var hex = false
	var oct = false
	var bin = false

	flag.StringVar(&value, "value", value, "The value to convert")
	flag.StringVar(&conv, "conv", conv, "The conversion {DEC | HEX | OCT | BIN}") 
	flag.BoolVar(&dec, "dec", dec, "Set the base of the value to decimal")
	flag.BoolVar(&bin, "bin", bin, "Set the base of the value to binary")
	flag.BoolVar(&oct, "oct", oct, "Set the base of the value to octal")
	flag.BoolVar(&hex, "hex", hex, "Set the base of the value to hexadecimal")
	flag.Parse()

	if value == "" {
		logging.Panic("No value specified. See `conv -h`")
	}

	if conv == "" {
		logging.Panic("No conversion specified. See `conv -h`")
	}
	
	checkValidity(value, dec, hex, oct, bin)

	v := converter.Value{}

	if dec {
		decVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logging.Panic("%v", err)
		}
		err = v.UpdateDec(decVal)

		if err != nil {
			logging.Panic("%v", err)
		}
	} else if hex {
		err := v.UpdateHex(value)
		if err != nil {
			logging.Panic("%v", err)
		}
	} else if oct {
		err := v.UpdateOct(value)
		if err != nil {
			logging.Panic("%v", err)
		}
	} else if bin {
		err := v.UpdateBin(value)
		if err != nil {
			logging.Panic("%v", err)
		}
	}

	switch conv {
	case CONVERSION_DEC:
		fmt.Println(v.Dec)
	case CONVERSION_HEX:
		fmt.Println(v.Hex)
	case CONVERSION_OCT:
		fmt.Println(v.Oct)
	case CONVERSION_BIN:
		fmt.Println(v.Bin)
	case CONVERSION_ALL:
		fmt.Println(v.ToString())
	default:
		logging.Panic("Invalid conversion: %s", conv)
	}
}
