package main

import "golang.org/x/tour/reader"

type (
	MyReader  struct{}
	ErrReader int
)

func (e ErrReader) Error() string {
	return "0-length slice; Can't do it"
}

func (readerInstance MyReader) Read(byteSlice []byte) (int, error) {

	if len(byteSlice) <= 0 {
		return 0, ErrReader(0)
	}

	for index := range byteSlice {
		byteSlice[index] = 65
	}

	return len(byteSlice), nil
}

func main() {
	reader.Validate(MyReader{})
}
