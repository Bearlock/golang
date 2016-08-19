package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

type ErrReader struct{}

func (e ErrReader) Error() string {
	return "0-length slice; cannot do"
}

func (readerInstance *rot13Reader) Read(byteSlice []byte) (totalBytes int, err error) {
	
	totalBytes, err = readerInstance.r.Read(byteSlice)

	if len(byteSlice) <= 0 {
		return 0, ErrReader{}
	}
	
	for index, value := range byteSlice {
		byteSlice[index] = asciiTransform(value)
	}
	
	return totalBytes, err
}
	
func asciiTransform(asciiValue byte) byte {
	if asciiValue > 64 && asciiValue < 91 {
		return upperCaseAsciiTransform(asciiValue)
	} else if asciiValue > 96 && asciiValue < 123 {
		return lowerCaseAsciiTransform(asciiValue)
	} else {
		return asciiValue
	}
}

func upperCaseAsciiTransform(asciiValue byte) byte {
	if asciiValue < 78 {
		return asciiValue + 13
	} else {
		return asciiValue - 13
	}
}

func lowerCaseAsciiTransform(asciiValue byte) byte {
	if asciiValue < 110 {
		return asciiValue + 13
	} else {
		return asciiValue - 13
	}
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
