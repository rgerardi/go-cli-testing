package main

import (
	"bytes"
	"io/ioutil"
)

// START OMIT
type myStruct struct {
	data []byte
}

func newFromFile(filename string) (*myStruct, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &myStruct{data: data}, nil
}

func (m *myStruct) save(filename string) error {
	return ioutil.WriteFile(filename, m.data, 0644)
}

func (m *myStruct) doMagic() { // HL
	m.data = bytes.ToUpper(m.data) // HL
} // HL

// END OMIT

func main() {

	m := myStruct{
		data: []byte("Just some test data..."),
	}

	m.doMagic()

	m.save("file1.txt")
}
